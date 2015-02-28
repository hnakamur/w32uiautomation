// +build amd64
package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/mattn/go-ole"
)

func createPropertyCondition(aut *IUIAutomation, propertyId PROPERTYID, value ole.VARIANT) (newCondition *IUIAutomationCondition, err error) {
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
		err = ole.NewError(hr)
	}
	return
}
