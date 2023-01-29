package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

// CF2376EA-CE8C-11D1-8B05-00600806D9B6
var IID_ISWbemNamedValueSet = syscall.GUID{0xCF2376EA, 0xCE8C, 0x11D1,
	[8]byte{0x8B, 0x05, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemNamedValueSet struct {
	ole.OleClient
}

func NewISWbemNamedValueSet(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemNamedValueSet {
	if pDisp == nil {
		return nil
	}
	p := &ISWbemNamedValueSet{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemNamedValueSetFromVar(v ole.Variant) *ISWbemNamedValueSet {
	return NewISWbemNamedValueSet(v.IDispatch(), false, false)
}

func (this *ISWbemNamedValueSet) IID() *syscall.GUID {
	return &IID_ISWbemNamedValueSet
}

func (this *ISWbemNamedValueSet) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemNamedValueSet) NewEnum_() *com.UnknownClass {
	retVal, _ := this.PropGet(-4, nil)
	return com.NewUnknownClass(retVal.PunkValVal(), true)
}

func (this *ISWbemNamedValueSet) ForEach(action func(item *ISWbemNamedValue) bool) {
	pEnum := this.NewEnum_()
	var pEnumVar *win32.IEnumVARIANT
	pEnum.QueryInterface(&win32.IID_IEnumVARIANT, unsafe.Pointer(&pEnumVar))
	defer pEnumVar.Release()
	for {
		var c uint32
		var v ole.Variant
		pEnumVar.Next(1, (*win32.VARIANT)(&v), &c)
		if c == 0 {
			break
		}
		pItem := (*ISWbemNamedValue)(v.ToPointer())
		ret := action(pItem)
		v.Clear()
		if !ret {
			break
		}
	}
}

var ISWbemNamedValueSet_Item_OptArgs = []string{
	"iFlags",
}

func (this *ISWbemNamedValueSet) Item(strName string, optArgs ...interface{}) *ISWbemNamedValue {
	optArgs = ole.ProcessOptArgs(ISWbemNamedValueSet_Item_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000000, []interface{}{strName}, optArgs...)
	return NewISWbemNamedValue(retVal.IDispatch(), false, true)
}

func (this *ISWbemNamedValueSet) Count() int32 {
	retVal, _ := this.PropGet(0x00000001, nil)
	return retVal.LValVal()
}

var ISWbemNamedValueSet_Add_OptArgs = []string{
	"iFlags",
}

func (this *ISWbemNamedValueSet) Add(strName string, varValue *ole.Variant, optArgs ...interface{}) *ISWbemNamedValue {
	optArgs = ole.ProcessOptArgs(ISWbemNamedValueSet_Add_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000002, []interface{}{strName, varValue}, optArgs...)
	return NewISWbemNamedValue(retVal.IDispatch(), false, true)
}

var ISWbemNamedValueSet_Remove_OptArgs = []string{
	"iFlags",
}

func (this *ISWbemNamedValueSet) Remove(strName string, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemNamedValueSet_Remove_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000003, []interface{}{strName}, optArgs...)
	_ = retVal
}

func (this *ISWbemNamedValueSet) Clone() *ISWbemNamedValueSet {
	retVal, _ := this.Call(0x00000004, nil)
	return NewISWbemNamedValueSet(retVal.IDispatch(), false, true)
}

func (this *ISWbemNamedValueSet) DeleteAll() {
	retVal, _ := this.Call(0x00000005, nil)
	_ = retVal
}

