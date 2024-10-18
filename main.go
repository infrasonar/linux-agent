package main

import (
	"log"

	"github.com/infrasonar/go-libagent"
)

func main() {
	// Start collector
	log.Printf("Starting InfraSonar Linux Agent Collector v%s\n", version)

	// Initialize random
	libagent.RandInit()

	// Initialize Helper
	libagent.GetHelper()

	// Set-up signal handler
	quit := make(chan bool)
	go libagent.SigHandler(quit)

	// Create Collector
	collector := libagent.NewCollector("linux", version)

	// Create Asset
	asset := libagent.NewAsset(collector)
	asset.Kind = "Linux"
	asset.Announce()

	// Create and plan checks
	checkSystem := libagent.Check{
		Key:             "system",
		Collector:       collector,
		Asset:           asset,
		IntervalEnv:     "CHECK_SYSTEM_INTERVAL",
		DefaultInterval: 300,
		NoCount:         false,
		SetTimestamp:    false,
		Fn:              CheckSystem,
	}
	go checkSystem.Plan(quit)

	checkDisk := libagent.Check{
		Key:             "disk",
		Collector:       collector,
		Asset:           asset,
		IntervalEnv:     "CHECK_DISK_INTERVAL",
		DefaultInterval: 300,
		NoCount:         false,
		SetTimestamp:    false,
		Fn:              CheckDisk,
	}
	go checkDisk.Plan(quit)

	checkMemInfo := libagent.Check{
		Key:             "memInfo",
		Collector:       collector,
		Asset:           asset,
		IntervalEnv:     "CHECK_MEMINFO_INTERVAL",
		DefaultInterval: 300,
		NoCount:         false,
		SetTimestamp:    false,
		Fn:              CheckMemInfo,
	}
	go checkMemInfo.Plan(quit)

	// Wait for quit
	<-quit
}
