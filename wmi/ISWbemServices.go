package wmi

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// 76A6415C-CB41-11D1-8B02-00600806D9B6
var IID_ISWbemServices = syscall.GUID{0x76A6415C, 0xCB41, 0x11D1, 
	[8]byte{0x8B, 0x02, 0x00, 0x60, 0x08, 0x06, 0xD9, 0xB6}}

type ISWbemServices struct {
	ole.OleClient
}

func NewISWbemServices(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemServices {
	 if pDisp == nil {
		return nil;
	}
	p := &ISWbemServices{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemServicesFromVar(v ole.Variant) *ISWbemServices {
	return NewISWbemServices(v.IDispatch(), false, false)
}

func (this *ISWbemServices) IID() *syscall.GUID {
	return &IID_ISWbemServices
}

func (this *ISWbemServices) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

var ISWbemServices_Get_OptArgs= []string{
	"strObjectPath", "iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServices) Get(optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemServices_Get_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000001, nil, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemServices_GetAsync_OptArgs= []string{
	"strObjectPath", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServices) GetAsync(objWbemSink *win32.IDispatch, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServices_GetAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000002, []interface{}{objWbemSink}, optArgs...)
	_= retVal
}

var ISWbemServices_Delete_OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServices) Delete(strObjectPath string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServices_Delete_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000003, []interface{}{strObjectPath}, optArgs...)
	_= retVal
}

var ISWbemServices_DeleteAsync_OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServices) DeleteAsync(objWbemSink *win32.IDispatch, strObjectPath string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServices_DeleteAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000004, []interface{}{objWbemSink, strObjectPath}, optArgs...)
	_= retVal
}

var ISWbemServices_InstancesOf_OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServices) InstancesOf(strClass string, optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemServices_InstancesOf_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000005, []interface{}{strClass}, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemServices_InstancesOfAsync_OptArgs= []string{
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServices) InstancesOfAsync(objWbemSink *win32.IDispatch, strClass string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServices_InstancesOfAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000006, []interface{}{objWbemSink, strClass}, optArgs...)
	_= retVal
}

var ISWbemServices_SubclassesOf_OptArgs= []string{
	"strSuperclass", "iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServices) SubclassesOf(optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemServices_SubclassesOf_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000007, nil, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemServices_SubclassesOfAsync_OptArgs= []string{
	"strSuperclass", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServices) SubclassesOfAsync(objWbemSink *win32.IDispatch, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServices_SubclassesOfAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000008, []interface{}{objWbemSink}, optArgs...)
	_= retVal
}

var ISWbemServices_ExecQuery_OptArgs= []string{
	"strQueryLanguage", "iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServices) ExecQuery(strQuery string, optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemServices_ExecQuery_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000009, []interface{}{strQuery}, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemServices_ExecQueryAsync_OptArgs= []string{
	"strQueryLanguage", "lFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServices) ExecQueryAsync(objWbemSink *win32.IDispatch, strQuery string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServices_ExecQueryAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000a, []interface{}{objWbemSink, strQuery}, optArgs...)
	_= retVal
}

var ISWbemServices_AssociatorsOf_OptArgs= []string{
	"strAssocClass", "strResultClass", "strResultRole", "strRole", 
	"bClassesOnly", "bSchemaOnly", "strRequiredAssocQualifier", "strRequiredQualifier", 
	"iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServices) AssociatorsOf(strObjectPath string, optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemServices_AssociatorsOf_OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000b, []interface{}{strObjectPath}, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemServices_AssociatorsOfAsync_OptArgs= []string{
	"strAssocClass", "strResultClass", "strResultRole", "strRole", 
	"bClassesOnly", "bSchemaOnly", "strRequiredAssocQualifier", "strRequiredQualifier", 
	"iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServices) AssociatorsOfAsync(objWbemSink *win32.IDispatch, strObjectPath string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServices_AssociatorsOfAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000c, []interface{}{objWbemSink, strObjectPath}, optArgs...)
	_= retVal
}

var ISWbemServices_ReferencesTo_OptArgs= []string{
	"strResultClass", "strRole", "bClassesOnly", "bSchemaOnly", 
	"strRequiredQualifier", "iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServices) ReferencesTo(strObjectPath string, optArgs ...interface{}) *ISWbemObjectSet {
	optArgs = ole.ProcessOptArgs(ISWbemServices_ReferencesTo_OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000d, []interface{}{strObjectPath}, optArgs...)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemServices_ReferencesToAsync_OptArgs= []string{
	"strResultClass", "strRole", "bClassesOnly", "bSchemaOnly", 
	"strRequiredQualifier", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServices) ReferencesToAsync(objWbemSink *win32.IDispatch, strObjectPath string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServices_ReferencesToAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000e, []interface{}{objWbemSink, strObjectPath}, optArgs...)
	_= retVal
}

var ISWbemServices_ExecNotificationQuery_OptArgs= []string{
	"strQueryLanguage", "iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServices) ExecNotificationQuery(strQuery string, optArgs ...interface{}) *ISWbemEventSource {
	optArgs = ole.ProcessOptArgs(ISWbemServices_ExecNotificationQuery_OptArgs, optArgs)
	retVal, _ := this.Call(0x0000000f, []interface{}{strQuery}, optArgs...)
	return NewISWbemEventSource(retVal.IDispatch(), false, true)
}

var ISWbemServices_ExecNotificationQueryAsync_OptArgs= []string{
	"strQueryLanguage", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServices) ExecNotificationQueryAsync(objWbemSink *win32.IDispatch, strQuery string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServices_ExecNotificationQueryAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000010, []interface{}{objWbemSink, strQuery}, optArgs...)
	_= retVal
}

var ISWbemServices_ExecMethod_OptArgs= []string{
	"objWbemInParameters", "iFlags", "objWbemNamedValueSet", 
}

func (this *ISWbemServices) ExecMethod(strObjectPath string, strMethodName string, optArgs ...interface{}) *ISWbemObject {
	optArgs = ole.ProcessOptArgs(ISWbemServices_ExecMethod_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000011, []interface{}{strObjectPath, strMethodName}, optArgs...)
	return NewISWbemObject(retVal.IDispatch(), false, true)
}

var ISWbemServices_ExecMethodAsync_OptArgs= []string{
	"objWbemInParameters", "iFlags", "objWbemNamedValueSet", "objWbemAsyncContext", 
}

func (this *ISWbemServices) ExecMethodAsync(objWbemSink *win32.IDispatch, strObjectPath string, strMethodName string, optArgs ...interface{})  {
	optArgs = ole.ProcessOptArgs(ISWbemServices_ExecMethodAsync_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000012, []interface{}{objWbemSink, strObjectPath, strMethodName}, optArgs...)
	_= retVal
}

func (this *ISWbemServices) Security_() *ISWbemSecurity {
	retVal, _ := this.PropGet(0x00000013, nil)
	return NewISWbemSecurity(retVal.IDispatch(), false, true)
}

