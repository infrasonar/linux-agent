package main

import (
	"github.com/c9s/goprocinfo/linux"
	"github.com/infrasonar/go-libagent"
)

func readMemInfo(state map[string][]map[string]any) error {
	mem, err := linux.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return err
	}

	item := map[string]any{
		"name":     	"memInfo",
		"memFree":  	mem.MemFree,
		"memTotal": 	mem.MemTotal,
		"memAvailable": mem.MemAvailable,
		"buffers":  	mem.Buffers,
		"cached":  		mem.Cached,
		"swapTotal":  	mem.SwapTotal,
		"swapFree":  	mem.SwapFree,
	}

	state["memInfo"] = []map[string]any{item}
	return nil
}

func CheckMemInfo(check *libagent.Check) (map[string][]map[string]any, error) {
	state := map[string][]map[string]any{}

	err := readMemInfo(state)
	if err != nil {
		return nil, err
	}

	return state, nil
}
