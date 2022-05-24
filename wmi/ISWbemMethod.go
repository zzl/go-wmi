package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// 422E8E90-D955-11D1-8B09-00600806D9B6
var IID_ISWbemMethod = syscall.GUID{0x422E8E90, 0xD955, 0x11D1, 
	[8]byte{0x8B, 0x09, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemMethod struct {
	ole.OleClient
}

func NewISWbemMethod(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemMethod {
	 if pDisp == nil {
		return nil;
	}
	p := &ISWbemMethod{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemMethodFromVar(v ole.Variant) *ISWbemMethod {
	return NewISWbemMethod(v.IDispatch(), false, false)
}

func (this *ISWbemMethod) IID() *syscall.GUID {
	return &IID_ISWbemMethod
}

func (this *ISWbemMethod) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemMethod) Name() string {
	retVal, _ := this.PropGet(0x00000001, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemMethod) Origin() string {
	retVal, _ := this.PropGet(0x00000002, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemMethod) InParameters() *ISWbemObject {
	retVal, _ := this.PropGet(0x00000003, nil)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

func (this *ISWbemMethod) OutParameters() *ISWbemObject {
	retVal, _ := this.PropGet(0x00000004, nil)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

func (this *ISWbemMethod) Qualifiers_() *ISWbemQualifierSet {
	retVal, _ := this.PropGet(0x00000005, nil)
	return NewISWbemQualifierSet(retVal.IDispatch(), false, true)
}

