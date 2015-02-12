package main

import (
	"fmt"
	"os/exec"

	wa "github.com/hnakamur/w32uiautomation"
	"github.com/mattn/go-ole"
)

func addLangJa() error {
	err := exec.Command("control.exe").Start()
	if err != nil {
		return err
	}

	auto, err := wa.NewUIAutomation()
	if err != nil {
		return err
	}

	root, err := auto.GetRootElement()
	if err != nil {
		return err
	}
	defer root.Release()

	condVal := wa.NewVariantString("Control Panel")
	fmt.Printf("condVal=%v, %s\n", condVal, condVal.ToString())
	condition, err := auto.CreatePropertyCondition(wa.UIA_NamePropertyId, condVal)
	fmt.Printf("condition=%v, err=%v\n", condition, err)
	if err != nil {
		return err
	}
	found, err := wa.WaitFindFirst(root, wa.TreeScope_Children, condition)
	fmt.Printf("found=%v, err=%v\n", found, err)
	if err != nil {
		return err
	}

	foundName, err := found.Get_CurrentName()
	if err != nil {
		return err
	}
	// I don't know why, but we get an empty string for foundName
	fmt.Printf("foundName=%v\n", foundName)

	foundAutomationId, err := found.Get_CurrentAutomationId()
	if err != nil {
		return err
	}
	fmt.Printf("foundAutomationId=%v\n", foundAutomationId)

	cPanel := found
	//cPanel, err := wa.FindFirstWithBreadthFirstSearch(auto, root,
	//	wa.NewElemMatcherFuncWithName(controlPanelWinName))
	//if err != nil {
	//	return err
	//}
	//cPanelName, err := cPanel.Get_CurrentName()
	//if err != nil {
	//	return err
	//}
	//// NOTE: Here we can get the actual name, "cPanelulator"
	//fmt.Printf("cPanel=%v, cPanelName=%v\n", cPanel, cPanelName)

	err = wa.WaitFindAndInvoke(auto, cPanel,
		wa.NewElemMatcherFuncWithName("Add a language"))
	if err != nil {
		return err
	}

	languageWin, err := wa.WaitFindFirstWithBreadthFirstSearch(
		auto, root, wa.NewElemMatcherFuncWithName("Language"))
	if err != nil {
		return err
	}
	fmt.Printf("languageWin=%v\n", languageWin)

	err = wa.WaitFindAndInvoke(auto, languageWin,
		wa.NewElemMatcherFuncWithName("Add a language"))
	if err != nil {
		return err
	}

	addLanguagesWin, err := wa.WaitFindFirstWithBreadthFirstSearch(
		auto, root, wa.NewElemMatcherFuncWithName("Add languages"))
	if err != nil {
		return err
	}
	fmt.Printf("addLanguagesWin=%v\n", addLanguagesWin)

	err = wa.WaitFindAndInvoke(auto, addLanguagesWin,
		wa.NewElemMatcherFuncWithName("Japanese"))
	if err != nil {
		return err
	}

	languageWin, err = wa.WaitFindFirstWithBreadthFirstSearch(
		auto, root, wa.NewElemMatcherFuncWithName("Language"))
	if err != nil {
		return err
	}
	fmt.Printf("languageWin=%v\n", languageWin)

	textElem, err := wa.WaitFindFirstWithBreadthFirstSearch(
		auto, languageWin, wa.NewElemMatcherFuncWithName("ja"))
	if err != nil {
		return err
	}
	fmt.Printf("textElem=%v\n", textElem)

	trueCondition, err := auto.CreateTrueCondition()
	if err != nil {
		return err
	}
	defer trueCondition.Release()

	walker, err := auto.CreateTreeWalker(trueCondition)
	if err != nil {
		return err
	}
	defer walker.Release()

	listItem, err := walker.GetParentElement(textElem)
	if err != nil {
		return err
	}

	err = wa.Invoke(listItem)
	if err != nil {
		return err
	}

	err = wa.WaitFindAndInvoke(auto, languageWin,
		wa.NewElemMatcherFuncWithAutomationId("MoveUpButton_TB"))
	if err != nil {
		return err
	}

	err = wa.WaitFindAndInvoke(auto, languageWin,
		wa.NewElemMatcherFuncWithAutomationId("Close"))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	err := addLangJa()
	if err != nil {
		panic(err)
	}
}
