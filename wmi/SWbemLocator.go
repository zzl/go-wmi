package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemLocator = syscall.GUID{0x76A64158, 0xCB41, 0x11D1,
	[8]byte{0x8B, 0x02, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type SWbemLocator struct {
	ISWbemLocator
}

func NewSWbemLocator(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemLocator {
	if pDisp == nil {
		return nil
	}
	p := &SWbemLocator{ISWbemLocator{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemLocatorFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemLocator {
	return NewSWbemLocator(v.IDispatch(), addRef, scoped)
}

func NewSWbemLocatorInstance(scoped bool) (*SWbemLocator, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemLocator, nil,
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemLocator, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemLocator(p, false, scoped), nil
}

