package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemRefreshableItem = syscall.GUID{0x8C6854BC, 0xDE4B, 0x11D3, 
	[8]byte{0xB3, 0x90, 0x00, 0x10, 0x5A, 0x1F, 0x47, 0x3A}}

type SWbemRefreshableItem struct {
	ISWbemRefreshableItem
}

func NewSWbemRefreshableItem(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemRefreshableItem {
	 if pDisp == nil {
		return nil;
	}
	p := &SWbemRefreshableItem{ISWbemRefreshableItem{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemRefreshableItemFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemRefreshableItem {
	return NewSWbemRefreshableItem(v.IDispatch(), addRef, scoped)
}

func NewSWbemRefreshableItemInstance(scoped bool) (*SWbemRefreshableItem, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemRefreshableItem, nil, 
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemRefreshableItem, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemRefreshableItem(p, false, scoped), nil
}

