package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"encoding/json"
)

type stat struct {
	Key string
    Typeperf_name string
    Value  string
	Vtype string
}

func newStat(report string, key string, stat_index int) *stat {
	// split report and return new stat
	lines := strings.Split(report, "\n")
	header := strings.Split(lines[1], ",")
	values := strings.Split(lines[2], ",")
	s := stat{
		Key: key,
		Typeperf_name: strings.Trim(header[stat_index+1], "\r"),
		Value: strings.Trim(values[stat_index+1], "\r"),
		Vtype: "float",
	}
    return &s
}

func main() {
	COMMANDS := map[string]string{
		"cpu": "\\Processor(_total)\\% Processor Time",
		"mem_used": "\\Memory\\Committed Bytes",
		"mem_free": "\\Memory\\Available Bytes",
		"disk_read": "\\PhysicalDisk(_Total)\\Avg. Disk Bytes/Read",
		"disk_write": "\\PhysicalDisk(_Total)\\Avg. Disk Bytes/Write",
	}
	out, err := exec.Command(
		"typeperf.exe",
		COMMANDS["cpu"], COMMANDS["mem_used"], COMMANDS["disk_read"], COMMANDS["disk_write"],
		"-sc",
		"1",
	).Output()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("-- report:\n", string(out))

	for i, s := range []string{"cpu", "mem_used", "disk_read", "disk_write"} {
		stat := newStat(string(out), s, i)
		json_stat, _ := json.Marshal(stat)
		fmt.Println(string(json_stat))
	}
}
