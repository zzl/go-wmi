package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemObjectEx = syscall.GUID{0xD6BDAFB2, 0x9435, 0x491F, 
	[8]byte{0xBB, 0x87, 0x6A, 0xA0, 0xF0, 0xBC, 0x31, 0xA2}}

type SWbemObjectEx struct {
	ISWbemObjectEx
}

func NewSWbemObjectEx(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemObjectEx {
	 if pDisp == nil {
		return nil;
	}
	p := &SWbemObjectEx{ISWbemObjectEx{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemObjectExFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemObjectEx {
	return NewSWbemObjectEx(v.IDispatch(), addRef, scoped)
}

func NewSWbemObjectExInstance(scoped bool) (*SWbemObjectEx, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemObjectEx, nil, 
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemObjectEx, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemObjectEx(p, false, scoped), nil
}

