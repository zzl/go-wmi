package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

// C93BA292-D955-11D1-8B09-00600806D9B6
var IID_ISWbemMethodSet = syscall.GUID{0xC93BA292, 0xD955, 0x11D1,
	[8]byte{0x8B, 0x09, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemMethodSet struct {
	ole.OleClient
}

func NewISWbemMethodSet(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemMethodSet {
	if pDisp == nil {
		return nil
	}
	p := &ISWbemMethodSet{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemMethodSetFromVar(v ole.Variant) *ISWbemMethodSet {
	return NewISWbemMethodSet(v.IDispatch(), false, false)
}

func (this *ISWbemMethodSet) IID() *syscall.GUID {
	return &IID_ISWbemMethodSet
}

func (this *ISWbemMethodSet) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemMethodSet) NewEnum_() *com.UnknownClass {
	retVal, _ := this.PropGet(-4, nil)
	return com.NewUnknownClass(retVal.PunkValVal(), true)
}

func (this *ISWbemMethodSet) ForEach(action func(item *ISWbemMethod) bool) {
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
		pItem := (*ISWbemMethod)(v.ToPointer())
		ret := action(pItem)
		v.Clear()
		if !ret {
			break
		}
	}
}

var ISWbemMethodSet_Item_OptArgs = []string{
	"iFlags",
}

func (this *ISWbemMethodSet) Item(strName string, optArgs ...interface{}) *ISWbemMethod {
	optArgs = ole.ProcessOptArgs(ISWbemMethodSet_Item_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000000, []interface{}{strName}, optArgs...)
	return NewISWbemMethod(retVal.IDispatch(), false, true)
}

func (this *ISWbemMethodSet) Count() int32 {
	retVal, _ := this.PropGet(0x00000001, nil)
	return retVal.LValVal()
}

