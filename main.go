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

	//
	// Create and plan checks
	//

	// system check
	checkSystem := libagent.Check{
		Key:          "system",
		Collector:    collector,
		Asset:        asset,
		IntervalEnv:  "CHECK_SYSTEM_INTERVAL",
		NoCount:      false,
		SetTimestamp: false,
		Fn:           CheckSystem,
	}
	go checkSystem.Plan(quit)

	// syslog check
	checkSyslog := libagent.Check{
		Key:          "syslog",
		Collector:    collector,
		Asset:        asset,
		IntervalEnv:  "CHECK_SYSLOG_INTERVAL",
		NoCount:      false,
		SetTimestamp: false,
		Fn:           CheckSyslog,
	}
	go checkSyslog.Plan(quit)

	// Wait for quit
	<-quit
}
