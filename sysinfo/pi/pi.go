package pi

import "os"

type Info struct {
	Model string
}

func Stat() (Info, error) {
	model, err := os.ReadFile("/proc/device-tree/model")
	if err != nil {
		return Info{}, err
	}
	return Info{Model: string(model)}, nil
}
