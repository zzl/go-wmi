package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// 269AD56A-8A67-4129-BC8C-0506DCFE9880
var IID_ISWbemObjectEx = syscall.GUID{0x269AD56A, 0x8A67, 0x4129,
	[8]byte{0xBC, 0x8C, 0x05, 0x06, 0xDC, 0xFE, 0x98, 0x80}}

type ISWbemObjectEx struct {
	ole.OleClient
}

func NewISWbemObjectEx(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemObjectEx {
	if pDisp == nil {
		return nil
	}
	p := &ISWbemObjectEx{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemObjectExFromVar(v ole.Variant) *ISWbemObjectEx {
	return NewISWbemObjectEx(v.IDispatch(), false, false)
}

func (this *ISWbemObjectEx) IID() *syscall.GUID {
	return &IID_ISWbemObjectEx
}

func (this *ISWbemObjectEx) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

var ISWbemObjectEx_Put__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemObjectEx) Put_(optArgs ...interface{}) *ISWbemObjectPath {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_Put__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000001, nil, optArgs...)
	return NewISWbemObjectPath(retVal.IDispatch(), false, true)
}

var ISWbemObjectEx_PutAsync__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext",
}

func (this *ISWbemObjectEx) PutAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_PutAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000002, []interface{}{objWbemSink}, optArgs...)
	_ = retVal
}

var ISWbemObjectEx_Delete__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemObjectEx) Delete_(optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_Delete__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000003, nil, optArgs...)
	_ = retVal
}

var ISWbemObjectEx_DeleteAsync__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext",
}

func (this *ISWbemObjectEx) DeleteAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_DeleteAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000004, []interface{}{objWbemSink}, optArgs...)
	_ = retVal
}

var ISWbemObjectEx_Instances__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemObjectEx) Instances_(optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_Instances__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000005, nil, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemObjectEx_InstancesAsync__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext",
}

func (this *ISWbemObjectEx) InstancesAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_InstancesAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000006, []interface{}{objWbemSink}, optArgs...)
	_ = retVal
}

var ISWbemObjectEx_Subclasses__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemObjectEx) Subclasses_(optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_Subclasses__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000007, nil, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemObjectEx_SubclassesAsync__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext",
}

func (this *ISWbemObjectEx) SubclassesAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_SubclassesAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000008, []interface{}{objWbemSink}, optArgs...)
	_ = retVal
}

var ISWbemObjectEx_Associators__OptArgs = []string{
	"strAssocClass", "strResultClass", "strResultRole", "strRole",
	"bClassesOnly", "bSchemaOnly", "strRequiredAssocQualifier", "strRequiredQualifier",
	"iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemObjectEx) Associators_(optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_Associators__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000009, nil, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemObjectEx_AssociatorsAsync__OptArgs = []string{
	"strAssocClass", "strResultClass", "strResultRole", "strRole",
	"bClassesOnly", "bSchemaOnly", "strRequiredAssocQualifier", "strRequiredQualifier",
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext",
}

func (this *ISWbemObjectEx) AssociatorsAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_AssociatorsAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000a, []interface{}{objWbemSink}, optArgs...)
	_ = retVal
}

var ISWbemObjectEx_References__OptArgs = []string{
	"strResultClass", "strRole", "bClassesOnly", "bSchemaOnly",
	"strRequiredQualifier", "iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemObjectEx) References_(optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_References__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000b, nil, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemObjectEx_ReferencesAsync__OptArgs = []string{
	"strResultClass", "strRole", "bClassesOnly", "bSchemaOnly",
	"strRequiredQualifier", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext",
}

func (this *ISWbemObjectEx) ReferencesAsync_(objWbemSink *win32.IDispatch, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_ReferencesAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000c, []interface{}{objWbemSink}, optArgs...)
	_ = retVal
}

