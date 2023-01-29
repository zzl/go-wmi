package wmi

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
)

// 76A64164-CB41-11D1-8B02-00600806D9B6
var IID_ISWbemNamedValue = syscall.GUID{0x76A64164, 0xCB41, 0x11D1,
	[8]byte{0x8B, 0x02, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemNamedValue struct {
	ole.OleClient
}

func NewISWbemNamedValue(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemNamedValue {
	if pDisp == nil {
		return nil
	}
	p := &ISWbemNamedValue{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemNamedValueFromVar(v ole.Variant) *ISWbemNamedValue {
	return NewISWbemNamedValue(v.IDispatch(), false, false)
}

func (this *ISWbemNamedValue) IID() *syscall.GUID {
	return &IID_ISWbemNamedValue
}

func (this *ISWbemNamedValue) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemNamedValue) Value() ole.Variant {
	retVal, _ := this.PropGet(0x00000000, nil)
	com.AddToScope(retVal)
	return *retVal
}

func (this *ISWbemNamedValue) SetValue(rhs *ole.Variant) {
	_ = this.PropPut(0x00000000, []interface{}{rhs})
}

func (this *ISWbemNamedValue) Name() string {
	retVal, _ := this.PropGet(0x00000002, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}
