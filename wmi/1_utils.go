package wmi

import (
	"github.com/zzl/go-com/ole"
	"log"
)

func ConnectDefault() *ISWbemServices {
	locator, err := NewSWbemLocatorInstance(false)
	if err != nil {
		log.Fatal(err)
	}
	defer locator.Release()
	service := locator.ConnectServer(ole.NamedArgs{
		"strServer":    ".",
		"strNamespace": "root\\cimv2",
	})
	service.Security_().SetImpersonationLevel(3)
	return service
}

func (this *ISWbemPropertySet) Get(name string) *ole.Variant {
	v := this.Item(name).Value()
	return &v
}

func (this *ISWbemObject) Get(name string) *ole.Variant {
	return this.Properties_().Get(name)
}

func (this *SWbemLastError) Get(name string) *ole.Variant {
	return this.Properties_().Get(name)
}
