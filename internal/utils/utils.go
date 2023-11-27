package utils

import (
	"fmt"
	"os"
	"runtime"
)

// Reads the CPU package temperature from /sys/class/themrmal
// Due to this, linux is the only supported platform.
func ReadCPUTemp() (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("system %s is not supported", runtime.GOOS)
	}

	thermal_zone0, err := os.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return "", err
	}

	return string(thermal_zone0), nil
}
