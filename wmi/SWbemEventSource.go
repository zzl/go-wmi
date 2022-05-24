package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemEventSource = syscall.GUID{0x04B83D58, 0x21AE, 0x11D2, 
	[8]byte{0x8B, 0x33, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type SWbemEventSource struct {
	ISWbemEventSource
}

func NewSWbemEventSource(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemEventSource {
	 if pDisp == nil {
		return nil;
	}
	p := &SWbemEventSource{ISWbemEventSource{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemEventSourceFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemEventSource {
	return NewSWbemEventSource(v.IDispatch(), addRef, scoped)
}

func NewSWbemEventSourceInstance(scoped bool) (*SWbemEventSource, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemEventSource, nil, 
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemEventSource, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemEventSource(p, false, scoped), nil
}

