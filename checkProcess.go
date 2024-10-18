package main

import (
	"github.com/c9s/goprocinfo/linux"
	"github.com/infrasonar/go-libagent"
	"fmt"
)

func readProcess(state map[string][]map[string]any) error {
	pids, err := linux.ListPID("/proc", 1000)  // TODO which limit
	if err != nil {
		return err
	}

	names := map[string]int{}
	items := []map[string]any{}

	for _, pid := range pids {
		process, err := linux.ReadProcess(pid, "/proc")

		name := process.Status.Name  // TODO replace / character?
		ct := names[process.Status.Name]
		names[process.Status.Name] += 1

		if ct > 0 {
			name = fmt.Sprintf("%s_%d", name, ct)
		}

		if err == nil {
			items = append(items, map[string]any{
				"name":        name,
				"pid":         process.Status.Pid,
				"state":       process.Status.State,
				"threads":     process.Status.Threads,
				"vmPeak":      process.Status.VmPeak,
				"vmSize":      process.Status.VmSize,
				"vmLck":       process.Status.VmLck,
				"vmHWM":       process.Status.VmHWM,
				"vmRSS":       process.Status.VmRSS,
				"vmData":      process.Status.VmData,
				"vmStk":       process.Status.VmStk,
				"vmExe":       process.Status.VmExe,
				"vmLib":       process.Status.VmLib,
				"vmPTE":       process.Status.VmPTE,
				"vmSwap":      process.Status.VmSwap,
			})
		}
	}

	state["process"] = items
	return nil
}

func CheckProcess(check *libagent.Check) (map[string][]map[string]any, error) {
	state := map[string][]map[string]any{}

	err := readProcess(state)
	if err != nil {
		return nil, err
	}

	return state, nil
}
