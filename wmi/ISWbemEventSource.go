package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// 27D54D92-0EBE-11D2-8B22-00600806D9B6
var IID_ISWbemEventSource = syscall.GUID{0x27D54D92, 0x0EBE, 0x11D2,
	[8]byte{0x8B, 0x22, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemEventSource struct {
	ole.OleClient
}

func NewISWbemEventSource(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemEventSource {
	if pDisp == nil {
		return nil
	}
	p := &ISWbemEventSource{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemEventSourceFromVar(v ole.Variant) *ISWbemEventSource {
	return NewISWbemEventSource(v.IDispatch(), false, false)
}

func (this *ISWbemEventSource) IID() *syscall.GUID {
	return &IID_ISWbemEventSource
}

func (this *ISWbemEventSource) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

var ISWbemEventSource_NextEvent_OptArgs = []string{
	"iTimeoutMs",
}

func (this *ISWbemEventSource) NextEvent(optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemEventSource_NextEvent_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000001, nil, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

func (this *ISWbemEventSource) Security_() *ISWbemSecurity {
	retVal, _ := this.PropGet(0x00000002, nil)
	return NewISWbemSecurity(retVal.IDispatch(), false, true)
}

