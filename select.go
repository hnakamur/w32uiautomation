package w32uiautomation

import (
	"fmt"
	"unsafe"
)

func Select(element *IUIAutomationElement) error {
	fmt.Printf("Select start. element=%v\n", element)
	unknown, err := element.GetCurrentPattern(UIA_SelectionItemPatternId)
	if err != nil {
		return err
	}
	fmt.Printf("Select. unknown=%v\n", unknown)
	//defer unknown.Release()

	disp, err := unknown.QueryInterface(IID_IUIAutomationSelectionItemPattern)
	if err != nil {
		return err
	}
	fmt.Printf("Select. disp=%v\n", disp)

	pattern := (*IUIAutomationSelectionItemPattern)(unsafe.Pointer(disp))
	fmt.Printf("Select. pattern=%v\n", pattern)
	//defer pattern.Release()
	err = pattern.Select()
	if err != nil {
		return err
	}
	return nil
}

func WaitFindAndSelect(auto *IUIAutomation, element *IUIAutomationElement, matcher ElemMatcherFunc) error {
	element, err := WaitFindFirstWithBreadthFirstSearch(
		auto, element, matcher)
	if err != nil {
		return err
	}

	err = Select(element)
	if err != nil {
		return err
	}
	return nil
}
