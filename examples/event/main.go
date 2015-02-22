package main

import (
	"fmt"
	"os/exec"
	"syscall"
	"time"
	"unsafe"

	wa "github.com/hnakamur/w32uiautomation"
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

func HandleStructureChangedEvent(this *wa.IUIAutomationStructureChangedEventHandler, sender *wa.IUIAutomationElement, changeType wa.StructureChangeType, runtimeId *ole.SAFEARRAY) syscall.Handle {
	fmt.Printf("HandleStructureChangedEvent. this=%v, sender=%v, changeType=%s, runtimeId=%v\n", this, sender, changeType.ToString(), runtimeId)
	return 0
}

func addLanguage() error {
	err := exec.Command("control.exe", "/name", "Microsoft.Language").Start()
	if err != nil {
		return err
	}
	time.Sleep(time.Second)

	auto, err := wa.NewUIAutomation()
	if err != nil {
		return err
	}

	root, err := auto.GetRootElement()
	if err != nil {
		return err
	}
	defer root.Release()

	condVal := wa.NewVariantString("Language")
	fmt.Printf("condVal=%v, %s\n", condVal, condVal.ToString())
	condition, err := auto.CreatePropertyCondition(wa.UIA_NamePropertyId, condVal)
	fmt.Printf("condition=%v, err=%v\n", condition, err)
	if err != nil {
		return err
	}
	languageWin, err := wa.WaitFindFirst(auto, root, wa.TreeScope_Children, condition)
	fmt.Printf("languageWin=%v, err=%v\n", languageWin, err)
	if err != nil {
		return err
	}

	handler := wa.NewStructureChangedEventHandler(HandleStructureChangedEvent)

	fmt.Println("Before AddStructureChangedEventHandler")
	err = auto.AddStructureChangedEventHandler(languageWin, wa.TreeScope_Subtree, nil, (*wa.IUIAutomationStructureChangedEventHandler)(unsafe.Pointer(&handler)))
	if err != nil {
		return err
	}
	fmt.Println("After AddStructureChangedEventHandler")

	addALanguageLink, err := findElementByName(auto, languageWin, "Add a language")
	if err != nil {
		return err
	}
	fmt.Println(`Found "Add a language" link`)
	err = wa.Invoke(addALanguageLink)
	if err != nil {
		return err
	}
	fmt.Println(`Invoked "Add a language" link`)

	addLanguagesWin, err := findChildElementByName(auto, root, "Add languages")
	if err != nil {
		return err
	}
	fmt.Println(`Found "Add languages" window`)

	japaneseListItem, err := findElementByName(auto, addLanguagesWin, "Japanese")
	if err != nil {
		return err
	}
	fmt.Println(`Found "Japanese" listItem`)
	err = wa.Invoke(japaneseListItem)
	if err != nil {
		return err
	}
	fmt.Println(`Invoked "Japanese" listItem`)

	err = auto.RemoveStructureChangedEventHandler(languageWin, (*wa.IUIAutomationStructureChangedEventHandler)(unsafe.Pointer(&handler)))
	if err != nil {
		return err
	}
	fmt.Println("After RemoveStructureChangedEventHandler")

	return nil
}

func findChildElementByName(auto *wa.IUIAutomation, start *wa.IUIAutomationElement, elementName string) (*wa.IUIAutomationElement, error) {
	condVal := wa.NewVariantString(elementName)
	condition, err := auto.CreatePropertyCondition(wa.UIA_NamePropertyId, condVal)
	if err != nil {
		return nil, err
	}
	return wa.WaitFindFirst(auto, start, wa.TreeScope_Children, condition)
}

func findElementByName(auto *wa.IUIAutomation, start *wa.IUIAutomationElement, elementName string) (*wa.IUIAutomationElement, error) {
	condVal := wa.NewVariantString(elementName)
	condition, err := auto.CreatePropertyCondition(wa.UIA_NamePropertyId, condVal)
	if err != nil {
		return nil, err
	}
	return wa.WaitFindFirst(auto, start, wa.TreeScope_Subtree, condition)
}

func main() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	err := addLanguage()
	if err != nil {
		panic(err)
	}
}
