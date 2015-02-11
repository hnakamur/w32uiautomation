package w32uiautomation

// NOTE: APIs in this file are deprecated.

type ElemMatcherFunc func(element *IUIAutomationElement, walkder *IUIAutomationTreeWalker) (matched bool, err error)

func NewElemMatcherFuncWithName(name string) ElemMatcherFunc {
	return func(element *IUIAutomationElement, walkder *IUIAutomationTreeWalker) (matched bool, err error) {
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
	return func(element *IUIAutomationElement, walkder *IUIAutomationTreeWalker) (matched bool, err error) {
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
