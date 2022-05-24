package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

var CLSID_SWbemSink = syscall.GUID{0x75718C9A, 0xF029, 0x11D1, 
	[8]byte{0xA1, 0xAC, 0x00, 0xC0, 0x4F, 0xB6, 0xC2, 0x23}}

type SWbemSink struct {
	ISWbemSink
}

func NewSWbemSink(pDisp *win32.IDispatch, addRef bool, scoped bool) *SWbemSink {
	 if pDisp == nil {
		return nil;
	}
	p := &SWbemSink{ISWbemSink{ole.OleClient{pDisp}}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewSWbemSinkFromVar(v ole.Variant, addRef bool, scoped bool) *SWbemSink {
	return NewSWbemSink(v.IDispatch(), addRef, scoped)
}

func NewSWbemSinkInstance(scoped bool) (*SWbemSink, error) {
	var p *win32.IDispatch
	hr := win32.CoCreateInstance(&CLSID_SWbemSink, nil, 
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_ISWbemSink, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewSWbemSink(p, false, scoped), nil
}

func (this *SWbemSink) RegisterEventHandlers(handlers ISWbemSinkEventsHandlers) uint32 {
	var cpc *win32.IConnectionPointContainer
	hr := this.QueryInterface(&win32.IID_IConnectionPointContainer, unsafe.Pointer(&cpc))
	win32.ASSERT_SUCCEEDED(hr)

	var cp *win32.IConnectionPoint
	hr = cpc.FindConnectionPoint(&IID_ISWbemSinkEvents, &cp)
	win32.ASSERT_SUCCEEDED(hr)

	dispImpl := &ISWbemSinkEventsDispImpl{Handlers: handlers}
	disp := NewISWbemSinkEventsComObj(dispImpl, false)
	
	var cookie uint32
	hr = cp.Advise(disp.IUnknown(), &cookie)
	win32.ASSERT_SUCCEEDED(hr)

	disp.Release()
	cp.Release()
	cpc.Release()
	return cookie
}

func (this *SWbemSink) UnRegisterEventHandlers(cookie uint32) {
	var cpc *win32.IConnectionPointContainer
	hr := this.QueryInterface(&win32.IID_IConnectionPointContainer, unsafe.Pointer(&cpc))
	win32.ASSERT_SUCCEEDED(hr)

	var cp *win32.IConnectionPoint
	hr = cpc.FindConnectionPoint(&IID_ISWbemSinkEvents, &cp)
	win32.ASSERT_SUCCEEDED(hr)

	hr = cp.Unadvise(cookie)
	win32.ASSERT_SUCCEEDED(hr)

	cp.Release()
	cpc.Release()
}

