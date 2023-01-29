package wmi

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
)

// 5AD4BF92-DAAB-11D3-B38F-00105A1F473A
var IID_ISWbemRefreshableItem = syscall.GUID{0x5AD4BF92, 0xDAAB, 0x11D3,
	[8]byte{0xB3, 0x8F, 0x00, 0x10, 0x5A, 0x1F, 0x47, 0x3A}}

type ISWbemRefreshableItem struct {
	ole.OleClient
}

func NewISWbemRefreshableItem(pDisp *win32.IDispatch, addRef bool, scoped bool) *ISWbemRefreshableItem {
	if pDisp == nil {
		return nil
	}
	p := &ISWbemRefreshableItem{ole.OleClient{pDisp}}
	if addRef {
		pDisp.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func ISWbemRefreshableItemFromVar(v ole.Variant) *ISWbemRefreshableItem {
	return NewISWbemRefreshableItem(v.IDispatch(), false, false)
}

func (this *ISWbemRefreshableItem) IID() *syscall.GUID {
	return &IID_ISWbemRefreshableItem
}

func (this *ISWbemRefreshableItem) GetIDispatch(addRef bool) *win32.IDispatch {
	if addRef {
		this.AddRef()
	}
	return this.IDispatch
}

func (this *ISWbemRefreshableItem) Index() int32 {
	retVal, _ := this.PropGet(0x00000001, nil)
	return retVal.LValVal()
}

func (this *ISWbemRefreshableItem) Refresher() *ISWbemRefresher {
	retVal, _ := this.PropGet(0x00000002, nil)
	return NewISWbemRefresher(retVal.IDispatch(), false, true)
}

func (this *ISWbemRefreshableItem) IsSet() bool {
	retVal, _ := this.PropGet(0x00000003, nil)
	return retVal.BoolValVal() != win32.VARIANT_FALSE
}

func (this *ISWbemRefreshableItem) Object() *ISWbemObjectEx {
	retVal, _ := this.PropGet(0x00000004, nil)
	return NewISWbemObjectEx(retVal.IDispatch(), false, true)
}

func (this *ISWbemRefreshableItem) ObjectSet() *ISWbemObjectSet {
	retVal, _ := this.PropGet(0x00000005, nil)
	return NewISWbemObjectSet(retVal.IDispatch(), false, true)
}

var ISWbemRefreshableItem_Remove_OptArgs = []string{
	"iFlags",
}

func (this *ISWbemRefreshableItem) Remove(optArgs ...interface{}) {
	optArgs = ole.ProcessOptArgs(ISWbemRefreshableItem_Remove_OptArgs, optArgs)
	retVal, _ := this.Call(0x00000006, nil, optArgs...)
	_ = retVal
}

