package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// 79B05932-D3B7-11D1-8B06-00600806D9B6
var IID_ISWbemQualifier = syscall.GUID{0x79B05932, 0xD3B7, 0x11D1, 
	[8]byte{0x8B, 0x06, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemQualifier struct {
	ole.OleClient
}

func NewISWbemQualifier(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemQualifier {
	 if pDisp == nil {
		return nil;
	}
	p := &ISWbemQualifier{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemQualifierFromVar(v ole.Variant) *ISWbemQualifier {
	return NewISWbemQualifier(v.IDispatch(), false, false)
}

func (this *ISWbemQualifier) IID() *syscall.GUID {
	return &IID_ISWbemQualifier
}

func (this *ISWbemQualifier) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemQualifier) Value() ole.Variant {
	retVal, _ := this.PropGet(0x00000000, nil)
	com.AddToScope(retVal)
	return *retVal
}

func (this *ISWbemQualifier) SetValue(rhs *ole.Variant)  {
	_ = this.PropPut(0x00000000, []interface{}{rhs})
}

func (this *ISWbemQualifier) Name() string {
	retVal, _ := this.PropGet(0x00000001, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemQualifier) IsLocal() bool {
	retVal, _ := this.PropGet(0x00000002, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemQualifier) PropagatesToSubclass() bool {
	retVal, _ := this.PropGet(0x00000003, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemQualifier) SetPropagatesToSubclass(rhs bool)  {
	_ = this.PropPut(0x00000003, []interface{}{rhs})
}

func (this *ISWbemQualifier) PropagatesToInstance() bool {
	retVal, _ := this.PropGet(0x00000004, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemQualifier) SetPropagatesToInstance(rhs bool)  {
	_ = this.PropPut(0x00000004, []interface{}{rhs})
}

func (this *ISWbemQualifier) IsOverridable() bool {
	retVal, _ := this.PropGet(0x00000005, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemQualifier) SetIsOverridable(rhs bool)  {
	_ = this.PropPut(0x00000005, []interface{}{rhs})
}

func (this *ISWbemQualifier) IsAmended() bool {
	retVal, _ := this.PropGet(0x00000006, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

