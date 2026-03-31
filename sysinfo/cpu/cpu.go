package cpu

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type Info struct {
	Cores       int
	Temperature int
}

func Stat() (Info, error) {
	temperature, err := temperature()
	if err != nil {
		return Info{}, err
	}
	return Info{
		Cores:       runtime.NumCPU(),
		Temperature: temperature,
	}, nil
}

func temperature() (int, error) {
	files, err := filepath.Glob("/sys/class/thermal/thermal_zone*/temp")
	if err != nil {
		return 0, err
	}
	if len(files) == 0 {
		return 0, fmt.Errorf("no thermal zone files found")
	}
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return 0, err
		}
		degrees := strings.TrimSuffix(string(content), "\n")
		deg, err := strconv.Atoi(degrees)
		if err != nil {
			return 0, err
		}
		return deg / 1000, nil
	}
	return 0, fmt.Errorf("no thermal zone files found")
}
