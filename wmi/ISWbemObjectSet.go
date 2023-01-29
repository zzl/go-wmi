package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

// 76A6415F-CB41-11D1-8B02-00600806D9B6
var IID_ISWbemObjectSet = syscall.GUID{0x76A6415F, 0xCB41, 0x11D1,
	[8]byte{0x8B, 0x02, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemObjectSet struct {
	ole.OleClient
}

func NewISWbemObjectSet(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemObjectSet {
	if pDisp == nil {
		return nil
	}
	p := &ISWbemObjectSet{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemObjectSetFromVar(v ole.Variant) *ISWbemObjectSet {
	return NewISWbemObjectSet(v.IDispatch(), false, false)
}

func (this *ISWbemObjectSet) IID() *syscall.GUID {
	return &IID_ISWbemObjectSet
}

func (this *ISWbemObjectSet) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemObjectSet) NewEnum_() *com.UnknownClass {
	retVal, _ := this.PropGet(-4, nil)
	return com.NewUnknownClass(retVal.PunkValVal(), true)
}

func (this *ISWbemObjectSet) ForEach(action func(item *ISWbemObject) bool) {
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
		pItem := (*ISWbemObject)(v.ToPointer())
		ret := action(pItem)
		v.Clear()
		if !ret {
			break
		}
	}
}

var ISWbemObjectSet_Item_OptArgs = []string{
	"iFlags",
}

func (this *ISWbemObjectSet) Item(strObjectPath string, optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemObjectSet_Item_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000000, []interface{}{strObjectPath}, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

func (this *ISWbemObjectSet) Count() int32 {
	retVal, _ := this.PropGet(0x00000001, nil)
	return retVal.LValVal()
}

func (this *ISWbemObjectSet) Security_() *ISWbemSecurity {
	retVal, _ := this.PropGet(0x00000004, nil)
	return NewISWbemSecurity(retVal.IDispatch(), false, true)
}

func (this *ISWbemObjectSet) ItemIndex(lIndex int32) *ISWbemObject {
	retVal, _ := this.Call(0x00000005, []interface{}{lIndex})
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

