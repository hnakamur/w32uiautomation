package w32uiautomation

import (
	"fmt"
	"unsafe"
)

func Invoke(element *IUIAutomationElement) error {
	fmt.Printf("Invoke start. element=%v\n", element)
	unknown, err := element.GetCurrentPattern(UIA_InvokePatternId)
	if err != nil {
		return err
	}
	fmt.Printf("Invoke. unknown=%v\n", unknown)
	//defer unknown.Release()

	disp, err := unknown.QueryInterface(IID_IUIAutomationInvokePattern)
	if err != nil {
		return err
	}
	fmt.Printf("Invoke. disp=%v\n", disp)

	pattern := (*IUIAutomationInvokePattern)(unsafe.Pointer(disp))
	fmt.Printf("Invoke. pattern=%v\n", pattern)
	//defer pattern.Release()
	err = pattern.Invoke()
	if err != nil {
		return err
	}
	return nil
}
