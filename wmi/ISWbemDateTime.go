package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"time"
)

// 5E97458A-CF77-11D3-B38F-00105A1F473A
var IID_ISWbemDateTime = syscall.GUID{0x5E97458A, 0xCF77, 0x11D3,
	[8]byte{0xB3, 0x8F, 0x00, 0x10, 0x5A, 0x1F, 0x47, 0x3A}}

type ISWbemDateTime struct {
	ole.OleClient
}

func NewISWbemDateTime(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemDateTime {
	if pDisp == nil {
		return nil
	}
	p := &ISWbemDateTime{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemDateTimeFromVar(v ole.Variant) *ISWbemDateTime {
	return NewISWbemDateTime(v.IDispatch(), false, false)
}

func (this *ISWbemDateTime) IID() *syscall.GUID {
	return &IID_ISWbemDateTime
}

func (this *ISWbemDateTime) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemDateTime) Value() string {
	retVal, _ := this.PropGet(0x00000000, nil)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

func (this *ISWbemDateTime) SetValue(rhs string) {
	_ = this.PropPut(0x00000000, []interface{}{rhs})
}

func (this *ISWbemDateTime) Year() int32 {
	retVal, _ := this.PropGet(0x00000001, nil)
	return retVal.LValVal()
}

func (this *ISWbemDateTime) SetYear(rhs int32) {
	_ = this.PropPut(0x00000001, []interface{}{rhs})
}

func (this *ISWbemDateTime) YearSpecified() bool {
	retVal, _ := this.PropGet(0x00000002, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemDateTime) SetYearSpecified(rhs bool) {
	_ = this.PropPut(0x00000002, []interface{}{rhs})
}

func (this *ISWbemDateTime) Month() int32 {
	retVal, _ := this.PropGet(0x00000003, nil)
	return retVal.LValVal()
}

func (this *ISWbemDateTime) SetMonth(rhs int32) {
	_ = this.PropPut(0x00000003, []interface{}{rhs})
}

func (this *ISWbemDateTime) MonthSpecified() bool {
	retVal, _ := this.PropGet(0x00000004, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemDateTime) SetMonthSpecified(rhs bool) {
	_ = this.PropPut(0x00000004, []interface{}{rhs})
}

func (this *ISWbemDateTime) Day() int32 {
	retVal, _ := this.PropGet(0x00000005, nil)
	return retVal.LValVal()
}

func (this *ISWbemDateTime) SetDay(rhs int32) {
	_ = this.PropPut(0x00000005, []interface{}{rhs})
}

func (this *ISWbemDateTime) DaySpecified() bool {
	retVal, _ := this.PropGet(0x00000006, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemDateTime) SetDaySpecified(rhs bool) {
	_ = this.PropPut(0x00000006, []interface{}{rhs})
}

func (this *ISWbemDateTime) Hours() int32 {
	retVal, _ := this.PropGet(0x00000007, nil)
	return retVal.LValVal()
}

func (this *ISWbemDateTime) SetHours(rhs int32) {
	_ = this.PropPut(0x00000007, []interface{}{rhs})
}

func (this *ISWbemDateTime) HoursSpecified() bool {
	retVal, _ := this.PropGet(0x00000008, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemDateTime) SetHoursSpecified(rhs bool) {
	_ = this.PropPut(0x00000008, []interface{}{rhs})
}

func (this *ISWbemDateTime) Minutes() int32 {
	retVal, _ := this.PropGet(0x00000009, nil)
	return retVal.LValVal()
}

func (this *ISWbemDateTime) SetMinutes(rhs int32) {
	_ = this.PropPut(0x00000009, []interface{}{rhs})
}

func (this *ISWbemDateTime) MinutesSpecified() bool {
	retVal, _ := this.PropGet(0x0000000a, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemDateTime) SetMinutesSpecified(rhs bool) {
	_ = this.PropPut(0x0000000a, []interface{}{rhs})
}

func (this *ISWbemDateTime) Seconds() int32 {
	retVal, _ := this.PropGet(0x0000000b, nil)
	return retVal.LValVal()
}

func (this *ISWbemDateTime) SetSeconds(rhs int32) {
	_ = this.PropPut(0x0000000b, []interface{}{rhs})
}

func (this *ISWbemDateTime) SecondsSpecified() bool {
	retVal, _ := this.PropGet(0x0000000c, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemDateTime) SetSecondsSpecified(rhs bool) {
	_ = this.PropPut(0x0000000c, []interface{}{rhs})
}

func (this *ISWbemDateTime) Microseconds() int32 {
	retVal, _ := this.PropGet(0x0000000d, nil)
	return retVal.LValVal()
}

func (this *ISWbemDateTime) SetMicroseconds(rhs int32) {
	_ = this.PropPut(0x0000000d, []interface{}{rhs})
}

func (this *ISWbemDateTime) MicrosecondsSpecified() bool {
	retVal, _ := this.PropGet(0x0000000e, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemDateTime) SetMicrosecondsSpecified(rhs bool) {
	_ = this.PropPut(0x0000000e, []interface{}{rhs})
}

func (this *ISWbemDateTime) UTC() int32 {
	retVal, _ := this.PropGet(0x0000000f, nil)
	return retVal.LValVal()
}

func (this *ISWbemDateTime) SetUTC(rhs int32) {
	_ = this.PropPut(0x0000000f, []interface{}{rhs})
}

func (this *ISWbemDateTime) UTCSpecified() bool {
	retVal, _ := this.PropGet(0x00000010, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemDateTime) SetUTCSpecified(rhs bool) {
	_ = this.PropPut(0x00000010, []interface{}{rhs})
}

func (this *ISWbemDateTime) IsInterval() bool {
	retVal, _ := this.PropGet(0x00000011, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemDateTime) SetIsInterval(rhs bool) {
	_ = this.PropPut(0x00000011, []interface{}{rhs})
}

var ISWbemDateTime_GetVarDate_OptArgs = []string{
	"bIsLocal",
}

func (this *ISWbemDateTime) GetVarDate(optArgs ...interface{}) time.Time {
	optArgs = ole.ProcessOptArgs(ISWbemDateTime_GetVarDate_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000012, nil, optArgs...)
	return ole.Date(retVal.DateVal()).ToGoTime()
}

var ISWbemDateTime_SetVarDate_OptArgs = []string{
	"bIsLocal",
}

func (this *ISWbemDateTime) SetVarDate(dVarDate time.Time, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemDateTime_SetVarDate_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000013, []interface{}{dVarDate}, optArgs...)
	_ = retVal
}

var ISWbemDateTime_GetFileTime_OptArgs = []string{
	"bIsLocal",
}

func (this *ISWbemDateTime) GetFileTime(optArgs ...interface{}) string {
	optArgs = ole.ProcessOptArgs(ISWbemDateTime_GetFileTime_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000014, nil, optArgs...)
	return win32.BstrToStrAndFree(retVal.BstrValVal())
}

var ISWbemDateTime_SetFileTime_OptArgs = []string{
	"bIsLocal",
}

func (this *ISWbemDateTime) SetFileTime(strFileTime string, optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemDateTime_SetFileTime_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000015, []interface{}{strFileTime}, optArgs...)
	_ = retVal
}

