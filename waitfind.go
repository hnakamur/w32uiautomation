package w32uiautomation

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/mattn/go-ole"
)

func WaitFindFirst(auto *IUIAutomation, elem *IUIAutomationElement, scope TreeScope, condition *IUIAutomationCondition) (found *IUIAutomationElement, err error) {
	for {
		found, err = elem.FindFirst(scope, condition)
		if err != nil {
			return nil, err
		}
		if found != nil {
			return found, nil
		}

		fmt.Printf("Start waitChildAdded\n")
		waitChildAdded(auto, elem)
		fmt.Printf("received from waitChildAdded\n")
	}
}

func waitChildAdded(auto *IUIAutomation, elem *IUIAutomationElement) error {
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
		return err
	}
	fmt.Printf("AddStructureChangedEventHandler done\n")
	var m ole.Msg
	for waiting {
		ole.GetMessage(&m, 0, 0, 0)
		ole.DispatchMessage(&m)
	}
	err = auto.RemoveStructureChangedEventHandler(elem, &handler)
	if err != nil {
		return err
	}
	fmt.Printf("RemoveStructureChangedEventHandler done\n")
	return nil
}
