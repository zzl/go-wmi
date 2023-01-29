package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemServicesEx = syscall.GUID{0x62E522DC, 0x8CF3, 0x40A8,
	[8]byte{0x8B, 0x2E, 0x37, 0xD5, 0x95, 0x65, 0x1E, 0x40}}

type SWbemServicesEx struct {
	ISWbemServicesEx
}

func NewSWbemServicesEx(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemServicesEx {
	if pDisp == nil {
		return nil
	}
	p := &SWbemServicesEx{ISWbemServicesEx{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemServicesExFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemServicesEx {
	return NewSWbemServicesEx(v.IDispatch(), addRef, scoped)
}

func NewSWbemServicesExInstance(scoped bool) (*SWbemServicesEx, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemServicesEx, nil,
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemServicesEx, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemServicesEx(p, false, scoped), nil
}

