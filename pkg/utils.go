package pkg

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

// Shorthand to set the content type to application/json
func contentJson(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

// Reads the CPU package temperature from /sys/class/themrmal
// Due to this, linux is the only supported platform.
// If the first reading from x86_pkg_temp fails fallback to acpiz one.
// If the system is not supported returns error.
func readCPUTempLINUX() (string, error) {
	if runtime.GOOS == "linux" {
		thermal_zone2, err := os.ReadFile("/sys/class/thermal/thermal_zone2/temp")
		if err != nil {
			thermal_zone1, err := os.ReadFile("/sys/class/thermal/thermal_zone1/temp")
			if err != nil {
				return "", err
			}
			return string(thermal_zone1), nil
		}
		return string(thermal_zone2), nil
	}
	return "", fmt.Errorf("system %s is not supported", runtime.GOOS)
}