var ISWbemObjectEx_ExecMethod__OptArgs = []string{
	"objWbemInParameters", "iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemObjectEx) ExecMethod_(strMethodName string, optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_ExecMethod__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000d, []interface{}{strMethodName}, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemObjectEx_ExecMethodAsync__OptArgs = []string{
	"objWbemInParameters", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext",
}

func (this *ISWbemObjectEx) ExecMethodAsync_(objWbemSink *win32.IDispatch, strMethodName string, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_ExecMethodAsync__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000e, []interface{}{objWbemSink, strMethodName}, optArgs...)
	_ = retVal
}

func (this *ISWbemObjectEx) Clone_() *ISWbemObject {
	retVal, _ := this.Call(0x0000000f, nil)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemObjectEx_GetObjectText__OptArgs = []string{
	"iFlags",
}

func (this *ISWbemObjectEx) GetObjectText_(optArgs ...interface{}) string {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_GetObjectText__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000010, nil, optArgs...)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

var ISWbemObjectEx_SpawnDerivedClass__OptArgs = []string{
	"iFlags",
}

func (this *ISWbemObjectEx) SpawnDerivedClass_(optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_SpawnDerivedClass__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000011, nil, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemObjectEx_SpawnInstance__OptArgs = []string{
	"iFlags",
}

func (this *ISWbemObjectEx) SpawnInstance_(optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_SpawnInstance__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000012, nil, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemObjectEx_CompareTo__OptArgs = []string{
	"iFlags",
}

func (this *ISWbemObjectEx) CompareTo_(objWbemObject *win32.IDispatch, optArgs ...interface{}) bool {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_CompareTo__OptArgs, optArgs)
	retVal, _ := this.Call(0x00000013, []interface{}{objWbemObject}, optArgs...)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemObjectEx) Qualifiers_() *ISWbemQualifierSet {
	retVal, _ := this.PropGet(0x00000014, nil)
	return NewISWbemQualifierSet(retVal.IDispatch(), false, true)
}

func (this *ISWbemObjectEx) Properties_() *ISWbemPropertySet {
	retVal, _ := this.PropGet(0x00000015, nil)
	return NewISWbemPropertySet(retVal.IDispatch(), false, true)
}

func (this *ISWbemObjectEx) Methods_() *ISWbemMethodSet {
	retVal, _ := this.PropGet(0x00000016, nil)
	return NewISWbemMethodSet(retVal.IDispatch(), false, true)
}

func (this *ISWbemObjectEx) Derivation_() ole.Variant {
	retVal, _ := this.PropGet(0x00000017, nil)
	com.AddToScope(retVal)
	return *retVal
}

func (this *ISWbemObjectEx) Path_() *ISWbemObjectPath {
	retVal, _ := this.PropGet(0x00000018, nil)
	return NewISWbemObjectPath(retVal.IDispatch(), false, true)
}

func (this *ISWbemObjectEx) Security_() *ISWbemSecurity {
	retVal, _ := this.PropGet(0x00000019, nil)
	return NewISWbemSecurity(retVal.IDispatch(), false, true)
}

var ISWbemObjectEx_Refresh__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemObjectEx) Refresh_(optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_Refresh__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000001a, nil, optArgs...)
	_ = retVal
}

func (this *ISWbemObjectEx) SystemProperties_() *ISWbemPropertySet {
	retVal, _ := this.PropGet(0x0000001b, nil)
	return NewISWbemPropertySet(retVal.IDispatch(), false, true)
}

var ISWbemObjectEx_GetText__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemObjectEx) GetText_(iObjectTextFormat int32, optArgs ...interface{}) string {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_GetText__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000001c, []interface{}{iObjectTextFormat}, optArgs...)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

var ISWbemObjectEx_SetFromText__OptArgs = []string{
	"iFlags", "objWbemNamedValueSet",
}

func (this *ISWbemObjectEx) SetFromText_(bsText string, iObjectTextFormat int32, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemObjectEx_SetFromText__OptArgs, optArgs)
	retVal, _ := this.Call(0x0000001d, []interface{}{bsText, iObjectTextFormat}, optArgs...)
	_ = retVal
}

