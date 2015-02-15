package main

import (
	"fmt"
	"os/exec"
	"syscall"
	"time"
	"unsafe"

	"github.com/hnakamur/w32uiautomation"
	"github.com/mattn/go-ole"
)

const (
	calculatorName          = "Calculator"
	clearButtonAutomationId = "81"
	twoButtonAutomationId   = "132"
	threeButtonAutomationId = "133"
	plusButtonAutomationId  = "93"
	equalButtonAutomationId = "121"
)

type MyHandler struct {
	w32uiautomation.IUIAutomationStructureChangedEventHandler
	refCount int32
}

func (h *MyHandler) QueryInterface(riid unsafe.Pointer, ppInterface **MyHandler) syscall.Handle {
	return myHandler_QueryInterface(h, riid, ppInterface)
}

func (h *MyHandler) AddRef() int32 {
	return myHandler_AddRef(h)
}

func (h *MyHandler) Release() int32 {
	return myHandler_Release(h)
}

func (h *MyHandler) HandleStructureChangedEvent(sender *w32uiautomation.IUIAutomationElement, changeType w32uiautomation.StructureChangeType, runtimeId *ole.SAFEARRAY) syscall.Handle {
	fmt.Printf("HandleStructureChangedEvent. h=%v, sender=%v, changeType=%v, runtimeId=%v\n", h, sender, changeType, runtimeId)
	return 0
}

func myHandler_QueryInterface(h *MyHandler, riid unsafe.Pointer, ppInterface **MyHandler) syscall.Handle {
	fmt.Println("myHandler_QueryInterface")
	h.AddRef()
	*ppInterface = h
	return 0
}

func myHandler_AddRef(h *MyHandler) int32 {
	fmt.Println("myHandler_AddRef")
	h.refCount += 1
	return h.refCount
}

func myHandler_Release(h *MyHandler) int32 {
	fmt.Println("myHandler_Release")
	h.refCount -= 1
	return h.refCount
}

func myHandler_HandleStructureChangedEvent(h *MyHandler, sender *w32uiautomation.IUIAutomationElement, changeType w32uiautomation.StructureChangeType, runtimeId *ole.SAFEARRAY) syscall.Handle {
	fmt.Printf("HandleStructureChangedEvent. h=%v, sender=%v, changeType=%v, runtimeId=%v\n", h, sender, changeType, runtimeId)
	return 0
}

var (
	myHandler_QueryInterfaceFunc              = myHandler_QueryInterface
	myHandler_AddRefFunc                      = myHandler_AddRef
	myHandler_ReleaseFunc                     = myHandler_Release
	myHandler_HandleStructureChangedEventFunc = myHandler_HandleStructureChangedEvent
)

func runCalc() error {
	err := exec.Command("control.exe", "/name", "Microsoft.Language").Start()
	if err != nil {
		return err
	}
	time.Sleep(time.Second)

	auto, err := w32uiautomation.NewUIAutomation()
	if err != nil {
		return err
	}

	root, err := auto.GetRootElement()
	if err != nil {
		return err
	}
	defer root.Release()

	condVal := w32uiautomation.NewVariantString("Language")
	fmt.Printf("condVal=%v, %s\n", condVal, condVal.ToString())
	condition, err := auto.CreatePropertyCondition(w32uiautomation.UIA_NamePropertyId, condVal)
	fmt.Printf("condition=%v, err=%v\n", condition, err)
	if err != nil {
		return err
	}
	calcWin, err := w32uiautomation.WaitFindFirst(root, w32uiautomation.TreeScope_Children, condition)
	fmt.Printf("calcWin=%v, err=%v\n", calcWin, err)
	if err != nil {
		return err
	}

	handler := new(MyHandler)
	vtable := w32uiautomation.IUIAutomationStructureChangedEventHandlerVtbl{
		ole.IUnknownVtbl{
			QueryInterface: uintptr(unsafe.Pointer(&myHandler_QueryInterfaceFunc)),
			AddRef:         uintptr(unsafe.Pointer(&myHandler_AddRefFunc)),
			Release:        uintptr(unsafe.Pointer(&myHandler_ReleaseFunc)),
		},
		uintptr(unsafe.Pointer(&myHandler_HandleStructureChangedEventFunc)),
	}
	handler.RawVTable = (*interface{})(unsafe.Pointer(&vtable))
	fmt.Println("Before AddStructureChangedEventHandler")
	err = auto.AddStructureChangedEventHandler(calcWin, w32uiautomation.TreeScope_Subtree, nil, (*w32uiautomation.IUIAutomationStructureChangedEventHandler)(unsafe.Pointer(handler)))
	if err != nil {
		return err
	}
	fmt.Println("After AddStructureChangedEventHandler")
	return nil
}

func main() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	err := runCalc()
	if err != nil {
		panic(err)
	}
}
