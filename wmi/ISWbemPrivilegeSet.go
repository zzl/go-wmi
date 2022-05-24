package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

// 26EE67BF-5804-11D2-8B4A-00600806D9B6
var IID_ISWbemPrivilegeSet = syscall.GUID{0x26EE67BF, 0x5804, 0x11D2, 
	[8]byte{0x8B, 0x4A, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemPrivilegeSet struct {
	ole.OleClient
}

func NewISWbemPrivilegeSet(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemPrivilegeSet {
	 if pDisp == nil {
		return nil;
	}
	p := &ISWbemPrivilegeSet{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemPrivilegeSetFromVar(v ole.Variant) *ISWbemPrivilegeSet {
	return NewISWbemPrivilegeSet(v.IDispatch(), false, false)
}

func (this *ISWbemPrivilegeSet) IID() *syscall.GUID {
	return &IID_ISWbemPrivilegeSet
}

func (this *ISWbemPrivilegeSet) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemPrivilegeSet) NewEnum_() *com.UnknownClass {
	retVal, _ := this.PropGet(-4, nil)
	return com.NewUnknownClass(retVal.PunkValVal(), true)
}

func (this *ISWbemPrivilegeSet) ForEach(action func(item *ISWbemPrivilege) bool) {
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
		pItem := (*ISWbemPrivilege)(v.ToPointer())
		ret := action(pItem)
		v.Clear()
		if !ret {
			break
		}
	}
}

func (this *ISWbemPrivilegeSet) Item(iPrivilege int32) *ISWbemPrivilege {
	retVal, _ := this.Call(0x00000000, []interface{}{iPrivilege})
	return NewISWbemPrivilege(retVal.IDispatch(), false, true)
}

func (this *ISWbemPrivilegeSet) Count() int32 {
	retVal, _ := this.PropGet(0x00000001, nil)
	return retVal.LValVal()
}

var ISWbemPrivilegeSet_Add_OptArgs= []string{
	"bIsEnabled", 
}

func (this *ISWbemPrivilegeSet) Add(iPrivilege int32, optArgs ...interface{}) *ISWbemPrivilege {
	optArgs = ole.ProcessOptArgs(ISWbemPrivilegeSet_Add_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000002, []interface{}{iPrivilege}, optArgs...)
	return NewISWbemPrivilege(retVal.IDispatch(), false, true)
}

func (this *ISWbemPrivilegeSet) Remove(iPrivilege int32)  {
	retVal, _ := this.Call(0x00000003, []interface{}{iPrivilege})
	_= retVal
}

func (this *ISWbemPrivilegeSet) DeleteAll()  {
	retVal, _ := this.Call(0x00000004, nil)
	_= retVal
}

var ISWbemPrivilegeSet_AddAsString_OptArgs= []string{
	"bIsEnabled", 
}

func (this *ISWbemPrivilegeSet) AddAsString(strPrivilege string, optArgs ...interface{}) *ISWbemPrivilege {
	optArgs = ole.ProcessOptArgs(ISWbemPrivilegeSet_AddAsString_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000005, []interface{}{strPrivilege}, optArgs...)
	return NewISWbemPrivilege(retVal.IDispatch(), false, true)
}

