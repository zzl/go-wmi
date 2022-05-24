package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemNamedValueSet = syscall.GUID{0x9AED384E, 0xCE8B, 0x11D1, 
	[8]byte{0x8B, 0x05, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type SWbemNamedValueSet struct {
	ISWbemNamedValueSet
}

func NewSWbemNamedValueSet(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemNamedValueSet {
	 if pDisp == nil {
		return nil;
	}
	p := &SWbemNamedValueSet{ISWbemNamedValueSet{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemNamedValueSetFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemNamedValueSet {
	return NewSWbemNamedValueSet(v.IDispatch(), addRef, scoped)
}

func NewSWbemNamedValueSetInstance(scoped bool) (*SWbemNamedValueSet, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemNamedValueSet, nil, 
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemNamedValueSet, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemNamedValueSet(p, false, scoped), nil
}

