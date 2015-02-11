package w32uiautomation

import (
	"unsafe"

	"github.com/mattn/go-ole"
)

func NewVariantString(s string) ole.VARIANT {
	return ole.NewVariant(
		ole.VT_BSTR,
		int64(uintptr(unsafe.Pointer(ole.SysAllocStringLen(s)))))
}
