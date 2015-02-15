package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/mattn/go-ole"
)

type IUIAutomationSelectionItemPattern struct {
	ole.IUnknown
}

type IUIAutomationSelectionItemPatternVtbl struct {
	ole.IUnknownVtbl
	Select                        uintptr
	AddToSelection                uintptr
	RemoveFromSelection           uintptr
	Get_CurrentIsSelected         uintptr
	Get_CurrentSelectionContainer uintptr
	Get_CachedIsSelected          uintptr
	Get_CachedSelectionContainer  uintptr
}

//        HRESULT ( STDMETHODCALLTYPE *Select )(
//            __RPC__in IUIAutomationSelectionItemPattern * This);
//
//        HRESULT ( STDMETHODCALLTYPE *AddToSelection )(
//            __RPC__in IUIAutomationSelectionItemPattern * This);
//
//        HRESULT ( STDMETHODCALLTYPE *RemoveFromSelection )(
//            __RPC__in IUIAutomationSelectionItemPattern * This);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentIsSelected )(
//            __RPC__in IUIAutomationSelectionItemPattern * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentSelectionContainer )(
//            __RPC__in IUIAutomationSelectionItemPattern * This,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedIsSelected )(
//            __RPC__in IUIAutomationSelectionItemPattern * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedSelectionContainer )(
//            __RPC__in IUIAutomationSelectionItemPattern * This,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **retVal);

var IID_IUIAutomationSelectionItemPattern = &ole.GUID{0xa8efa66a, 0x0fda, 0x421a, [8]byte{0x91, 0x94, 0x38, 0x02, 0x1f, 0x35, 0x78, 0xea}}

func (pat *IUIAutomationSelectionItemPattern) VTable() *IUIAutomationSelectionItemPatternVtbl {
	return (*IUIAutomationSelectionItemPatternVtbl)(unsafe.Pointer(pat.RawVTable))
}

func (pat *IUIAutomationSelectionItemPattern) Select() error {
	return select_(pat)
}

func select_(pat *IUIAutomationSelectionItemPattern) error {
	hr, _, _ := syscall.Syscall(
		pat.VTable().Select,
		1,
		uintptr(unsafe.Pointer(pat)),
		0,
		0)
	if hr != 0 {
		return ole.NewError(hr)
	}
	return nil
}
