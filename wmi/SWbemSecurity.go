package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemSecurity = syscall.GUID{0xB54D66E9, 0x2287, 0x11D2,
	[8]byte{0x8B, 0x33, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type SWbemSecurity struct {
	ISWbemSecurity
}

func NewSWbemSecurity(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemSecurity {
	if pDisp == nil {
		return nil
	}
	p := &SWbemSecurity{ISWbemSecurity{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemSecurityFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemSecurity {
	return NewSWbemSecurity(v.IDispatch(), addRef, scoped)
}

func NewSWbemSecurityInstance(scoped bool) (*SWbemSecurity, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemSecurity, nil,
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemSecurity, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemSecurity(p, false, scoped), nil
}

