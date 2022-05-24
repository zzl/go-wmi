package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemLastError = syscall.GUID{0xC2FEEEAC, 0xCFCD, 0x11D1, 
	[8]byte{0x8B, 0x05, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type SWbemLastError struct {
	ISWbemLastError
}

func NewSWbemLastError(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemLastError {
	 if pDisp == nil {
		return nil;
	}
	p := &SWbemLastError{ISWbemLastError{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemLastErrorFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemLastError {
	return NewSWbemLastError(v.IDispatch(), addRef, scoped)
}

func NewSWbemLastErrorInstance(scoped bool) (*SWbemLastError, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemLastError, nil, 
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemLastError, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemLastError(p, false, scoped), nil
}

