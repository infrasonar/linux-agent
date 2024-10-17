package main

import (
	"strings"

	"github.com/c9s/goprocinfo/linux"
	"github.com/infrasonar/go-libagent"
)

func readMounts(state map[string][]map[string]any) error {
	mounts, err := linux.ReadMounts("/proc/mounts")
	if err != nil {
		return err
	}

	items := []map[string]any{}

	for _, mount := range mounts.Mounts {
		if strings.HasPrefix(mount.Device, "/dev/loop") {
			// Skip loop devices
			continue
		}

		disk, err := linux.ReadDisk(mount.MountPoint)
		if err == nil && disk.All > 0 {

			items = append(items, map[string]any{
				"name":        mount.MountPoint,
				"device":      mount.Device,
				"fsType":      mount.FSType,
				"total":       disk.All,
				"free":        disk.Free,
				"used":        disk.Used,
				"percentUsed": libagent.IFloat64(float64(disk.Used) / float64(disk.All) * 100.0),
				"freeInodes":  disk.FreeInodes,
			})
		}
	}
	state["mounts"] = items
	return nil
}

func readDiskStats(state map[string][]map[string]any) error {
	diskStats, err := linux.ReadDiskStats("/proc/diskstats")
	if err != nil {
		return err
	}

	items := []map[string]any{}

	for _, ds := range diskStats {
		if strings.HasPrefix(ds.Name, "loop") {
			// Skip loop devices
			continue
		}
		items = append(items, map[string]any{
			"name":        ds.Name,
			"readBytes":   ds.GetReadBytes(),
			"writeBytes":  ds.GetWriteBytes(),
			"ioTicks":     ds.IOTicks,
			"inFlight":    ds.InFlight,
			"major":       ds.Major,
			"minor":       ds.Minor,
			"readIOs":     ds.ReadIOs,
			"readMerges":  ds.ReadMerges,
			"readTicks":   ds.ReadTicks,
			"timeInQueue": ds.TimeInQueue,
			"writeIOs":    ds.WriteIOs,
			"writeMerges": ds.WriteMerges,
			"writeTicks":  ds.WriteTicks,
		})
	}

	state["diskStats"] = items
	return nil
}

func CheckDisk(_ *libagent.Check) (map[string][]map[string]any, error) {
	state := map[string][]map[string]any{}
	var err error

	if err = readMounts(state); err != nil {
		return nil, err
	}

	if err = readDiskStats(state); err != nil {
		return nil, err
	}
	// Print debug dump
	// b, _ := json.MarshalIndent(state, "", "    ")
	// log.Fatal(string(b))

	return state, nil
}
