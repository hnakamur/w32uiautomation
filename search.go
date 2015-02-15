package w32uiautomation

import "time"

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
	walker, err := NewTreeWalker(auto)
	if err != nil {
		return nil, err
	}
	defer walker.Release()

	return findFirstWithBreadthFirstSearchHelper(walker, start, matcher)
}

func findFirstWithBreadthFirstSearchHelper(walker *IUIAutomationTreeWalker, start *IUIAutomationElement, matcher ElemMatcherFunc) (*IUIAutomationElement, error) {
	//fmt.Printf("findFirstWithBreadthFirstSearchHelper start=%v\n", start)
	element, err := findFirstInSiblings(walker, start, matcher)
	if err != nil {
		return nil, err
	}
	if element != nil {
		//fmt.Printf("findFirstWithBreadthFirstSearchHelper found#1=%v\n", element)
		return element, nil
	}

	for parent := start; parent != nil; {
		child, err := walker.GetFirstChildElement(parent)
		if err != nil {
			return nil, err
		}

		if child != nil {
			//fmt.Printf("findFirstWithBreadthFirstSearchHelper first chid=%v\n", child)
			element, err := findFirstWithBreadthFirstSearchHelper(walker, child, matcher)
			if err != nil {
				return nil, err
			}
			if element != nil {
				//fmt.Printf("findFirstWithBreadthFirstSearchHelper found#2=%v\n", element)
				return element, nil
			}
		}

		parent, err = walker.GetNextSiblingElement(parent)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("findFirstWithBreadthFirstSearchHelper next parent=%v\n", parent)
	}
	//fmt.Printf("findFirstWithBreadthFirstSearchHelper exiting\n")
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
		//fmt.Printf("findFirstInSiblings next sibling=%v\n", element)
	}
	return nil, nil
}
