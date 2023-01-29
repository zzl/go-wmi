package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

// 9B16ED16-D3DF-11D1-8B08-00600806D9B6
var IID_ISWbemQualifierSet = syscall.GUID{0x9B16ED16, 0xD3DF, 0x11D1,
	[8]byte{0x8B, 0x08, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemQualifierSet struct {
	ole.OleClient
}

func NewISWbemQualifierSet(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemQualifierSet {
	if pDisp == nil {
		return nil
	}
	p := &ISWbemQualifierSet{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemQualifierSetFromVar(v ole.Variant) *ISWbemQualifierSet {
	return NewISWbemQualifierSet(v.IDispatch(), false, false)
}

func (this *ISWbemQualifierSet) IID() *syscall.GUID {
	return &IID_ISWbemQualifierSet
}

func (this *ISWbemQualifierSet) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemQualifierSet) NewEnum_() *com.UnknownClass {
	retVal, _ := this.PropGet(-4, nil)
	return com.NewUnknownClass(retVal.PunkValVal(), true)
}

func (this *ISWbemQualifierSet) ForEach(action func(item *ISWbemQualifier) bool) {
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
		pItem := (*ISWbemQualifier)(v.ToPointer())
		ret := action(pItem)
		v.Clear()
		if !ret {
			break
		}
	}
}

var ISWbemQualifierSet_Item_OptArgs = []string{
	"iFlags",
}

func (this *ISWbemQualifierSet) Item(name string, optArgs ...interface{}) *ISWbemQualifier {
	optArgs = ole.ProcessOptArgs(ISWbemQualifierSet_Item_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000000, []interface{}{name}, optArgs...)
	return NewISWbemQualifier(retVal.IDispatch(), false, true)
}

func (this *ISWbemQualifierSet) Count() int32 {
	retVal, _ := this.PropGet(0x00000001, nil)
	return retVal.LValVal()
}

var ISWbemQualifierSet_Add_OptArgs = []string{
	"bPropagatesToSubclass", "bPropagatesToInstance", "bIsOverridable", "iFlags",
}

func (this *ISWbemQualifierSet) Add(strName string, varVal *ole.Variant, optArgs ...interface{}) *ISWbemQualifier {
	optArgs = ole.ProcessOptArgs(ISWbemQualifierSet_Add_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000002, []interface{}{strName, varVal}, optArgs...)
	return NewISWbemQualifier(retVal.IDispatch(), false, true)
}

var ISWbemQualifierSet_Remove_OptArgs = []string{
	"iFlags",
}

func (this *ISWbemQualifierSet) Remove(strName string, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemQualifierSet_Remove_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000003, []interface{}{strName}, optArgs...)
	_ = retVal
}

