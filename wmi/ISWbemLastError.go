package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// D962DB84-D4BB-11D1-8B09-00600806D9B6
var IID_ISWbemLastError = syscall.GUID{0xD962DB84, 0xD4BB, 0x11D1,
	[8]byte{0x8B, 0x09, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemLastError struct {
	ole.OleClient
}

func NewISWbemLastError(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemLastError {
	if pDisp == nil {
		return nil
	}
	p := &ISWbemLastError{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemLastErrorFromVar(v ole.Variant) *ISWbemLastError {
	return NewISWbemLastError(v.IDispatch(), false, false)
}

func (this *ISWbemLastError) IID() *syscall.GUID {
	return &IID_ISWbemLastError
}

func (this *ISWbemLastError) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

var ISWbemLastError_Put__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemLastError) Put_(optArgs ...interface{}) *ISWbemObjectPath {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_Put__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000001, nil, optArgs...)
	return NewISWbemObjectPath(retVal.IDispatch(), false, true)
}

var ISWbemLastError_PutAsync__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext",
}

func (this *ISWbemLastError) PutAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_PutAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000002, []interface{}{objWbemSink}, optArgs...)
	_ = retVal
}

var ISWbemLastError_Delete__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemLastError) Delete_(optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_Delete__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000003, nil, optArgs...)
	_ = retVal
}

var ISWbemLastError_DeleteAsync__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext",
}

func (this *ISWbemLastError) DeleteAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_DeleteAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000004, []interface{}{objWbemSink}, optArgs...)
	_ = retVal
}

var ISWbemLastError_Instances__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemLastError) Instances_(optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_Instances__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000005, nil, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemLastError_InstancesAsync__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext",
}

func (this *ISWbemLastError) InstancesAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_InstancesAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000006, []interface{}{objWbemSink}, optArgs...)
	_ = retVal
}

var ISWbemLastError_Subclasses__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemLastError) Subclasses_(optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_Subclasses__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000007, nil, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemLastError_SubclassesAsync__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext",
}

func (this *ISWbemLastError) SubclassesAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_SubclassesAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000008, []interface{}{objWbemSink}, optArgs...)
	_ = retVal
}

var ISWbemLastError_Associators__OptArgs = []string{
	"strAssocClass", "strResultClass", "strResultRole", "strRole",
	"bClassesOnly", "bSchemaOnly", "strRequiredAssocQualifier", "strRequiredQualifier",
	"iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemLastError) Associators_(optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_Associators__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000009, nil, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemLastError_AssociatorsAsync__OptArgs = []string{
	"strAssocClass", "strResultClass", "strResultRole", "strRole",
	"bClassesOnly", "bSchemaOnly", "strRequiredAssocQualifier", "strRequiredQualifier",
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext",
}

func (this *ISWbemLastError) AssociatorsAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_AssociatorsAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000a, []interface{}{objWbemSink}, optArgs...)
	_ = retVal
}

var ISWbemLastError_References__OptArgs = []string{
	"strResultClass", "strRole", "bClassesOnly", "bSchemaOnly",
	"strRequiredQualifier", "iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemLastError) References_(optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_References__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000b, nil, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemLastError_ReferencesAsync__OptArgs = []string{
	"strResultClass", "strRole", "bClassesOnly", "bSchemaOnly",
	"strRequiredQualifier", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext",
}

func (this *ISWbemLastError) ReferencesAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_ReferencesAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000c, []interface{}{objWbemSink}, optArgs...)
	_ = retVal
}

var ISWbemLastError_ExecMethod__OptArgs = []string{
	"objWbemInParameters", "iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemLastError) ExecMethod_(strMethodName string, optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_ExecMethod__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000d, []interface{}{strMethodName}, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemLastError_ExecMethodAsync__OptArgs = []string{
	"objWbemInParameters", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext",
}

func (this *ISWbemLastError) ExecMethodAsync_(objWbemSink *win32.IDispatch, strMethodName string, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_ExecMethodAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000e, []interface{}{objWbemSink, strMethodName}, optArgs...)
	_ = retVal
}

func (this *ISWbemLastError) Clone_() *ISWbemObject {
	retVal, _ := this.Call(0x0000000f, nil)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemLastError_GetObjectText__OptArgs = []string{
	"iFlags",
}

func (this *ISWbemLastError) GetObjectText_(optArgs ...interface{}) string {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_GetObjectText__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000010, nil, optArgs...)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

var ISWbemLastError_SpawnDerivedClass__OptArgs = []string{
	"iFlags",
}

func (this *ISWbemLastError) SpawnDerivedClass_(optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_SpawnDerivedClass__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000011, nil, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemLastError_SpawnInstance__OptArgs = []string{
	"iFlags",
}

func (this *ISWbemLastError) SpawnInstance_(optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_SpawnInstance__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000012, nil, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemLastError_CompareTo__OptArgs = []string{
	"iFlags",
}

func (this *ISWbemLastError) CompareTo_(objWbemObject *win32.IDispatch, optArgs ...interface{}) bool {
	optArgs = ole.ProcessOptArgs(ISWbemLastError_CompareTo__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000013, []interface{}{objWbemObject}, optArgs...)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemLastError) Qualifiers_() *ISWbemQualifierSet {
	retVal, _ := this.PropGet(0x00000014, nil)
	return NewISWbemQualifierSet(retVal.IDispatch(), false, true)
}

func (this *ISWbemLastError) Properties_() *ISWbemPropertySet {
	retVal, _ := this.PropGet(0x00000015, nil)
	return NewISWbemPropertySet(retVal.IDispatch(), false, true)
}

func (this *ISWbemLastError) Methods_() *ISWbemMethodSet {
	retVal, _ := this.PropGet(0x00000016, nil)
	return NewISWbemMethodSet(retVal.IDispatch(), false, true)
}

func (this *ISWbemLastError) Derivation_() ole.Variant {
	retVal, _ := this.PropGet(0x00000017, nil)
	com.AddToScope(retVal)
	return *retVal
}

func (this *ISWbemLastError) Path_() *ISWbemObjectPath {
	retVal, _ := this.PropGet(0x00000018, nil)
	return NewISWbemObjectPath(retVal.IDispatch(), false, true)
}

func (this *ISWbemLastError) Security_() *ISWbemSecurity {
	retVal, _ := this.PropGet(0x00000019, nil)
	return NewISWbemSecurity(retVal.IDispatch(), false, true)
}

