package main

import (
	"fmt"
	"net"
	"os"
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/host"
)

func main() {
	Hostname, err := os.Hostname()
	checkNilError(err)

	cpuInfo, err := cpu.Info()
	checkNilError(err)

	cpuCounts, err := cpu.Counts(true)
	checkNilError(err)

	virtualMemory, err := mem.VirtualMemory()
	checkNilError(err)

	currentDisk, err := disk.Usage("/")
	checkNilError(err)

	networkInterfaces, err := net.Interfaces()
	checkNilError(err)
	//formatting network details to string to write it in a file
	var networkDetails string
	for _, iface := range networkInterfaces {
		networkDetails += fmt.Sprintf("Name: %s, MAC: %s\n", iface.Name, iface.HardwareAddr)
	}

	uptime, err := host.Uptime()
	checkNilError(err)
	// Format uptime into days, hours, minutes, and seconds
	days := uptime / (24 * 3600)
	hours := (uptime % (24 * 3600)) / 3600
	minutes := (uptime % 3600) / 60
	seconds := uptime % 60
	// Formatted uptime string
	uptimeStr := fmt.Sprintf("%d days, %02d hours, %02d minutes, %02d seconds", days, hours, minutes, seconds)

	fmt.Println("Operating Sytem:- ", runtime.GOOS)
	fmt.Println("Hostname is:- ", Hostname)
	fmt.Println("System Architecture:- ", runtime.GOARCH)
	fmt.Println("CPU INFO:- ", cpuInfo)
	fmt.Println("CPU COUNTS:- ", cpuCounts)
	fmt.Println("Virtual Memory Details:- ", virtualMemory)
	fmt.Println("Current Disk Details:- ", currentDisk)
	fmt.Println("Network Details:- ", networkDetails)
	fmt.Println("Host Uptime:- ", uptime)

	//formatitng the outputs
	output := fmt.Sprintf(
		"Operating System:- %s\nHostname is:- %s\nSystem Architecture:- %s\nCPU INFO:- %s\nCPU COUNTS:- %d\nVirtual Memory Details:- %s\nCurrent Disk Details:- %s\nNetwork Details:- %s\nHost Uptime:- %s\n",
		runtime.GOOS, Hostname, runtime.GOARCH, cpuInfo, cpuCounts, virtualMemory, currentDisk, networkDetails, uptimeStr,
	)

	err = os.WriteFile("SystemInfo.txt", []byte(output), 0644)
	checkNilError(err)
	fmt.Println("System info written successfully to system_info.txt!")
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
