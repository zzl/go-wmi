package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// 26EE67BD-5804-11D2-8B4A-00600806D9B6
var IID_ISWbemPrivilege = syscall.GUID{0x26EE67BD, 0x5804, 0x11D2, 
	[8]byte{0x8B, 0x4A, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemPrivilege struct {
	ole.OleClient
}

func NewISWbemPrivilege(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemPrivilege {
	 if pDisp == nil {
		return nil;
	}
	p := &ISWbemPrivilege{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemPrivilegeFromVar(v ole.Variant) *ISWbemPrivilege {
	return NewISWbemPrivilege(v.IDispatch(), false, false)
}

func (this *ISWbemPrivilege) IID() *syscall.GUID {
	return &IID_ISWbemPrivilege
}

func (this *ISWbemPrivilege) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemPrivilege) IsEnabled() bool {
	retVal, _ := this.PropGet(0x00000000, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemPrivilege) SetIsEnabled(rhs bool)  {
	_ = this.PropPut(0x00000000, []interface{}{rhs})
}

func (this *ISWbemPrivilege) Name() string {
	retVal, _ := this.PropGet(0x00000001, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemPrivilege) DisplayName() string {
	retVal, _ := this.PropGet(0x00000002, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemPrivilege) Identifier() int32 {
	retVal, _ := this.PropGet(0x00000003, nil)
	return retVal.LValVal()
}

