package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemMethod = syscall.GUID{0x04B83D5B, 0x21AE, 0x11D2,
	[8]byte{0x8B, 0x33, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type SWbemMethod struct {
	ISWbemMethod
}

func NewSWbemMethod(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemMethod {
	if pDisp == nil {
		return nil
	}
	p := &SWbemMethod{ISWbemMethod{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemMethodFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemMethod {
	return NewSWbemMethod(v.IDispatch(), addRef, scoped)
}

func NewSWbemMethodInstance(scoped bool) (*SWbemMethod, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemMethod, nil,
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemMethod, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemMethod(p, false, scoped), nil
}

