package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// 76A6415B-CB41-11D1-8B02-00600806D9B6
var IID_ISWbemLocator = syscall.GUID{0x76A6415B, 0xCB41, 0x11D1, 
	[8]byte{0x8B, 0x02, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemLocator struct {
	ole.OleClient
}

func NewISWbemLocator(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemLocator {
	 if pDisp == nil {
		return nil;
	}
	p := &ISWbemLocator{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemLocatorFromVar(v ole.Variant) *ISWbemLocator {
	return NewISWbemLocator(v.IDispatch(), false, false)
}

func (this *ISWbemLocator) IID() *syscall.GUID {
	return &IID_ISWbemLocator
}

func (this *ISWbemLocator) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

var ISWbemLocator_ConnectServer_OptArgs= []string{
	"strServer", "strNamespace", "strUser", "strPassword", 
	"strLocale", "strAuthority", "iSecurityFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemLocator) ConnectServer(optArgs ...interface{}) *ISWbemServices {
	optArgs = ole.ProcessOptArgs(ISWbemLocator_ConnectServer_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000001, nil, optArgs...)
	return NewISWbemServices(retVal.IDispatch(), false, true)
}

func (this *ISWbemLocator) Security_() *ISWbemSecurity {
	retVal, _ := this.PropGet(0x00000002, nil)
	return NewISWbemSecurity(retVal.IDispatch(), false, true)
}

