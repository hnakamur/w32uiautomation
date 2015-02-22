package main

import (
	"fmt"
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
	calc, err := w32uiautomation.WaitFindFirst2(auto, root, w32uiautomation.TreeScope_Children, condition)
	if err != nil {
		return err
	}

	calcName, err := calc.Get_CurrentName()
	if err != nil {
		return err
	}
	fmt.Printf("calcName=%v\n", calcName)

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

	button, err := w32uiautomation.WaitFindFirst(calc,
		w32uiautomation.TreeScope_Subtree,
		condition)
	if err != nil {
		return err
	}
	return w32uiautomation.Invoke(button)
}

func main() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	err := runCalc()
	if err != nil {
		panic(err)
	}
}
