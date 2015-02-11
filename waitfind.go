package w32uiautomation

import (
	"fmt"
	"time"
)

func WaitFindFirst(elem *IUIAutomationElement, scope TreeScope, condition *IUIAutomationCondition) (found *IUIAutomationElement, err error) {
	fmt.Printf("WaitFindFirst start. elem=%v, scope=%x, condition=%v\n", elem, scope, condition)
	for {
		found, err = elem.FindFirst(scope, condition)
		if err != nil {
			return nil, err
		}
		if found != nil {
			return found, nil
		}
		fmt.Printf("WaitFindFirst not found. sleep and try again. scope=%x, condition=%v\n", scope, condition)
		time.Sleep(100 * time.Millisecond)
	}
}
