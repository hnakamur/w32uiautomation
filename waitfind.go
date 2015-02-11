package w32uiautomation

import (
	"fmt"
	"time"
	"unsafe"
)

func WaitFindFirst(elem *IUIAutomationElement, scope TreeScope, condition *IUIAutomationCondition) (found *IUIAutomationElement, err error) {
	fmt.Printf("WaitFindFirst start. elem=%v, scope=%x, condition=%v\n", elem, scope, condition)
	for {
		unkFound, err := elem.FindFirst(scope, condition)
		fmt.Printf("WaitFindFirst unkFound=%v\n", unkFound)
		if err != nil {
			return nil, err
		}
		if unkFound != nil {
			disp, err := unkFound.QueryInterface(IID_IUIAutomationElement)
			if err != nil {
				return nil, err
			}
			fmt.Printf("WaitFindFirst. disp=%v\n", disp)

			found = (*IUIAutomationElement)(unsafe.Pointer(disp))
			return found, nil
		}
		fmt.Printf("WaitFindFirst not found. sleep and try again. scope=%x, condition=%v\n", scope, condition)
		time.Sleep(100 * time.Millisecond)
	}
}
