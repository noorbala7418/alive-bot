package main

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type Process struct {
	pid int
	cpu float64
}

func Cpuusage() int {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	processes := make([]*Process, 0)
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		// log.Println(len(ft), ft)
		pid, err := strconv.Atoi(ft[1])
		if err != nil {
			continue
		}
		cpu, err := strconv.ParseFloat(ft[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		processes = append(processes, &Process{pid, cpu})
	}
	cpusum := 0
	for _, p := range processes {
		// log.Println("Process ", p.pid, " takes ", p.cpu, " % of the CPU")
		cpusum += int(p.cpu)
	}
	return cpusum
}

func Memoryusage() int {
	meminfo, _ := exec.Command("cat", "/proc/meminfo").Output()
	memoryraw := strings.Split(string(meminfo), "\n")

	totalmemory, _ := strconv.ParseInt(strings.FieldsFunc(memoryraw[0], func(r rune) bool { return r == ' ' })[1], 10, 64)
	usedmemory, _ := strconv.ParseInt(strings.FieldsFunc(memoryraw[1], func(r rune) bool { return r == ' ' })[1], 10, 64)
	result := int(usedmemory * 100 / totalmemory)

	return result
}

func rebootSystem() {
	syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART2)
}

func systemUptime() string {
	cmd, _ := exec.Command("bash", "-c", "uptime").Output()
	return strings.Split(string(cmd), ",")[0]
}
