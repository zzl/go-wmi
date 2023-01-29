package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemRefresher = syscall.GUID{0xD269BF5C, 0xD9C1, 0x11D3,
	[8]byte{0xB3, 0x8F, 0x00, 0x10, 0x5A, 0x1F, 0x47, 0x3A}}

type SWbemRefresher struct {
	ISWbemRefresher
}

func NewSWbemRefresher(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemRefresher {
	if pDisp == nil {
		return nil
	}
	p := &SWbemRefresher{ISWbemRefresher{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemRefresherFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemRefresher {
	return NewSWbemRefresher(v.IDispatch(), addRef, scoped)
}

func NewSWbemRefresherInstance(scoped bool) (*SWbemRefresher, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemRefresher, nil,
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemRefresher, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemRefresher(p, false, scoped), nil
}

