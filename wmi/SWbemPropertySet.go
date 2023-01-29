package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemPropertySet = syscall.GUID{0x04B83D5C, 0x21AE, 0x11D2,
	[8]byte{0x8B, 0x33, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type SWbemPropertySet struct {
	ISWbemPropertySet
}

func NewSWbemPropertySet(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemPropertySet {
	if pDisp == nil {
		return nil
	}
	p := &SWbemPropertySet{ISWbemPropertySet{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemPropertySetFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemPropertySet {
	return NewSWbemPropertySet(v.IDispatch(), addRef, scoped)
}

func NewSWbemPropertySetInstance(scoped bool) (*SWbemPropertySet, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemPropertySet, nil,
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemPropertySet, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemPropertySet(p, false, scoped), nil
}

