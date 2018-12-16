package w32uiautomation

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
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

func (t StructureChangeType) ToString() string {
	switch t {
	case StructureChangeType_ChildAdded:
		return "StructureChangeType_ChildAdded"
	case StructureChangeType_ChildRemoved:
		return "StructureChangeType_ChildRemoved"
	case StructureChangeType_ChildrenInvalidated:
		return "StructureChangeType_ChildrenInvalidated"
	case StructureChangeType_ChildrenBulkAdded:
		return "StructureChangeType_ChildrenBulkAdded"
	case StructureChangeType_ChildrenBulkRemoved:
		return "StructureChangeType_ChildrenBulkRemoved"
	case StructureChangeType_ChildrenReordered:
		return "StructureChangeType_ChildrenReordered"
	default:
		panic(fmt.Errorf("Unknown StructureChangeType: %d", t))
	}
}

type IUIAutomationStructureChangedEventHandler struct {
	ole.IUnknown
	ref int32
}

type IUIAutomationStructureChangedEventHandlerVtbl struct {
	ole.IUnknownVtbl
	HandleStructureChangedEvent uintptr
}

var IID_IUIAutomationStructureChangedEventHandler = &ole.GUID{0xe81d1b4e, 0x11c5, 0x42f8, [8]byte{0x97, 0x54, 0xe7, 0x03, 0x6c, 0x79, 0xf0, 0x54}}

func (h *IUIAutomationStructureChangedEventHandler) VTable() *IUIAutomationStructureChangedEventHandlerVtbl {
	return (*IUIAutomationStructureChangedEventHandlerVtbl)(unsafe.Pointer(h.RawVTable))
}

func structureChangedEventHandler_queryInterface(this *ole.IUnknown, iid *ole.GUID, punk **ole.IUnknown) uint32 {
	*punk = nil
	if ole.IsEqualGUID(iid, ole.IID_IUnknown) ||
		ole.IsEqualGUID(iid, ole.IID_IDispatch) {
		structureChangedEventHandler_addRef(this)
		*punk = this
		return ole.S_OK
	}
	if ole.IsEqualGUID(iid, IID_IUIAutomationStructureChangedEventHandler) {
		structureChangedEventHandler_addRef(this)
		*punk = this
		return ole.S_OK
	}
	return ole.E_NOINTERFACE
}

func structureChangedEventHandler_addRef(this *ole.IUnknown) int32 {
	pthis := (*IUIAutomationStructureChangedEventHandler)(unsafe.Pointer(this))
	pthis.ref++
	return pthis.ref
}

func structureChangedEventHandler_release(this *ole.IUnknown) int32 {
	pthis := (*IUIAutomationStructureChangedEventHandler)(unsafe.Pointer(this))
	pthis.ref--
	return pthis.ref
}

func NewStructureChangedEventHandler(handlerFunc func(this *IUIAutomationStructureChangedEventHandler, sender *IUIAutomationElement, changeType StructureChangeType, runtimeId *ole.SAFEARRAY) syscall.Handle) IUIAutomationStructureChangedEventHandler {
	lpVtbl := &IUIAutomationStructureChangedEventHandlerVtbl{
		IUnknownVtbl: ole.IUnknownVtbl{
			QueryInterface: syscall.NewCallback(structureChangedEventHandler_queryInterface),
			AddRef:         syscall.NewCallback(structureChangedEventHandler_addRef),
			Release:        syscall.NewCallback(structureChangedEventHandler_release),
		},
		HandleStructureChangedEvent: syscall.NewCallback(handlerFunc),
	}
	return IUIAutomationStructureChangedEventHandler{
		IUnknown: ole.IUnknown{RawVTable: (*interface{})(unsafe.Pointer(lpVtbl))},
	}
}
