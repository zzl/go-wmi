package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// 76A6415A-CB41-11D1-8B02-00600806D9B6
var IID_ISWbemObject = syscall.GUID{0x76A6415A, 0xCB41, 0x11D1, 
	[8]byte{0x8B, 0x02, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemObject struct {
	ole.OleClient
}

func NewISWbemObject(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemObject {
	 if pDisp == nil {
		return nil;
	}
	p := &ISWbemObject{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemObjectFromVar(v ole.Variant) *ISWbemObject {
	return NewISWbemObject(v.IDispatch(), false, false)
}

func (this *ISWbemObject) IID() *syscall.GUID {
	return &IID_ISWbemObject
}

func (this *ISWbemObject) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

var ISWbemObject_Put__OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemObject) Put_(optArgs ...interface{}) *ISWbemObjectPath {
	optArgs = ole.ProcessOptArgs(ISWbemObject_Put__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000001, nil, optArgs...)
	return NewISWbemObjectPath(retVal.IDispatch(), false, true)
}

var ISWbemObject_PutAsync__OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemObject) PutAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemObject_PutAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000002, []interface{}{objWbemSink}, optArgs...)
	_= retVal
}

var ISWbemObject_Delete__OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemObject) Delete_(optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemObject_Delete__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000003, nil, optArgs...)
	_= retVal
}

var ISWbemObject_DeleteAsync__OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemObject) DeleteAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemObject_DeleteAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000004, []interface{}{objWbemSink}, optArgs...)
	_= retVal
}

var ISWbemObject_Instances__OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemObject) Instances_(optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemObject_Instances__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000005, nil, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemObject_InstancesAsync__OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemObject) InstancesAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemObject_InstancesAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000006, []interface{}{objWbemSink}, optArgs...)
	_= retVal
}

var ISWbemObject_Subclasses__OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemObject) Subclasses_(optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemObject_Subclasses__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000007, nil, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemObject_SubclassesAsync__OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemObject) SubclassesAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemObject_SubclassesAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000008, []interface{}{objWbemSink}, optArgs...)
	_= retVal
}

var ISWbemObject_Associators__OptArgs= []string{
	"strAssocClass", "strResultClass", "strResultRole", "strRole", 
	"bClassesOnly", "bSchemaOnly", "strRequiredAssocQualifier", "strRequiredQualifier", 
	"iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemObject) Associators_(optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemObject_Associators__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000009, nil, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemObject_AssociatorsAsync__OptArgs= []string{
	"strAssocClass", "strResultClass", "strResultRole", "strRole", 
	"bClassesOnly", "bSchemaOnly", "strRequiredAssocQualifier", "strRequiredQualifier", 
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemObject) AssociatorsAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemObject_AssociatorsAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000a, []interface{}{objWbemSink}, optArgs...)
	_= retVal
}

var ISWbemObject_References__OptArgs= []string{
	"strResultClass", "strRole", "bClassesOnly", "bSchemaOnly", 
	"strRequiredQualifier", "iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemObject) References_(optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemObject_References__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000b, nil, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemObject_ReferencesAsync__OptArgs= []string{
	"strResultClass", "strRole", "bClassesOnly", "bSchemaOnly", 
	"strRequiredQualifier", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemObject) ReferencesAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemObject_ReferencesAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000c, []interface{}{objWbemSink}, optArgs...)
	_= retVal
}

var ISWbemObject_ExecMethod__OptArgs= []string{
	"objWbemInParameters", "iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemObject) ExecMethod_(strMethodName string, optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemObject_ExecMethod__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000d, []interface{}{strMethodName}, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemObject_ExecMethodAsync__OptArgs= []string{
	"objWbemInParameters", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemObject) ExecMethodAsync_(objWbemSink *win32.IDispatch, strMethodName string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemObject_ExecMethodAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000e, []interface{}{objWbemSink, strMethodName}, optArgs...)
	_= retVal
}

func (this *ISWbemObject) Clone_() *ISWbemObject {
	retVal, _ := this.Call(0x0000000f, nil)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemObject_GetObjectText__OptArgs= []string{
	"iFlags", 
}

func (this *ISWbemObject) GetObjectText_(optArgs ...interface{}) string {
	optArgs = ole.ProcessOptArgs(ISWbemObject_GetObjectText__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000010, nil, optArgs...)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

var ISWbemObject_SpawnDerivedClass__OptArgs= []string{
	"iFlags", 
}

func (this *ISWbemObject) SpawnDerivedClass_(optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemObject_SpawnDerivedClass__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000011, nil, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemObject_SpawnInstance__OptArgs= []string{
	"iFlags", 
}

func (this *ISWbemObject) SpawnInstance_(optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemObject_SpawnInstance__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000012, nil, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemObject_CompareTo__OptArgs= []string{
	"iFlags", 
}

func (this *ISWbemObject) CompareTo_(objWbemObject *win32.IDispatch, optArgs ...interface{}) bool {
	optArgs = ole.ProcessOptArgs(ISWbemObject_CompareTo__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000013, []interface{}{objWbemObject}, optArgs...)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemObject) Qualifiers_() *ISWbemQualifierSet {
	retVal, _ := this.PropGet(0x00000014, nil)
	return NewISWbemQualifierSet(retVal.IDispatch(), false, true)
}

func (this *ISWbemObject) Properties_() *ISWbemPropertySet {
	retVal, _ := this.PropGet(0x00000015, nil)
	return NewISWbemPropertySet(retVal.IDispatch(), false, true)
}

func (this *ISWbemObject) Methods_() *ISWbemMethodSet {
	retVal, _ := this.PropGet(0x00000016, nil)
	return NewISWbemMethodSet(retVal.IDispatch(), false, true)
}

func (this *ISWbemObject) Derivation_() ole.Variant {
	retVal, _ := this.PropGet(0x00000017, nil)
	com.AddToScope(retVal)
	return *retVal
}

func (this *ISWbemObject) Path_() *ISWbemObjectPath {
	retVal, _ := this.PropGet(0x00000018, nil)
	return NewISWbemObjectPath(retVal.IDispatch(), false, true)
}

func (this *ISWbemObject) Security_() *ISWbemSecurity {
	retVal, _ := this.PropGet(0x00000019, nil)
	return NewISWbemSecurity(retVal.IDispatch(), false, true)
}

