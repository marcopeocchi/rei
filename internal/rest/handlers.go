package rest

import (
	"net/http"
	"valeera/m/internal/config"
	"valeera/m/internal/utils"

	"github.com/goccy/go-json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// HTTP handler function for retrieving "top-like" data of the system
func Top(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cpu, _ := cpu.Info()
	host, _ := host.Info()
	vmstat, _ := mem.VirtualMemory()

	err := json.NewEncoder(w).EncodeContext(r.Context(), SystemTop{
		CPU:       cpu[0].ModelName,
		CoreCount: cpu[0].Cores + 1,
		Hostname:  host.Hostname,
		OS:        host.OS,
		Platform:  host.Platform,
		Uptime:    host.Uptime,
		RAMFree:   vmstat.Available,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HTTP handler function for retrieving cpu package temperature
func Temp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	temp, _ := utils.ReadCPUTempLINUX()

	err := json.NewEncoder(w).EncodeContext(r.Context(), SystemTemp{
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
