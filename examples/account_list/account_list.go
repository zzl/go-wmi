package main

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-wmi/wmi"
)

func main() {

	com.Initialize()
	defer com.NewScope().Leave()

	service := wmi.ConnectDefault()

	accounts := service.ExecQuery("select * from Win32_UserAccount")
	accounts.ForEach(func(account *wmi.ISWbemObject) bool {
		println(account.Get("Name").String())
		return true
	})

}
