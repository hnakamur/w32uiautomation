package w32uiautomation

import "time"

// NOTE: APIs in this file are deprecated.

func FindToplevelElementWithName(auto *IUIAutomation, name string) (*IUIAutomationElement, error) {
	root, err := auto.GetRootElement()
	if err != nil {
		return nil, err
	}
	defer root.Release()

	condition, err := auto.CreateTrueCondition()
	if err != nil {
		return nil, err
	}
	defer condition.Release()

	walker, err := auto.CreateTreeWalker(condition)
	if err != nil {
		return nil, err
	}
	defer walker.Release()

	element, err := walker.GetFirstChildElement(root)
	if err != nil {
		return nil, err
	}
	for element != nil {
		elementName, err := element.Get_CurrentName()
		if err != nil {
			return nil, err
		}
		if elementName == name {
			return element, nil
		}

		element, err = walker.GetNextSiblingElement(element)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func WaitFindToplevelElementWithName(auto *IUIAutomation, name string) (*IUIAutomationElement, error) {
	for {
		found, err := FindToplevelElementWithName(auto, name)
		if err != nil {
			return nil, err
		}
		if found != nil {
			return found, err
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func WaitFindFirstWithBreadthFirstSearch(auto *IUIAutomation, start *IUIAutomationElement, matcher ElemMatcherFunc) (*IUIAutomationElement, error) {
	for {
		found, err := FindFirstWithBreadthFirstSearch(auto, start, matcher)
		if err != nil {
			return nil, err
		}
		if found != nil {
			return found, err
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func FindFirstWithBreadthFirstSearch(auto *IUIAutomation, start *IUIAutomationElement, matcher ElemMatcherFunc) (*IUIAutomationElement, error) {
	condition, err := auto.CreateTrueCondition()
	if err != nil {
		return nil, err
	}
	defer condition.Release()

	walker, err := auto.CreateTreeWalker(condition)
	if err != nil {
		return nil, err
	}
	defer walker.Release()

	return findFirstWithBreadthFirstSearchHelper(walker, start, matcher)
}

func findFirstWithBreadthFirstSearchHelper(walker *IUIAutomationTreeWalker, start *IUIAutomationElement, matcher ElemMatcherFunc) (*IUIAutomationElement, error) {
	element, err := findFirstInSiblings(walker, start, matcher)
	if err != nil {
		return nil, err
	}
	if element != nil {
		return element, nil
	}

	for parent := start; parent != nil; {
		child, err := walker.GetFirstChildElement(parent)
		if err != nil {
			return nil, err
		}

		if child != nil {
			element, err := findFirstWithBreadthFirstSearchHelper(walker, child, matcher)
			if err != nil {
				return nil, err
			}
			if element != nil {
				return element, nil
			}
		}

		parent, err = walker.GetNextSiblingElement(parent)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func findFirstInSiblings(walker *IUIAutomationTreeWalker, start *IUIAutomationElement, matcher ElemMatcherFunc) (*IUIAutomationElement, error) {
	element := start
	for element != nil {
		matched, err := matcher(element, walker)
		if err != nil {
			return nil, err
		}
		if matched {
			return element, nil
		}

		element, err = walker.GetNextSiblingElement(element)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}
