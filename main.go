package main

import (
	"fmt"
	"log"
	"os/exec"
)

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
	fmt.Println(string(out))
}
