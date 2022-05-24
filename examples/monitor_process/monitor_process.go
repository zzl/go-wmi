package main

import (
	"fmt"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-wmi/wmi"
)

func main() {

	go monitorStart()
	go monitorStop()

	println("Press any key to quit.")
	win32.GetCh()
}

func monitorStart() {
	com.Initialize()
	defer com.Uninitialize()
	defer com.NewScope().Leave()

	service := wmi.ConnectDefault()

	wql := "SELECT * FROM Win32_ProcessStartTrace"
	eventSource := service.ExecNotificationQuery(wql)
	if eventSource == nil {
		// if the error is "access denied",
		// try running with administrative privileges.
		println(com.GetLastErrorInfo().String())
		return
	}
	for {
		com.WithScope(func() {
			event := eventSource.NextEvent()
			name := event.Get("ProcessName").String()
			pid := event.Get("ProcessID").Uint32()
			fmt.Printf("%s(%d) Started.\n", name, pid)
		})
	}
}

func monitorStop() {
	com.Initialize()
	defer com.Uninitialize()
	defer com.NewScope().Leave()

	service := wmi.ConnectDefault()

	wql := "SELECT * FROM Win32_ProcessStopTrace"
	eventSource := service.ExecNotificationQuery(wql)
	if eventSource == nil {
		println(com.GetLastErrorInfo().String())
		return
	}
	for {
		com.WithScope(func() {
			event := eventSource.NextEvent()
			name := event.Get("ProcessName").String()
			pid := event.Get("ProcessID").Uint32()
			exitCode := event.Get("ExitStatus").Int32()
			fmt.Printf("%s(%d) Exited: %v.\n", name, pid, exitCode)
		})
	}
}
