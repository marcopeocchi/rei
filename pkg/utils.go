package pkg

import (
	"net/http"
	"os"
	"runtime"
)

func contentJson(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

func readCPUTempLINUX() (string, error) {
	if runtime.GOOS == "linux" {
		thermal_zone1, err := os.ReadFile("/sys/class/thermal/thermal_zone1/temp")
		if err != nil {
			thermal_zone2, err := os.ReadFile("/sys/class/thermal/thermal_zone2/temp")
			if err != nil {
				return "", err
			}
			return string(thermal_zone2), nil
		}
		return string(thermal_zone1), nil
	}
	return "0", nil
}
