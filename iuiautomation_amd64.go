// +build amd64
package w32uiautomation

import (
	"syscall"
	"unsafe"

	ole "github.com/mattn/go-ole"
)

func createPropertyCondition(aut *IUIAutomation, propertyId PROPERTYID, value ole.VARIANT) (*IUIAutomationCondition, error) {
	var newCondition *IUIAutomationCondition
	v := VariantToUintptrArray(value)
	hr, _, _ := syscall.Syscall6(
		aut.VTable().CreatePropertyCondition,
		6,
		uintptr(unsafe.Pointer(aut)),
		uintptr(propertyId),
		v[0],
		v[1],
		v[2],
		uintptr(unsafe.Pointer(&newCondition)))
	if hr != 0 {
		return nil, ole.NewError(hr)
	}
	return newCondition, nil
}
