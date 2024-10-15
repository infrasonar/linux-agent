package main

import (
	"bufio"
	"log"
	"os"

	"github.com/influxdata/go-syslog/v3/rfc3164"
)

func CheckSyslog() (map[string][]map[string]any, error) {
	state := map[string][]map[string]any{}
	items := []map[string]any{}

	fn := os.Getenv("SYSLOG_PATH")
	if fn == "" {
		fn = "/var/log/syslog"
	}

	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	line, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	p := rfc3164.NewParser(rfc3164.WithBestEffort(), rfc3164.WithRFC3339())
	m, e := p.Parse(line)
	log.Fatal(m, e)

	// item := map[string]any{
	// 	"name":  "uptime",
	// 	"idle":  uptime.Idle,
	// 	"total": uptime.Total,
	// }

	state["syslog"] = items
	return state, nil
}
