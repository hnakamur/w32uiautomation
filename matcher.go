package w32uiautomation

// NOTE: APIs in this file may be deleted in the future.

type ElemMatcherFunc func(element *IUIAutomationElement, walker *IUIAutomationTreeWalker) (matched bool, err error)

func NewElemMatcherFuncWithName(name string) ElemMatcherFunc {
	return func(element *IUIAutomationElement, walker *IUIAutomationTreeWalker) (matched bool, err error) {
		n, err := element.Get_CurrentName()
		if err != nil {
			return
		}
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
	return func(element *IUIAutomationElement, walker *IUIAutomationTreeWalker) (matched bool, err error) {
		child, err := walker.GetFirstChildElement(element)
		if err != nil {
			return
		}

		for child != nil {
			var n string
			n, err = child.Get_CurrentName()
			if err != nil {
				break
			}

			if n == name {
				matched = true
				break
			}

			child, err = walker.GetNextSiblingElement(child)
			if err != nil {
				break
			}
		}
		return
	}
}
