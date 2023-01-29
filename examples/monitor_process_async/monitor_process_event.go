package main

import (
	"fmt"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-wmi/wmi"
)

func main() {

	defer com.Initialize().Uninitialize()
	defer com.NewScope().Leave()

	service := wmi.ConnectDefault()

	//
	startEventSink, _ := wmi.NewSWbemSinkInstance(true)
	startEventSink.RegisterEventHandlers(wmi.ISWbemSinkEventsHandlers{
		OnObjectReady: func(event *wmi.ISWbemObject, context *wmi.ISWbemNamedValueSet) {
			defer com.NewScope().Leave()
			name := event.Get("ProcessName").String()
			pid := event.Get("ProcessID").Uint32()
			fmt.Printf("%s(%d) Started.\n", name, pid)
		},
	})

	wql := "SELECT * FROM Win32_ProcessStartTrace"
	service.ExecNotificationQueryAsync(startEventSink.IDispatch, wql)

	//
	stopEventSink, _ := wmi.NewSWbemSinkInstance(true)
	stopEventSink.RegisterEventHandlers(wmi.ISWbemSinkEventsHandlers{
		OnObjectReady: func(event *wmi.ISWbemObject, context *wmi.ISWbemNamedValueSet) {
			defer com.NewScope().Leave()
			name := event.Get("ProcessName").String()
			pid := event.Get("ProcessID").Uint32()
			exitCode := event.Get("ExitStatus").Int32()
			fmt.Printf("%s(%d) Exited: %v.\n", name, pid, exitCode)
		},
	})

	wql = "SELECT * FROM Win32_ProcessStopTrace"
	service.ExecNotificationQueryAsync(stopEventSink.IDispatch, wql)

	//
	tid := com.GetContext().TID
	go func() {
		println("Press any key to quit.")
		win32.GetCh()
		win32.PostThreadMessage(tid, win32.WM_QUIT, 0, 0)
	}()

	com.MessageLoop()
}
