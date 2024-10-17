package main

import (
	"github.com/c9s/goprocinfo/linux"
	"github.com/infrasonar/go-libagent"
)

func readUptime(state map[string][]map[string]any) error {
	uptime, err := linux.ReadUptime("/proc/uptime")
	if err != nil {
		return err
	}

	item := map[string]any{
		"name":  "uptime",
		"idle":  libagent.IFloat64(uptime.Idle),
		"total": libagent.IFloat64(uptime.Total),
	}

	state["uptime"] = []map[string]any{item}
	return nil
}

func readLoadAvg(state map[string][]map[string]any, check *libagent.Check) error {
	loadAvg, err := linux.ReadLoadAvg("/proc/loadavg")
	if err != nil {
		return err
	}

	var load float64

	if check.Interval < 300 {
		load = loadAvg.Last1Min
	} else if check.Interval < 900 {
		load = loadAvg.Last5Min
	} else {
		load = loadAvg.Last15Min
	}

	item := map[string]any{
		"name":           "loadAvg",
		"load":           libagent.IFloat64(load),
		"processRunning": loadAvg.ProcessRunning,
		"processTotal":   loadAvg.ProcessTotal,
	}

	state["loadAvg"] = []map[string]any{item}
	return nil
}

func CheckSystem(check *libagent.Check) (map[string][]map[string]any, error) {
	state := map[string][]map[string]any{}
	var err error

	if e := readUptime(state); e != nil {
		err = e
	}
	if e := readLoadAvg(state, check); e != nil {
		err = e
	}

	return state, err
}
