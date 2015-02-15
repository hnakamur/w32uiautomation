package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/mattn/go-ole"
)

type StructureChangeType uintptr

const (
	StructureChangeType_ChildAdded = iota
	StructureChangeType_ChildRemoved
	StructureChangeType_ChildrenInvalidated
	StructureChangeType_ChildrenBulkAdded
	StructureChangeType_ChildrenBulkRemoved
	StructureChangeType_ChildrenReordered
)

type IUIAutomationStructureChangedEventHandler struct {
	ole.IUnknown
}

type IUIAutomationStructureChangedEventHandlerVtbl struct {
	ole.IUnknownVtbl
	HandleStructureChangedEvent uintptr
}

var IID_IUIAutomationStructureChangedEventHandler = &ole.GUID{0xe81d1b4e, 0x11c5, 0x42f8, [8]byte{0x97, 0x54, 0xe7, 0x03, 0x6c, 0x79, 0xf0, 0x54}}

func (h *IUIAutomationStructureChangedEventHandler) VTable() *IUIAutomationStructureChangedEventHandlerVtbl {
	return (*IUIAutomationStructureChangedEventHandlerVtbl)(unsafe.Pointer(h.RawVTable))
}

func (h *IUIAutomationStructureChangedEventHandler) HandleStructureChangedEvent(sender *IUIAutomationElement, changeType StructureChangeType, runtimeId *ole.SAFEARRAY) error {
	return handleStructureChangedEvent(h, sender, changeType, runtimeId)
}

func handleStructureChangedEvent(h *IUIAutomationStructureChangedEventHandler, sender *IUIAutomationElement, changeType StructureChangeType, runtimeId *ole.SAFEARRAY) error {
	hr, _, _ := syscall.Syscall6(
		h.VTable().HandleStructureChangedEvent,
		4,
		uintptr(unsafe.Pointer(h)),
		uintptr(unsafe.Pointer(sender)),
		uintptr(changeType),
		uintptr(unsafe.Pointer(runtimeId)),
		0,
		0)
	if hr != 0 {
		return ole.NewError(hr)
	}
	return nil
}
