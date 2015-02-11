package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/mattn/go-ole"
)

type IUIAutomationTreeWalker struct {
	ole.IUnknown
}

type IUIAutomationTreeWalkerVtbl struct {
	ole.IUnknownVtbl
	GetParentElement                    uintptr
	GetFirstChildElement                uintptr
	GetLastChildElement                 uintptr
	GetNextSiblingElement               uintptr
	GetPreviousSiblingElement           uintptr
	NormalizeElement                    uintptr
	GetParentElementBuildCache          uintptr
	GetFirstChildElementBuildCache      uintptr
	GetLastChildElementBuildCache       uintptr
	GetNextSiblingElementBuildCache     uintptr
	GetPreviousSiblingElementBuildCache uintptr
	NormalizeElementBuildCache          uintptr
	Get_Condition                       uintptr
}

func (w *IUIAutomationTreeWalker) VTable() *IUIAutomationTreeWalkerVtbl {
	return (*IUIAutomationTreeWalkerVtbl)(unsafe.Pointer(w.RawVTable))
}

func (w *IUIAutomationTreeWalker) GetFirstChildElement(element *IUIAutomationElement) (first *IUIAutomationElement, err error) {
	return getFirstChildElement(w, element)
}

func (w *IUIAutomationTreeWalker) GetNextSiblingElement(element *IUIAutomationElement) (next *IUIAutomationElement, err error) {
	return getNextSiblingElement(w, element)
}

func getFirstChildElement(w *IUIAutomationTreeWalker, element *IUIAutomationElement) (first *IUIAutomationElement, err error) {
	hr, _, _ := syscall.Syscall(
		w.VTable().GetFirstChildElement,
		3,
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(element)),
		uintptr(unsafe.Pointer(&first)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getNextSiblingElement(w *IUIAutomationTreeWalker, element *IUIAutomationElement) (next *IUIAutomationElement, err error) {
	hr, _, _ := syscall.Syscall(
		w.VTable().GetNextSiblingElement,
		3,
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(element)),
		uintptr(unsafe.Pointer(&next)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

//        HRESULT ( STDMETHODCALLTYPE *GetParentElement )(
//            __RPC__in IUIAutomationTreeWalker * This,
//            /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **parent);
//
//        HRESULT ( STDMETHODCALLTYPE *GetFirstChildElement )(
//            __RPC__in IUIAutomationTreeWalker * This,
//            /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **first);
//
//        HRESULT ( STDMETHODCALLTYPE *GetLastChildElement )(
//            __RPC__in IUIAutomationTreeWalker * This,
//            /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **last);
//
//        HRESULT ( STDMETHODCALLTYPE *GetNextSiblingElement )(
//            __RPC__in IUIAutomationTreeWalker * This,
//            /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **next);
//
//        HRESULT ( STDMETHODCALLTYPE *GetPreviousSiblingElement )(
//            __RPC__in IUIAutomationTreeWalker * This,
//            /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **previous);
//
//        HRESULT ( STDMETHODCALLTYPE *NormalizeElement )(
//            __RPC__in IUIAutomationTreeWalker * This,
//            /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **normalized);
//
//        HRESULT ( STDMETHODCALLTYPE *GetParentElementBuildCache )(
//            __RPC__in IUIAutomationTreeWalker * This,
//            /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//            /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **parent);
//
//        HRESULT ( STDMETHODCALLTYPE *GetFirstChildElementBuildCache )(
//            __RPC__in IUIAutomationTreeWalker * This,
//            /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//            /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **first);
//
//        HRESULT ( STDMETHODCALLTYPE *GetLastChildElementBuildCache )(
//            __RPC__in IUIAutomationTreeWalker * This,
//            /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//            /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **last);
//
//        HRESULT ( STDMETHODCALLTYPE *GetNextSiblingElementBuildCache )(
//            __RPC__in IUIAutomationTreeWalker * This,
//            /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//            /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **next);
//
//        HRESULT ( STDMETHODCALLTYPE *GetPreviousSiblingElementBuildCache )(
//            __RPC__in IUIAutomationTreeWalker * This,
//            /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//            /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **previous);
//
//        HRESULT ( STDMETHODCALLTYPE *NormalizeElementBuildCache )(
//            __RPC__in IUIAutomationTreeWalker * This,
//            /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//            /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **normalized);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_Condition )(
//            __RPC__in IUIAutomationTreeWalker * This,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **condition);
