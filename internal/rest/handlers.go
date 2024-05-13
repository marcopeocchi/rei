package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/marcopeocchi/rei/internal/config"
	"github.com/marcopeocchi/rei/internal/models"
	"github.com/marcopeocchi/rei/internal/utils"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func top() *models.SystemTop {
	var (
		cpu, _    = cpu.Info()
		host, _   = host.Info()
		vmstat, _ = mem.VirtualMemory()
	)

	res := models.SystemTop{
		Hostname: host.Hostname,
		OS:       host.OS,
		Platform: host.Platform,
		Uptime:   host.Uptime,
		MemFree:  vmstat.Available,
		MemTotal: vmstat.Total,
	}

	if runtime.GOOS == "linux" {
		res.CPU = cpu[0].ModelName
		res.CoreCount = cpu[0].Cores + 1
	}

	return &res
}

func Top(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := top()
	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TopFmt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		vmstat, _    = mem.VirtualMemory()
		cpuUsage, _  = cpu.Percent(time.Millisecond*500, false)
		diskUsage, _ = disk.Usage("/")
	)

	fmt.Fprintf(
		w,
		"CPU: %0.f%% MEM: %.0f%% HDD: %.0f%%",
		cpuUsage[0],
		vmstat.UsedPercent,
		diskUsage.UsedPercent,
	)
}

func Temp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	temp, _ := utils.ReadCPUTemp()

	err := json.NewEncoder(w).Encode(models.SystemTemp{
		CPUTemp: temp,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Config(c *config.SafeConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		c.JsonEncoder(w)
	}
}
