package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// 75718C9F-F029-11D1-A1AC-00C04FB6C223
var IID_ISWbemSink = syscall.GUID{0x75718C9F, 0xF029, 0x11D1, 
	[8]byte{0xA1, 0xAC, 0x00, 0xC0, 0x4F, 0xB6, 0xC2, 0x23}}

type ISWbemSink struct {
	ole.OleClient
}

func NewISWbemSink(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemSink {
	 if pDisp == nil {
		return nil;
	}
	p := &ISWbemSink{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemSinkFromVar(v ole.Variant) *ISWbemSink {
	return NewISWbemSink(v.IDispatch(), false, false)
}

func (this *ISWbemSink) IID() *syscall.GUID {
	return &IID_ISWbemSink
}

func (this *ISWbemSink) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemSink) Cancel()  {
	retVal, _ := this.Call(0x00000001, nil)
	_= retVal
}

