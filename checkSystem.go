package main

import (
	"github.com/c9s/goprocinfo/linux"
)

func CheckSystem() (map[string][]map[string]any, error) {
	state := map[string][]map[string]any{}

	uptime, err := linux.ReadUptime("/proc/uptime")
	if err != nil {
		return nil, err
	}

	item := map[string]any{
		"name":  "uptime",
		"idle":  uptime.Idle,
		"total": uptime.Total,
	}

	items := []map[string]any{item}

	state["uptime"] = items
	return state, nil
}
