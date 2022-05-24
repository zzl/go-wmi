package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

// DEA0A7B2-D4BA-11D1-8B09-00600806D9B6
var IID_ISWbemPropertySet = syscall.GUID{0xDEA0A7B2, 0xD4BA, 0x11D1, 
	[8]byte{0x8B, 0x09, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemPropertySet struct {
	ole.OleClient
}

func NewISWbemPropertySet(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemPropertySet {
	 if pDisp == nil {
		return nil;
	}
	p := &ISWbemPropertySet{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemPropertySetFromVar(v ole.Variant) *ISWbemPropertySet {
	return NewISWbemPropertySet(v.IDispatch(), false, false)
}

func (this *ISWbemPropertySet) IID() *syscall.GUID {
	return &IID_ISWbemPropertySet
}

func (this *ISWbemPropertySet) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemPropertySet) NewEnum_() *com.UnknownClass {
	retVal, _ := this.PropGet(-4, nil)
	return com.NewUnknownClass(retVal.PunkValVal(), true)
}

func (this *ISWbemPropertySet) ForEach(action func(item *ISWbemProperty) bool) {
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
		pItem := (*ISWbemProperty)(v.ToPointer())
		ret := action(pItem)
		v.Clear()
		if !ret {
			break
		}
	}
}

var ISWbemPropertySet_Item_OptArgs= []string{
	"iFlags", 
}

func (this *ISWbemPropertySet) Item(strName string, optArgs ...interface{}) *ISWbemProperty {
	optArgs = ole.ProcessOptArgs(ISWbemPropertySet_Item_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000000, []interface{}{strName}, optArgs...)
	return NewISWbemProperty(retVal.IDispatch(), false, true)
}

func (this *ISWbemPropertySet) Count() int32 {
	retVal, _ := this.PropGet(0x00000001, nil)
	return retVal.LValVal()
}

var ISWbemPropertySet_Add_OptArgs= []string{
	"bIsArray", "iFlags", 
}

func (this *ISWbemPropertySet) Add(strName string, iCimType int32, optArgs ...interface{}) *ISWbemProperty {
	optArgs = ole.ProcessOptArgs(ISWbemPropertySet_Add_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000002, []interface{}{strName, iCimType}, optArgs...)
	return NewISWbemProperty(retVal.IDispatch(), false, true)
}

var ISWbemPropertySet_Remove_OptArgs= []string{
	"iFlags", 
}

func (this *ISWbemPropertySet) Remove(strName string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemPropertySet_Remove_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000003, []interface{}{strName}, optArgs...)
	_= retVal
}

