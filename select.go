package w32uiautomation

import "unsafe"

func Select(element *IUIAutomationElement) error {
	unknown, err := element.GetCurrentPattern(UIA_SelectionItemPatternId)
	if err != nil {
		return err
	}
	defer unknown.Release()

	disp, err := unknown.QueryInterface(IID_IUIAutomationSelectionItemPattern)
	if err != nil {
		return err
	}

	pattern := (*IUIAutomationSelectionItemPattern)(unsafe.Pointer(disp))
	defer pattern.Release()
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
