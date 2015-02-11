package w32uiautomation

import "time"

func WaitFindFirst(elem *IUIAutomationElement, scope TreeScope, condition *IUIAutomationCondition) (found *IUIAutomationElement, err error) {
	for {
		found, err := elem.FindFirst(scope, condition)
		if err != nil {
			return nil, err
		}
		if found != nil {
			return found, err
		}
		time.Sleep(100 * time.Millisecond)
	}
}
