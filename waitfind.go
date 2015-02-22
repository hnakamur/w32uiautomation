package w32uiautomation

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"

	"github.com/mattn/go-ole"
)

func WaitFindFirst(elem *IUIAutomationElement, scope TreeScope, condition *IUIAutomationCondition) (found *IUIAutomationElement, err error) {

	for {
		found, err = elem.FindFirst(scope, condition)
		if err != nil {
			return nil, err
		}
		if found != nil {
			return found, nil
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func WaitFindFirst2(auto *IUIAutomation, elem *IUIAutomationElement, scope TreeScope, condition *IUIAutomationCondition) (found *IUIAutomationElement, err error) {
	for {
		found, err = elem.FindFirst(scope, condition)
		if err != nil {
			return nil, err
		}
		if found != nil {
			return found, nil
		}

		c := make(chan struct{})
		errc := make(chan error)
		fmt.Printf("Start waitChildAdded\n")
		go waitChildAdded(auto, elem, c, errc)
		<-c
		fmt.Printf("received from waitChildAdded\n")
	}
}

func waitChildAdded(auto *IUIAutomation, elem *IUIAutomationElement, c chan struct{}, errc chan error) {
	waiting := true
	handler := NewStructureChangedEventHandler(nil)
	lpVtbl := (*IUIAutomationStructureChangedEventHandlerVtbl)(unsafe.Pointer(handler.IUnknown.RawVTable))
	lpVtbl.HandleStructureChangedEvent = syscall.NewCallback(func(this *IUIAutomationStructureChangedEventHandler, sender *IUIAutomationElement, changeType StructureChangeType, runtimeId *ole.SAFEARRAY) syscall.Handle {
		switch changeType {
		case StructureChangeType_ChildAdded, StructureChangeType_ChildrenBulkAdded:
			fmt.Printf("Got event with changeType=%v\n", changeType)
			waiting = false
		}
		return ole.S_OK
	})
	err := auto.AddStructureChangedEventHandler(elem, TreeScope_Subtree, nil, &handler)
	if err != nil {
		errc <- err
	}
	fmt.Printf("AddStructureChangedEventHandler done\n")
	var m ole.Msg
	for waiting {
		ole.GetMessage(&m, 0, 0, 0)
		ole.DispatchMessage(&m)
	}
	err = auto.RemoveStructureChangedEventHandler(elem, &handler)
	if err != nil {
		errc <- err
	}
	fmt.Printf("RemoveStructureChangedEventHandler done\n")
	c <- struct{}{}
}
