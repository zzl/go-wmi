package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemDateTime = syscall.GUID{0x47DFBE54, 0xCF76, 0x11D3, 
	[8]byte{0xB3, 0x8F, 0x00, 0x10, 0x5A, 0x1F, 0x47, 0x3A}}

type SWbemDateTime struct {
	ISWbemDateTime
}

func NewSWbemDateTime(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemDateTime {
	 if pDisp == nil {
		return nil;
	}
	p := &SWbemDateTime{ISWbemDateTime{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemDateTimeFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemDateTime {
	return NewSWbemDateTime(v.IDispatch(), addRef, scoped)
}

func NewSWbemDateTimeInstance(scoped bool) (*SWbemDateTime, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemDateTime, nil, 
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemDateTime, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemDateTime(p, false, scoped), nil
}

