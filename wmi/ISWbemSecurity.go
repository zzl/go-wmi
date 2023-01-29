package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// B54D66E6-2287-11D2-8B33-00600806D9B6
var IID_ISWbemSecurity = syscall.GUID{0xB54D66E6, 0x2287, 0x11D2,
	[8]byte{0x8B, 0x33, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemSecurity struct {
	ole.OleClient
}

func NewISWbemSecurity(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemSecurity {
	if pDisp == nil {
		return nil
	}
	p := &ISWbemSecurity{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemSecurityFromVar(v ole.Variant) *ISWbemSecurity {
	return NewISWbemSecurity(v.IDispatch(), false, false)
}

func (this *ISWbemSecurity) IID() *syscall.GUID {
	return &IID_ISWbemSecurity
}

func (this *ISWbemSecurity) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemSecurity) ImpersonationLevel() int32 {
	retVal, _ := this.PropGet(0x00000001, nil)
	return retVal.LValVal()
}

func (this *ISWbemSecurity) SetImpersonationLevel(rhs int32) {
	_ = this.PropPut(0x00000001, []interface{}{rhs})
}

func (this *ISWbemSecurity) AuthenticationLevel() int32 {
	retVal, _ := this.PropGet(0x00000002, nil)
	return retVal.LValVal()
}

func (this *ISWbemSecurity) SetAuthenticationLevel(rhs int32) {
	_ = this.PropPut(0x00000002, []interface{}{rhs})
}

func (this *ISWbemSecurity) Privileges() *ISWbemPrivilegeSet {
	retVal, _ := this.PropGet(0x00000003, nil)
	return NewISWbemPrivilegeSet(retVal.IDispatch(), false, true)
}

