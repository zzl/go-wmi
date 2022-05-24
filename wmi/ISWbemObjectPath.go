package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// 5791BC27-CE9C-11D1-97BF-0000F81E849C
var IID_ISWbemObjectPath = syscall.GUID{0x5791BC27, 0xCE9C, 0x11D1, 
	[8]byte{0x97, 0xBF, 0x00, 0x00, 0xF8, 0x1E, 0x84, 0x9C}}

type ISWbemObjectPath struct {
	ole.OleClient
}

func NewISWbemObjectPath(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemObjectPath {
	 if pDisp == nil {
		return nil;
	}
	p := &ISWbemObjectPath{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemObjectPathFromVar(v ole.Variant) *ISWbemObjectPath {
	return NewISWbemObjectPath(v.IDispatch(), false, false)
}

func (this *ISWbemObjectPath) IID() *syscall.GUID {
	return &IID_ISWbemObjectPath
}

func (this *ISWbemObjectPath) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemObjectPath) Path() string {
	retVal, _ := this.PropGet(0x00000000, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemObjectPath) SetPath(rhs string)  {
	_ = this.PropPut(0x00000000, []interface{}{rhs})
}

func (this *ISWbemObjectPath) RelPath() string {
	retVal, _ := this.PropGet(0x00000001, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemObjectPath) SetRelPath(rhs string)  {
	_ = this.PropPut(0x00000001, []interface{}{rhs})
}

func (this *ISWbemObjectPath) Server() string {
	retVal, _ := this.PropGet(0x00000002, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemObjectPath) SetServer(rhs string)  {
	_ = this.PropPut(0x00000002, []interface{}{rhs})
}

func (this *ISWbemObjectPath) Namespace() string {
	retVal, _ := this.PropGet(0x00000003, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemObjectPath) SetNamespace(rhs string)  {
	_ = this.PropPut(0x00000003, []interface{}{rhs})
}

func (this *ISWbemObjectPath) ParentNamespace() string {
	retVal, _ := this.PropGet(0x00000004, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemObjectPath) DisplayName() string {
	retVal, _ := this.PropGet(0x00000005, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemObjectPath) SetDisplayName(rhs string)  {
	_ = this.PropPut(0x00000005, []interface{}{rhs})
}

func (this *ISWbemObjectPath) Class() string {
	retVal, _ := this.PropGet(0x00000006, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemObjectPath) SetClass(rhs string)  {
	_ = this.PropPut(0x00000006, []interface{}{rhs})
}

func (this *ISWbemObjectPath) IsClass() bool {
	retVal, _ := this.PropGet(0x00000007, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemObjectPath) SetAsClass()  {
	retVal, _ := this.Call(0x00000008, nil)
	_= retVal
}

func (this *ISWbemObjectPath) IsSingleton() bool {
	retVal, _ := this.PropGet(0x00000009, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemObjectPath) SetAsSingleton()  {
	retVal, _ := this.Call(0x0000000a, nil)
	_= retVal
}

func (this *ISWbemObjectPath) Keys() *ISWbemNamedValueSet {
	retVal, _ := this.PropGet(0x0000000b, nil)
	return NewISWbemNamedValueSet(retVal.IDispatch(), false, true)
}

func (this *ISWbemObjectPath) Security_() *ISWbemSecurity {
	retVal, _ := this.PropGet(0x0000000c, nil)
	return NewISWbemSecurity(retVal.IDispatch(), false, true)
}

func (this *ISWbemObjectPath) Locale() string {
	retVal, _ := this.PropGet(0x0000000d, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemObjectPath) SetLocale(rhs string)  {
	_ = this.PropPut(0x0000000d, []interface{}{rhs})
}

func (this *ISWbemObjectPath) Authority() string {
	retVal, _ := this.PropGet(0x0000000e, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemObjectPath) SetAuthority(rhs string)  {
	_ = this.PropPut(0x0000000e, []interface{}{rhs})
}

