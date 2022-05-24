package main

import (
	"fmt"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-wmi/wmi"
	"sort"
)

func main() {

	com.Initialize()
	defer com.NewScope().Leave()

	service := wmi.ConnectDefault()

	type ProcessInfo struct {
		Name string
		PID  uint32
		Mem  uint64
	}
	var processeInfos []ProcessInfo

	processes := service.ExecQuery("select * from Win32_Process")
	processes.ForEach(func(process *wmi.ISWbemObject) bool {
		props := process.Properties_()
		processeInfos = append(processeInfos, ProcessInfo{
			props.Get("Name").String(),
			props.Get("ProcessID").Uint32(),
			props.Get("WorkingSetSize").Uint64(),
		})
		return true
	})

	sort.Slice(processeInfos, func(i, j int) bool {
		return processeInfos[i].Name < processeInfos[j].Name
	})

	fmt.Println("Name                                 PID           Working set")
	fmt.Println("---------------------------------------------------------------")
	format := "%-30s%10d%20d K\n"
	for _, info := range processeInfos {
		fmt.Printf(format, info.Name, info.PID, info.Mem/1024)
	}
}
