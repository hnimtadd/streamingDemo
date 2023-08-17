package utils

import (
	"cameraClient/entities"
	"runtime"
)

func GetHardwareInfo() (entities.HardwareInfo, error) {
	info := entities.HardwareInfo{
		Os:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}
	return info, nil
}
