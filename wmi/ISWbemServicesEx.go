package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// D2F68443-85DC-427E-91D8-366554CC754C
var IID_ISWbemServicesEx = syscall.GUID{0xD2F68443, 0x85DC, 0x427E, 
	[8]byte{0x91, 0xD8, 0x36, 0x65, 0x54, 0xCC, 0x75, 0x4C}}

type ISWbemServicesEx struct {
	ole.OleClient
}

func NewISWbemServicesEx(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemServicesEx {
	 if pDisp == nil {
		return nil;
	}
	p := &ISWbemServicesEx{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemServicesExFromVar(v ole.Variant) *ISWbemServicesEx {
	return NewISWbemServicesEx(v.IDispatch(), false, false)
}

func (this *ISWbemServicesEx) IID() *syscall.GUID {
	return &IID_ISWbemServicesEx
}

func (this *ISWbemServicesEx) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

var ISWbemServicesEx_Get_OptArgs= []string{
	"strObjectPath", "iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServicesEx) Get(optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_Get_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000001, nil, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemServicesEx_GetAsync_OptArgs= []string{
	"strObjectPath", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServicesEx) GetAsync(objWbemSink *win32.IDispatch, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_GetAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000002, []interface{}{objWbemSink}, optArgs...)
	_= retVal
}

var ISWbemServicesEx_Delete_OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServicesEx) Delete(strObjectPath string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_Delete_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000003, []interface{}{strObjectPath}, optArgs...)
	_= retVal
}

var ISWbemServicesEx_DeleteAsync_OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServicesEx) DeleteAsync(objWbemSink *win32.IDispatch, strObjectPath string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_DeleteAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000004, []interface{}{objWbemSink, strObjectPath}, optArgs...)
	_= retVal
}

var ISWbemServicesEx_InstancesOf_OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServicesEx) InstancesOf(strClass string, optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_InstancesOf_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000005, []interface{}{strClass}, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemServicesEx_InstancesOfAsync_OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServicesEx) InstancesOfAsync(objWbemSink *win32.IDispatch, strClass string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_InstancesOfAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000006, []interface{}{objWbemSink, strClass}, optArgs...)
	_= retVal
}

var ISWbemServicesEx_SubclassesOf_OptArgs= []string{
	"strSuperclass", "iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServicesEx) SubclassesOf(optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_SubclassesOf_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000007, nil, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemServicesEx_SubclassesOfAsync_OptArgs= []string{
	"strSuperclass", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServicesEx) SubclassesOfAsync(objWbemSink *win32.IDispatch, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_SubclassesOfAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000008, []interface{}{objWbemSink}, optArgs...)
	_= retVal
}

var ISWbemServicesEx_ExecQuery_OptArgs= []string{
	"strQueryLanguage", "iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServicesEx) ExecQuery(strQuery string, optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_ExecQuery_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000009, []interface{}{strQuery}, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemServicesEx_ExecQueryAsync_OptArgs= []string{
	"strQueryLanguage", "lFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServicesEx) ExecQueryAsync(objWbemSink *win32.IDispatch, strQuery string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_ExecQueryAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000a, []interface{}{objWbemSink, strQuery}, optArgs...)
	_= retVal
}

var ISWbemServicesEx_AssociatorsOf_OptArgs= []string{
	"strAssocClass", "strResultClass", "strResultRole", "strRole", 
	"bClassesOnly", "bSchemaOnly", "strRequiredAssocQualifier", "strRequiredQualifier", 
	"iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServicesEx) AssociatorsOf(strObjectPath string, optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_AssociatorsOf_OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000b, []interface{}{strObjectPath}, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemServicesEx_AssociatorsOfAsync_OptArgs= []string{
	"strAssocClass", "strResultClass", "strResultRole", "strRole", 
	"bClassesOnly", "bSchemaOnly", "strRequiredAssocQualifier", "strRequiredQualifier", 
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServicesEx) AssociatorsOfAsync(objWbemSink *win32.IDispatch, strObjectPath string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_AssociatorsOfAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000c, []interface{}{objWbemSink, strObjectPath}, optArgs...)
	_= retVal
}

var ISWbemServicesEx_ReferencesTo_OptArgs= []string{
	"strResultClass", "strRole", "bClassesOnly", "bSchemaOnly", 
	"strRequiredQualifier", "iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServicesEx) ReferencesTo(strObjectPath string, optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_ReferencesTo_OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000d, []interface{}{strObjectPath}, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemServicesEx_ReferencesToAsync_OptArgs= []string{
	"strResultClass", "strRole", "bClassesOnly", "bSchemaOnly", 
	"strRequiredQualifier", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServicesEx) ReferencesToAsync(objWbemSink *win32.IDispatch, strObjectPath string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_ReferencesToAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000e, []interface{}{objWbemSink, strObjectPath}, optArgs...)
	_= retVal
}

var ISWbemServicesEx_ExecNotificationQuery_OptArgs= []string{
	"strQueryLanguage", "iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServicesEx) ExecNotificationQuery(strQuery string, optArgs ...interface{}) *ISWbemEventSource {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_ExecNotificationQuery_OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000f, []interface{}{strQuery}, optArgs...)
	return NewISWbemEventSource(retVal.IDispatch(), false, true)
}

var ISWbemServicesEx_ExecNotificationQueryAsync_OptArgs= []string{
	"strQueryLanguage", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServicesEx) ExecNotificationQueryAsync(objWbemSink *win32.IDispatch, strQuery string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_ExecNotificationQueryAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000010, []interface{}{objWbemSink, strQuery}, optArgs...)
	_= retVal
}

var ISWbemServicesEx_ExecMethod_OptArgs= []string{
	"objWbemInParameters", "iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServicesEx) ExecMethod(strObjectPath string, strMethodName string, optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_ExecMethod_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000011, []interface{}{strObjectPath, strMethodName}, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemServicesEx_ExecMethodAsync_OptArgs= []string{
	"objWbemInParameters", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServicesEx) ExecMethodAsync(objWbemSink *win32.IDispatch, strObjectPath string, strMethodName string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_ExecMethodAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000012, []interface{}{objWbemSink, strObjectPath, strMethodName}, optArgs...)
	_= retVal
}

func (this *ISWbemServicesEx) Security_() *ISWbemSecurity {
	retVal, _ := this.PropGet(0x00000013, nil)
	return NewISWbemSecurity(retVal.IDispatch(), false, true)
}

var ISWbemServicesEx_Put_OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServicesEx) Put(objWbemObject *ISWbemObjectEx, optArgs ...interface{}) *ISWbemObjectPath {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_Put_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000014, []interface{}{objWbemObject}, optArgs...)
	return NewISWbemObjectPath(retVal.IDispatch(), false, true)
}

var ISWbemServicesEx_PutAsync_OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServicesEx) PutAsync(objWbemSink *ISWbemSink, objWbemObject *ISWbemObjectEx, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServicesEx_PutAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000015, []interface{}{objWbemSink, objWbemObject}, optArgs...)
	_= retVal
}

