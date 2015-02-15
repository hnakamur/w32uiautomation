package w32uiautomation

import "fmt"

// NOTE: APIs in this file are deprecated.

type ElemMatcherFunc func(element *IUIAutomationElement, walker *IUIAutomationTreeWalker) (matched bool, err error)

func NewElemMatcherFuncWithName(name string) ElemMatcherFunc {
	return func(element *IUIAutomationElement, walker *IUIAutomationTreeWalker) (matched bool, err error) {
		n, err := element.Get_CurrentName()
		if err != nil {
			return
		}
		fmt.Printf("name matcher. n=%s\n", n)
		if n == name {
			matched = true
		}
		return
	}
}

func NewElemMatcherFuncWithAutomationId(automationId string) ElemMatcherFunc {
	return func(element *IUIAutomationElement, walker *IUIAutomationTreeWalker) (matched bool, err error) {
		id, err := element.Get_CurrentAutomationId()
		if err != nil {
			return
		}
		if id == automationId {
			matched = true
		}
		return
	}
}

func NewElemMatcherFuncWithChildName(name string) ElemMatcherFunc {
	fmt.Printf("NewElemMatcherFuncWithChildName. name=%s\n", name)
	return func(element *IUIAutomationElement, walker *IUIAutomationTreeWalker) (matched bool, err error) {
		fmt.Printf("Child matcher. element=%v\n", element)
		child, err := walker.GetFirstChildElement(element)
		if err != nil {
			return
		}
		fmt.Printf("Child matcher. first child=%v\n", child)

		for child != nil {
			var n string
			n, err = child.Get_CurrentName()
			if err != nil {
				break
			}
			fmt.Printf("Child matcher. child name=%s\n", n)

			if n == name {
				matched = true
				break
			}

			child, err = walker.GetNextSiblingElement(child)
			if err != nil {
				break
			}
			fmt.Printf("Child matcher. next sibling=%v\n", child)
		}
		return
	}
}
