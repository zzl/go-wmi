package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// 1A388F98-D4BA-11D1-8B09-00600806D9B6
var IID_ISWbemProperty = syscall.GUID{0x1A388F98, 0xD4BA, 0x11D1,
	[8]byte{0x8B, 0x09, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemProperty struct {
	ole.OleClient
}

func NewISWbemProperty(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemProperty {
	if pDisp == nil {
		return nil
	}
	p := &ISWbemProperty{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemPropertyFromVar(v ole.Variant) *ISWbemProperty {
	return NewISWbemProperty(v.IDispatch(), false, false)
}

func (this *ISWbemProperty) IID() *syscall.GUID {
	return &IID_ISWbemProperty
}

func (this *ISWbemProperty) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemProperty) Value() ole.Variant {
	retVal, _ := this.PropGet(0x00000000, nil)
	com.AddToScope(retVal)
	return *retVal
}

func (this *ISWbemProperty) SetValue(rhs *ole.Variant) {
	_ = this.PropPut(0x00000000, []interface{}{rhs})
}

func (this *ISWbemProperty) Name() string {
	retVal, _ := this.PropGet(0x00000001, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemProperty) IsLocal() bool {
	retVal, _ := this.PropGet(0x00000002, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemProperty) Origin() string {
	retVal, _ := this.PropGet(0x00000003, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemProperty) CIMType() int32 {
	retVal, _ := this.PropGet(0x00000004, nil)
	return retVal.LValVal()
}

func (this *ISWbemProperty) Qualifiers_() *ISWbemQualifierSet {
	retVal, _ := this.PropGet(0x00000005, nil)
	return NewISWbemQualifierSet(retVal.IDispatch(), false, true)
}

func (this *ISWbemProperty) IsArray() bool {
	retVal, _ := this.PropGet(0x00000006, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

