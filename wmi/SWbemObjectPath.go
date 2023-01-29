package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemObjectPath = syscall.GUID{0x5791BC26, 0xCE9C, 0x11D1,
	[8]byte{0x97, 0xBF, 0x00, 0x00, 0xF8, 0x1E, 0x84, 0x9C}}

type SWbemObjectPath struct {
	ISWbemObjectPath
}

func NewSWbemObjectPath(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemObjectPath {
	if pDisp == nil {
		return nil
	}
	p := &SWbemObjectPath{ISWbemObjectPath{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemObjectPathFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemObjectPath {
	return NewSWbemObjectPath(v.IDispatch(), addRef, scoped)
}

func NewSWbemObjectPathInstance(scoped bool) (*SWbemObjectPath, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemObjectPath, nil,
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemObjectPath, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemObjectPath(p, false, scoped), nil
}

