package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemPrivilegeSet = syscall.GUID{0x26EE67BE, 0x5804, 0x11D2, 
	[8]byte{0x8B, 0x4A, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type SWbemPrivilegeSet struct {
	ISWbemPrivilegeSet
}

func NewSWbemPrivilegeSet(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemPrivilegeSet {
	 if pDisp == nil {
		return nil;
	}
	p := &SWbemPrivilegeSet{ISWbemPrivilegeSet{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemPrivilegeSetFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemPrivilegeSet {
	return NewSWbemPrivilegeSet(v.IDispatch(), addRef, scoped)
}

func NewSWbemPrivilegeSetInstance(scoped bool) (*SWbemPrivilegeSet, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemPrivilegeSet, nil, 
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemPrivilegeSet, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemPrivilegeSet(p, false, scoped), nil
}

