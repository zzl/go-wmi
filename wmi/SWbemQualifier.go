package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemQualifier = syscall.GUID{0x04B83D5F, 0x21AE, 0x11D2,
	[8]byte{0x8B, 0x33, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type SWbemQualifier struct {
	ISWbemQualifier
}

func NewSWbemQualifier(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemQualifier {
	if pDisp == nil {
		return nil
	}
	p := &SWbemQualifier{ISWbemQualifier{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemQualifierFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemQualifier {
	return NewSWbemQualifier(v.IDispatch(), addRef, scoped)
}

func NewSWbemQualifierInstance(scoped bool) (*SWbemQualifier, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemQualifier, nil,
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemQualifier, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemQualifier(p, false, scoped), nil
}

