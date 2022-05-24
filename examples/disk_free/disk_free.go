package main

import (
	"fmt"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-wmi/wmi"
	"math"
)

func main() {

	com.Initialize()
	defer com.NewScope().Leave()

	services := wmi.ConnectDefault()
	_ = services

	format := "%-6s%-12s%10s%10s%10s%8s\n"
	fmt.Printf(format, "Drive", "Label", "Size", "Free", "Used", "Used%")
	dispPartObjs := services.InstancesOf("Win32_LogicalDisk")
	dispPartObjs.ForEach(func(obj *wmi.ISWbemObject) bool {
		props := obj.Properties_()
		size := props.Get("Size").Int()
		free := props.Get("FreeSpace").Int()
		used := size - free
		usedPercent := math.Round(float64(used) * 100 / float64(size))
		fmt.Printf(format,
			props.Get("Name").String(),
			props.Get("VolumeName").String(),
			formatSize(size),
			formatSize(free),
			formatSize(used),
			fmt.Sprintf("%.0f%%", usedPercent),
		)
		return true
	})

}

func formatSize(byteCount int) string {
	return fmt.Sprintf("%d GB", byteCount/(1024*1024*1024))
}
