package main

import (
	"github.com/c9s/goprocinfo/linux"
	"github.com/infrasonar/go-libagent"
)

func readMemory(state map[string][]map[string]any) error {
	mem, err := linux.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return err
	}

	item := map[string]any{
		"name":  "memory",
		"free":  mem.MemFree,
		"total": mem.MemTotal,
	}

	state["memory"] = []map[string]any{item}
	return nil
}

func CheckMemory(check *libagent.Check) (map[string][]map[string]any, error) {
	state := map[string][]map[string]any{}
	var err error

	if e := readMemory(state); e != nil {
		err = e
	}

	return state, err
}
