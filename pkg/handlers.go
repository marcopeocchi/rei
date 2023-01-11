package pkg

import (
	"log"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// HTTP handler function for retrieving "top-like" data of the system
func handleTop(w http.ResponseWriter, r *http.Request) {
	contentJson(&w)
	cpu, _ := cpu.Info()
	host, _ := host.Info()
	vmstat, _ := mem.VirtualMemory()

	res, err := json.Marshal(SystemTop{
		CPU:       cpu[0].ModelName,
		CoreCount: cpu[0].Cores + 1,
		Hostname:  host.Hostname,
		OS:        host.OS,
		Platform:  host.Platform,
		Uptime:    host.Uptime,
		RAMFree:   vmstat.Available,
	})

	if err != nil {
		log.Fatalln(err)
	}

	w.Write(res)
}

// HTTP handler function for retrieving cpu package temperature
func handleTemp(w http.ResponseWriter, r *http.Request) {
	contentJson(&w)
	temp, _ := readCPUTempLINUX()
	res, err := json.Marshal(SystemTemp{
		CPUTemp: temp,
	})

	if err != nil {
		log.Fatalln(err)
	}

	w.Write(res)
}

// HTTP handler function for retrieving services
func handleConfig(w http.ResponseWriter, r *http.Request) {
	contentJson(&w)
	w.Write(config.Json())
}
