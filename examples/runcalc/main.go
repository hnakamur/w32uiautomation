package main

import (
	"os/exec"

	"github.com/hnakamur/w32uiautomation"
	"github.com/mattn/go-ole"
)

const (
	calculatorName          = "Calculator"
	clearButtonAutomationId = "81"
	twoButtonAutomationId   = "132"
	threeButtonAutomationId = "133"
	plusButtonAutomationId  = "93"
	equalButtonAutomationId = "121"
)

func runCalc() error {
	err := exec.Command("calc.exe").Start()
	if err != nil {
		return err
	}

	auto, err := w32uiautomation.NewUIAutomation()
	if err != nil {
		return err
	}

	root, err := auto.GetRootElement()
	if err != nil {
		return err
	}
	defer root.Release()

	condVal := w32uiautomation.NewVariantString(calculatorName)
	condition, err := auto.CreatePropertyCondition(w32uiautomation.UIA_NamePropertyId, condVal)
	if err != nil {
		return err
	}
	found, err := w32uiautomation.WaitFindFirst(root, w32uiautomation.TreeScope_Children, condition)
	if err != nil {
		return err
	}

	foundName, err := found.Get_CurrentName()
	if err != nil {
		return err
	}
	// I don't know why, but we get an empty string for foundName
	//fmt.Printf("foundName=%v\n", foundName)

	//foundAutomationId, err := found.Get_CurrentAutomationId()
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("foundAutomationId=%v\n", foundAutomationId)

	calc := found
	//calc, err := w32uiautomation.FindFirstWithBreadthFirstSearch(auto, root,
	//	w32uiautomation.NewElemMatcherFuncWithName(calculatorName))
	//if err != nil {
	//	return err
	//}
	//calcName, err := calc.Get_CurrentName()
	//if err != nil {
	//	return err
	//}
	//// NOTE: Here we can get the actual name, "Calculator"
	//fmt.Printf("calc=%v, calcName=%v\n", calc, calcName)

	pushButton(auto, calc, clearButtonAutomationId)
	if err != nil {
		return err
	}

	pushButton(auto, calc, twoButtonAutomationId)
	if err != nil {
		return err
	}

	pushButton(auto, calc, plusButtonAutomationId)
	if err != nil {
		return err
	}

	pushButton(auto, calc, threeButtonAutomationId)
	if err != nil {
		return err
	}

	pushButton(auto, calc, equalButtonAutomationId)
	if err != nil {
		return err
	}

	return nil
}

func pushButton(auto *w32uiautomation.IUIAutomation, calc *w32uiautomation.IUIAutomationElement, automationId string) error {
	condition, err := auto.CreatePropertyCondition(
		w32uiautomation.UIA_AutomationIdPropertyId,
		w32uiautomation.NewVariantString(automationId))
	if err != nil {
		return err
	}

	//button, err := w32uiautomation.WaitFindFirst(calc,
	//	w32uiautomation.TreeScope_Subtree,
	//	condition)
	//fmt.Printf("pushButton button#1=%v\n", button)
	//if err != nil {
	//	return err
	//}
	// NOTE: I don't know why but I got an error if I call Invoke() with
	// button returned by WaitFindFirst above.
	//
	// Invoke start. element=&{{0x7183c068}}
	// Invoke. unknown=<nil>
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal 0xc0000005 code=0x0 addr=0x0 pc=0x4539b0]
	//err = w32uiautomation.Invoke(button)
	//if err != nil {
	//	return err
	//}

	button, err := w32uiautomation.WaitFindFirstWithBreadthFirstSearch(
		auto, calc,
		w32uiautomation.NewElemMatcherFuncWithAutomationId(automationId))
	if err != nil {
		return err
	}

	err = w32uiautomation.Invoke(button)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	err := runCalc()
	if err != nil {
		panic(err)
	}
}
