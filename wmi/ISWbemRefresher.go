package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

// 14D8250E-D9C2-11D3-B38F-00105A1F473A
var IID_ISWbemRefresher = syscall.GUID{0x14D8250E, 0xD9C2, 0x11D3, 
	[8]byte{0xB3, 0x8F, 0x00, 0x10, 0x5A, 0x1F, 0x47, 0x3A}}

type ISWbemRefresher struct {
	ole.OleClient
}

func NewISWbemRefresher(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemRefresher {
	 if pDisp == nil {
		return nil;
	}
	p := &ISWbemRefresher{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemRefresherFromVar(v ole.Variant) *ISWbemRefresher {
	return NewISWbemRefresher(v.IDispatch(), false, false)
}

func (this *ISWbemRefresher) IID() *syscall.GUID {
	return &IID_ISWbemRefresher
}

func (this *ISWbemRefresher) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemRefresher) NewEnum_() *com.UnknownClass {
	retVal, _ := this.PropGet(-4, nil)
	return com.NewUnknownClass(retVal.PunkValVal(), true)
}

func (this *ISWbemRefresher) ForEach(action func(item *ISWbemRefreshableItem) bool) {
	pEnum := this.NewEnum_()
	var pEnumVar *win32.IEnumVARIANT
	pEnum.QueryInterface(&win32.IID_IEnumVARIANT, unsafe.Pointer(&pEnumVar))
	defer pEnumVar.Release();
	for {
		var c uint32
		var v ole.Variant
		pEnumVar.Next(1, (*win32.VARIANT)(&v), &c)
		if c == 0 {
			break
		}
		pItem := (*ISWbemRefreshableItem)(v.ToPointer())
		ret := action(pItem)
		v.Clear()
		if !ret {
			break
		}
	}
}

func (this *ISWbemRefresher) Item(iIndex int32) *ISWbemRefreshableItem {
	retVal, _ := this.Call(0x00000000, []interface{}{iIndex})
	return NewISWbemRefreshableItem(retVal.IDispatch(), false, true)
}

func (this *ISWbemRefresher) Count() int32 {
	retVal, _ := this.PropGet(0x00000001, nil)
	return retVal.LValVal()
}

var ISWbemRefresher_Add_OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemRefresher) Add(objWbemServices *ISWbemServicesEx, bsInstancePath string, optArgs ...interface{}) *ISWbemRefreshableItem {
	optArgs = ole.ProcessOptArgs(ISWbemRefresher_Add_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000002, []interface{}{objWbemServices, bsInstancePath}, optArgs...)
	return NewISWbemRefreshableItem(retVal.IDispatch(), false, true)
}

var ISWbemRefresher_AddEnum_OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemRefresher) AddEnum(objWbemServices *ISWbemServicesEx, bsClassName string, optArgs ...interface{}) *ISWbemRefreshableItem {
	optArgs = ole.ProcessOptArgs(ISWbemRefresher_AddEnum_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000003, []interface{}{objWbemServices, bsClassName}, optArgs...)
	return NewISWbemRefreshableItem(retVal.IDispatch(), false, true)
}

var ISWbemRefresher_Remove_OptArgs= []string{
	"iFlags", 
}

func (this *ISWbemRefresher) Remove(iIndex int32, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemRefresher_Remove_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000004, []interface{}{iIndex}, optArgs...)
	_= retVal
}

var ISWbemRefresher_Refresh_OptArgs= []string{
	"iFlags", 
}

func (this *ISWbemRefresher) Refresh(optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemRefresher_Refresh_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000005, nil, optArgs...)
	_= retVal
}

func (this *ISWbemRefresher) AutoReconnect() bool {
	retVal, _ := this.PropGet(0x00000006, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemRefresher) SetAutoReconnect(rhs bool)  {
	_ = this.PropPut(0x00000006, []interface{}{rhs})
}

func (this *ISWbemRefresher) DeleteAll()  {
	retVal, _ := this.Call(0x00000007, nil)
	_= retVal
}

