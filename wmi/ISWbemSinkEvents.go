package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

// 75718CA0-F029-11D1-A1AC-00C04FB6C223
var IID_ISWbemSinkEvents = syscall.GUID{0x75718CA0, 0xF029, 0x11D1,
	[8]byte{0xA1, 0xAC, 0x00, 0xC0, 0x4F, 0xB6, 0xC2, 0x23}}

type ISWbemSinkEventsDispInterface interface {
	OnObjectReady(objWbemObject *ISWbemObject, objWbemAsyncContext *ISWbemNamedValueSet)
	OnCompleted(iHResult int32, objWbemErrorObject *ISWbemObject, objWbemAsyncContext *ISWbemNamedValueSet)
	OnProgress(iUpperBound int32, iCurrent int32, strMessage string, objWbemAsyncContext *ISWbemNamedValueSet)
	OnObjectPut(objWbemObjectPath *ISWbemObjectPath, objWbemAsyncContext *ISWbemNamedValueSet)
}

type ISWbemSinkEventsHandlers struct {
	OnObjectReady func(objWbemObject *ISWbemObject, objWbemAsyncContext *ISWbemNamedValueSet)
	OnCompleted   func(iHResult int32, objWbemErrorObject *ISWbemObject, objWbemAsyncContext *ISWbemNamedValueSet)
	OnProgress    func(iUpperBound int32, iCurrent int32, strMessage string, objWbemAsyncContext *ISWbemNamedValueSet)
	OnObjectPut   func(objWbemObjectPath *ISWbemObjectPath, objWbemAsyncContext *ISWbemNamedValueSet)
}

type ISWbemSinkEventsDispImpl struct {
	Handlers ISWbemSinkEventsHandlers
}

func (this *ISWbemSinkEventsDispImpl) OnObjectReady(objWbemObject *ISWbemObject, objWbemAsyncContext *ISWbemNamedValueSet) {
	if this.Handlers.OnObjectReady != nil {
		this.Handlers.OnObjectReady(objWbemObject, objWbemAsyncContext)
	}
}

func (this *ISWbemSinkEventsDispImpl) OnCompleted(iHResult int32, objWbemErrorObject *ISWbemObject, objWbemAsyncContext *ISWbemNamedValueSet) {
	if this.Handlers.OnCompleted != nil {
		this.Handlers.OnCompleted(iHResult, objWbemErrorObject, objWbemAsyncContext)
	}
}

func (this *ISWbemSinkEventsDispImpl) OnProgress(iUpperBound int32, iCurrent int32, strMessage string, objWbemAsyncContext *ISWbemNamedValueSet) {
	if this.Handlers.OnProgress != nil {
		this.Handlers.OnProgress(iUpperBound, iCurrent, strMessage, objWbemAsyncContext)
	}
}

func (this *ISWbemSinkEventsDispImpl) OnObjectPut(objWbemObjectPath *ISWbemObjectPath, objWbemAsyncContext *ISWbemNamedValueSet) {
	if this.Handlers.OnObjectPut != nil {
		this.Handlers.OnObjectPut(objWbemObjectPath, objWbemAsyncContext)
	}
}

type ISWbemSinkEventsImpl struct {
	ole.IDispatchImpl
	DispImpl ISWbemSinkEventsDispInterface
}

func (this *ISWbemSinkEventsImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ISWbemSinkEvents {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IDispatchImpl.QueryInterface(riid, ppvObject)
}

func (this *ISWbemSinkEventsImpl) Invoke(dispIdMember int32, riid *syscall.GUID, lcid uint32,
	wFlags uint16, pDispParams *win32.DISPPARAMS, pVarResult *win32.VARIANT,
	pExcepInfo *win32.EXCEPINFO, puArgErr *uint32) win32.HRESULT {
	var unwrapActions ole.Actions
	defer unwrapActions.Execute()
	switch dispIdMember {
	case 1:
		vArgs, _ := ole.ProcessInvokeArgs(pDispParams, 2)
		p1 := (*ISWbemObject)(vArgs[0].ToPointer())
		p2 := (*ISWbemNamedValueSet)(vArgs[1].ToPointer())
		this.DispImpl.OnObjectReady(p1, p2)
		return win32.S_OK
	case 2:
		vArgs, _ := ole.ProcessInvokeArgs(pDispParams, 3)
		p1, _ := vArgs[0].ToInt32()
		p2 := (*ISWbemObject)(vArgs[1].ToPointer())
		p3 := (*ISWbemNamedValueSet)(vArgs[2].ToPointer())
		this.DispImpl.OnCompleted(p1, p2, p3)
		return win32.S_OK
	case 3:
		vArgs, _ := ole.ProcessInvokeArgs(pDispParams, 4)
		p1, _ := vArgs[0].ToInt32()
		p2, _ := vArgs[1].ToInt32()
		p3, _ := vArgs[2].ToString()
		p4 := (*ISWbemNamedValueSet)(vArgs[3].ToPointer())
		this.DispImpl.OnProgress(p1, p2, p3, p4)
		return win32.S_OK
	case 4:
		vArgs, _ := ole.ProcessInvokeArgs(pDispParams, 2)
		p1 := (*ISWbemObjectPath)(vArgs[0].ToPointer())
		p2 := (*ISWbemNamedValueSet)(vArgs[1].ToPointer())
		this.DispImpl.OnObjectPut(p1, p2)
		return win32.S_OK
	}
	return win32.E_NOTIMPL
}

type ISWbemSinkEventsComObj struct {
	ole.IDispatchComObj
}

func NewISWbemSinkEventsComObj(dispImpl ISWbemSinkEventsDispInterface, scoped bool) *ISWbemSinkEventsComObj {
	comObj := com.NewComObj[ISWbemSinkEventsComObj](
		&ISWbemSinkEventsImpl{DispImpl: dispImpl})
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

