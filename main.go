package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type brejax_Config struct {
	Port      int               `json:"port"`
	Endpoints map[string]string `json:"endpoints"`
}

type brejax_SystemStats struct {
	CPUUsage    []float64 `json:"cpu_usage"`
	MemoryUsage struct {
		Total       uint64  `json:"total"`
		Used        uint64  `json:"used"`
		UsedPercent float64 `json:"used_percent"`
	} `json:"memory_usage"`
	DiskUsage struct {
		Total       uint64  `json:"total"`
		Used        uint64  `json:"used"`
		UsedPercent float64 `json:"used_percent"`
	} `json:"disk_usage"`
	NetworkStats []net.IOCountersStat `json:"network_stats"`
	SystemLoad   load.AvgStat         `json:"system_load"`
}

func brejax_LoadConfig(filePath string) (brejax_Config, error) {
	var config brejax_Config
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}

func brejax_GetStats() (brejax_SystemStats, error) {
	var stats brejax_SystemStats

	// CPU Usage
	cpuUsage, err := cpu.Percent(0, false)
	if err != nil {
		return stats, err
	}
	stats.CPUUsage = cpuUsage

	// Memory Usage
	memStats, err := mem.VirtualMemory()
	if err != nil {
		return stats, err
	}
	stats.MemoryUsage.Total = memStats.Total
	stats.MemoryUsage.Used = memStats.Used
	stats.MemoryUsage.UsedPercent = memStats.UsedPercent

	// Disk Usage
	diskStats, err := disk.Usage("/")
	if err != nil {
		return stats, err
	}
	stats.DiskUsage.Total = diskStats.Total
	stats.DiskUsage.Used = diskStats.Used
	stats.DiskUsage.UsedPercent = diskStats.UsedPercent

	// Network Stats
	netStats, err := net.IOCounters(true)
	if err != nil {
		return stats, err
	}
	stats.NetworkStats = netStats

	// System Load
	systemLoad, err := load.Avg()
	if err != nil {
		return stats, err
	}
	stats.SystemLoad = *systemLoad

	return stats, nil
}

func brejax_StatsHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := brejax_GetStats()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving stats: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func brejax_CreateHandler(dataFunc func() (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := dataFunc()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving data: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

func main() {
	// Load configuration
	config, err := brejax_LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Set up routes
	http.HandleFunc(config.Endpoints["stats"], brejax_StatsHandler)
	http.HandleFunc(config.Endpoints["cpu"], brejax_CreateHandler(func() (interface{}, error) {
		return cpu.Percent(0, false)
	}))
	http.HandleFunc(config.Endpoints["memory"], brejax_CreateHandler(func() (interface{}, error) {
		return mem.VirtualMemory()
	}))
	http.HandleFunc(config.Endpoints["disk"], brejax_CreateHandler(func() (interface{}, error) {
		return disk.Usage("/")
	}))
	http.HandleFunc(config.Endpoints["network"], brejax_CreateHandler(func() (interface{}, error) {
		return net.IOCounters(true)
	}))
	http.HandleFunc(config.Endpoints["load"], brejax_CreateHandler(func() (interface{}, error) {
		return load.Avg()
	}))

	// Start the server
	port := config.Port
	log.Printf("Server is running on port %d...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
