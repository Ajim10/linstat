package memory

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Info struct {
	Total     uint64
	Free      uint64
	Available uint64
	Buffers   uint64
	Cached    uint64
}

func MemStat() (Info, error) {
	file, err := os.OpenFile("/proc/meminfo", os.O_RDONLY, 0444)
	if err != nil {
		return Info{}, err
	}
	defer file.Close()

	var info Info
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		key, value, found := strings.Cut(line, ":")
		if !found {
			continue
		}
		value = strings.TrimSpace(value)
		value = strings.TrimSuffix(value, " kB")

		t, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			continue
		}

		switch key {
		case "MemTotal":
			info.Total = t * 1024
		case "MemFree":
			info.Free = t * 1024
		case "MemAvailable":
			info.Available = t * 1024
		case "Buffers":
			info.Buffers = t * 1024
		case "Cached":
			info.Cached = t * 1024
		}
	}

	if err := scanner.Err(); err != nil {
		return Info{}, err
	}

	if info.Total == 0 {
		return Info{}, fmt.Errorf("MemTotal not found")
	}

	return info, nil
}
