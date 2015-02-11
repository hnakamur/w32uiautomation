package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/mattn/go-ole"
)

type IUIAutomationElement struct {
	ole.IUnknown
}

type IUIAutomationElementVtbl struct {
	ole.IUnknownVtbl
	SetFocus                        uintptr
	GetRuntimeId                    uintptr
	FindFirst                       uintptr
	FindAll                         uintptr
	FindFirstBuildCache             uintptr
	FindAllBuildCache               uintptr
	BuildUpdatedCache               uintptr
	GetCurrentPropertyValue         uintptr
	GetCurrentPropertyValueEx       uintptr
	GetCachedPropertyValue          uintptr
	GetCachedPropertyValueEx        uintptr
	GetCurrentPatternAs             uintptr
	GetCachedPatternAs              uintptr
	GetCurrentPattern               uintptr
	GetCachedPattern                uintptr
	GetCachedParent                 uintptr
	GetCachedChildren               uintptr
	Get_CurrentProcessId            uintptr
	Get_CurrentControlType          uintptr
	Get_CurrentLocalizedControlType uintptr
	Get_CurrentName                 uintptr
	Get_CurrentAcceleratorKey       uintptr
	Get_CurrentAccessKey            uintptr
	Get_CurrentHasKeyboardFocus     uintptr
	Get_CurrentIsKeyboardFocusable  uintptr
	Get_CurrentIsEnabled            uintptr
	Get_CurrentAutomationId         uintptr
	Get_CurrentClassName            uintptr
	Get_CurrentHelpText             uintptr
	Get_CurrentCulture              uintptr
	Get_CurrentIsControlElement     uintptr
	Get_CurrentIsContentElement     uintptr
	Get_CurrentIsPassword           uintptr
	Get_CurrentNativeWindowHandle   uintptr
	Get_CurrentItemType             uintptr
	Get_CurrentIsOffscreen          uintptr
	Get_CurrentOrientation          uintptr
	Get_CurrentFrameworkId          uintptr
	Get_CurrentIsRequiredForForm    uintptr
	Get_CurrentItemStatus           uintptr
	Get_CurrentBoundingRectangle    uintptr
	Get_CurrentLabeledBy            uintptr
	Get_CurrentAriaRole             uintptr
	Get_CurrentAriaProperties       uintptr
	Get_CurrentIsDataValidForForm   uintptr
	Get_CurrentControllerFor        uintptr
	Get_CurrentDescribedBy          uintptr
	Get_CurrentFlowsTo              uintptr
	Get_CurrentProviderDescription  uintptr
	Get_CachedProcessId             uintptr
	Get_CachedControlType           uintptr
	Get_CachedLocalizedControlType  uintptr
	Get_CachedName                  uintptr
	Get_CachedAcceleratorKey        uintptr
	Get_CachedAccessKey             uintptr
	Get_CachedHasKeyboardFocus      uintptr
	Get_CachedIsKeyboardFocusable   uintptr
	Get_CachedIsEnabled             uintptr
	Get_CachedAutomationId          uintptr
	Get_CachedClassName             uintptr
	Get_CachedHelpText              uintptr
	Get_CachedCulture               uintptr
	Get_CachedIsControlElement      uintptr
	Get_CachedIsContentElement      uintptr
	Get_CachedIsPassword            uintptr
	Get_CachedNativeWindowHandle    uintptr
	Get_CachedItemType              uintptr
	Get_CachedIsOffscreen           uintptr
	Get_CachedOrientation           uintptr
	Get_CachedFrameworkId           uintptr
	Get_CachedIsRequiredForForm     uintptr
	Get_CachedItemStatus            uintptr
	Get_CachedBoundingRectangle     uintptr
	Get_CachedLabeledBy             uintptr
	Get_CachedAriaRole              uintptr
	Get_CachedAriaProperties        uintptr
	Get_CachedIsDataValidForForm    uintptr
	Get_CachedControllerFor         uintptr
	Get_CachedDescribedBy           uintptr
	Get_CachedFlowsTo               uintptr
	Get_CachedProviderDescription   uintptr
	GetClickablePoint               uintptr
}

var IID_IUIAutomationElement = &ole.GUID{0xd22108aa, 0x8ac5, 0x49a5, [8]byte{0x83, 0x7b, 0x37, 0xbb, 0xb3, 0xd7, 0x59, 0x1e}}

func (elem *IUIAutomationElement) VTable() *IUIAutomationElementVtbl {
	return (*IUIAutomationElementVtbl)(unsafe.Pointer(elem.RawVTable))
}

func (elem *IUIAutomationElement) FindFirst(scope TreeScope, condition *IUIAutomationCondition) (found *IUIAutomationElement, err error) {
	return findFirst(elem, scope, condition)
}

func (elem *IUIAutomationElement) GetCurrentPattern(patternId PATTERNID) (*ole.IUnknown, error) {
	return getCurrentPattern(elem, patternId)
}

func (elem *IUIAutomationElement) Get_CurrentAutomationId() (string, error) {
	return get_CurrentAutomationId(elem)
}

func (elem *IUIAutomationElement) Get_CurrentCurrentClassName() (string, error) {
	return get_CurrentClassName(elem)
}

func (elem *IUIAutomationElement) Get_CurrentName() (string, error) {
	return get_CurrentName(elem)
}

func findFirst(elem *IUIAutomationElement, scope TreeScope, condition *IUIAutomationCondition) (found *IUIAutomationElement, err error) {
	hr, _, _ := syscall.Syscall6(
		elem.VTable().FindFirst,
		4,
		uintptr(unsafe.Pointer(elem)),
		uintptr(scope),
		uintptr(unsafe.Pointer(condition)),
		uintptr(unsafe.Pointer(&found)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getCurrentPattern(elem *IUIAutomationElement, patternId PATTERNID) (pattern *ole.IUnknown, err error) {
	hr, _, _ := syscall.Syscall(
		elem.VTable().GetCurrentPattern,
		3,
		uintptr(unsafe.Pointer(elem)),
		uintptr(patternId),
		uintptr(unsafe.Pointer(&pattern)))
	if hr != 0 {
		err = ole.NewError(hr)
		return
	}
	return
}

func get_CurrentAutomationId(elem *IUIAutomationElement) (id string, err error) {
	var bstrAutomationId *uint16
	hr, _, _ := syscall.Syscall(
		elem.VTable().Get_CurrentAutomationId,
		2,
		uintptr(unsafe.Pointer(elem)),
		uintptr(unsafe.Pointer(&bstrAutomationId)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
		return
	}
	id = ole.BstrToString(bstrAutomationId)
	return
}

func get_CurrentClassName(elem *IUIAutomationElement) (name string, err error) {
	var bstrName *uint16
	hr, _, _ := syscall.Syscall(
		elem.VTable().Get_CurrentClassName,
		2,
		uintptr(unsafe.Pointer(elem)),
		uintptr(unsafe.Pointer(&bstrName)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
		return
	}
	name = ole.BstrToString(bstrName)
	return
}

func get_CurrentName(elem *IUIAutomationElement) (name string, err error) {
	var bstrName *uint16
	hr, _, _ := syscall.Syscall(
		elem.VTable().Get_CurrentName,
		2,
		uintptr(unsafe.Pointer(elem)),
		uintptr(unsafe.Pointer(&bstrName)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
		return
	}
	name = ole.BstrToString(bstrName)
	return
}

//    typedef struct IUIAutomationElementVtbl
//    {
//        BEGIN_INTERFACE
//
//        HRESULT ( STDMETHODCALLTYPE *QueryInterface )(
//            __RPC__in IUIAutomationElement * This,
//            /* [in] */ __RPC__in REFIID riid,
//            /* [annotation][iid_is][out] */
//            __RPC__deref_out  void **ppvObject);
//
//        ULONG ( STDMETHODCALLTYPE *AddRef )(
//            __RPC__in IUIAutomationElement * This);
//
//        ULONG ( STDMETHODCALLTYPE *Release )(
//            __RPC__in IUIAutomationElement * This);
//
//        HRESULT ( STDMETHODCALLTYPE *SetFocus )(
//            __RPC__in IUIAutomationElement * This);
//
//        HRESULT ( STDMETHODCALLTYPE *GetRuntimeId )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt SAFEARRAY * *runtimeId);
//
//        HRESULT ( STDMETHODCALLTYPE *FindFirst )(
//            __RPC__in IUIAutomationElement * This,
//            /* [in] */ enum TreeScope scope,
//            /* [in] */ __RPC__in_opt IUIAutomationCondition *condition,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **found);
//
//        HRESULT ( STDMETHODCALLTYPE *FindAll )(
//            __RPC__in IUIAutomationElement * This,
//            /* [in] */ enum TreeScope scope,
//            /* [in] */ __RPC__in_opt IUIAutomationCondition *condition,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElementArray **found);
//
//        HRESULT ( STDMETHODCALLTYPE *FindFirstBuildCache )(
//            __RPC__in IUIAutomationElement * This,
//            /* [in] */ enum TreeScope scope,
//            /* [in] */ __RPC__in_opt IUIAutomationCondition *condition,
//            /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **found);
//
//        HRESULT ( STDMETHODCALLTYPE *FindAllBuildCache )(
//            __RPC__in IUIAutomationElement * This,
//            /* [in] */ enum TreeScope scope,
//            /* [in] */ __RPC__in_opt IUIAutomationCondition *condition,
//            /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElementArray **found);
//
//        HRESULT ( STDMETHODCALLTYPE *BuildUpdatedCache )(
//            __RPC__in IUIAutomationElement * This,
//            /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **updatedElement);
//
//        HRESULT ( STDMETHODCALLTYPE *GetCurrentPropertyValue )(
//            __RPC__in IUIAutomationElement * This,
//            /* [in] */ PROPERTYID propertyId,
//            /* [retval][out] */ __RPC__out VARIANT *retVal);
//
//        HRESULT ( STDMETHODCALLTYPE *GetCurrentPropertyValueEx )(
//            __RPC__in IUIAutomationElement * This,
//            /* [in] */ PROPERTYID propertyId,
//            /* [in] */ BOOL ignoreDefaultValue,
//            /* [retval][out] */ __RPC__out VARIANT *retVal);
//
//        HRESULT ( STDMETHODCALLTYPE *GetCachedPropertyValue )(
//            __RPC__in IUIAutomationElement * This,
//            /* [in] */ PROPERTYID propertyId,
//            /* [retval][out] */ __RPC__out VARIANT *retVal);
//
//        HRESULT ( STDMETHODCALLTYPE *GetCachedPropertyValueEx )(
//            __RPC__in IUIAutomationElement * This,
//            /* [in] */ PROPERTYID propertyId,
//            /* [in] */ BOOL ignoreDefaultValue,
//            /* [retval][out] */ __RPC__out VARIANT *retVal);
//
//        HRESULT ( STDMETHODCALLTYPE *GetCurrentPatternAs )(
//            __RPC__in IUIAutomationElement * This,
//            /* [in] */ PATTERNID patternId,
//            /* [in] */ __RPC__in REFIID riid,
//            /* [retval][iid_is][out] */ __RPC__deref_out_opt void **patternObject);
//
//        HRESULT ( STDMETHODCALLTYPE *GetCachedPatternAs )(
//            __RPC__in IUIAutomationElement * This,
//            /* [in] */ PATTERNID patternId,
//            /* [in] */ __RPC__in REFIID riid,
//            /* [retval][iid_is][out] */ __RPC__deref_out_opt void **patternObject);
//
//        HRESULT ( STDMETHODCALLTYPE *GetCurrentPattern )(
//            __RPC__in IUIAutomationElement * This,
//            /* [in] */ PATTERNID patternId,
//            /* [retval][out] */ __RPC__deref_out_opt IUnknown **patternObject);
//
//        HRESULT ( STDMETHODCALLTYPE *GetCachedPattern )(
//            __RPC__in IUIAutomationElement * This,
//            /* [in] */ PATTERNID patternId,
//            /* [retval][out] */ __RPC__deref_out_opt IUnknown **patternObject);
//
//        HRESULT ( STDMETHODCALLTYPE *GetCachedParent )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **parent);
//
//        HRESULT ( STDMETHODCALLTYPE *GetCachedChildren )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElementArray **children);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentProcessId )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out int *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentControlType )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out CONTROLTYPEID *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentLocalizedControlType )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentName )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentAcceleratorKey )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentAccessKey )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentHasKeyboardFocus )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentIsKeyboardFocusable )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentIsEnabled )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentAutomationId )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentClassName )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentHelpText )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentCulture )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out int *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentIsControlElement )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentIsContentElement )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentIsPassword )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentNativeWindowHandle )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt UIA_HWND *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentItemType )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentIsOffscreen )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentOrientation )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out enum OrientationType *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentFrameworkId )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentIsRequiredForForm )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentItemStatus )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentBoundingRectangle )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out RECT *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentLabeledBy )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentAriaRole )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentAriaProperties )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentIsDataValidForForm )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentControllerFor )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElementArray **retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentDescribedBy )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElementArray **retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentFlowsTo )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElementArray **retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CurrentProviderDescription )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedProcessId )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out int *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedControlType )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out CONTROLTYPEID *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedLocalizedControlType )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedName )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedAcceleratorKey )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedAccessKey )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedHasKeyboardFocus )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedIsKeyboardFocusable )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedIsEnabled )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedAutomationId )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedClassName )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedHelpText )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedCulture )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out int *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedIsControlElement )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedIsContentElement )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedIsPassword )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedNativeWindowHandle )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt UIA_HWND *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedItemType )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedIsOffscreen )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedOrientation )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out enum OrientationType *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedFrameworkId )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedIsRequiredForForm )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedItemStatus )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedBoundingRectangle )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out RECT *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedLabeledBy )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedAriaRole )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedAriaProperties )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedIsDataValidForForm )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__out BOOL *retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedControllerFor )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElementArray **retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedDescribedBy )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElementArray **retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedFlowsTo )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElementArray **retVal);
//
//        /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_CachedProviderDescription )(
//            __RPC__in IUIAutomationElement * This,
//            /* [retval][out] */ __RPC__deref_out_opt BSTR *retVal);
//
//        HRESULT ( STDMETHODCALLTYPE *GetClickablePoint )(
//            __RPC__in IUIAutomationElement * This,
//            /* [out] */ __RPC__out POINT *clickable,
//            /* [retval][out] */ __RPC__out BOOL *gotClickable);
//
//        END_INTERFACE
//    } IUIAutomationElementVtbl;
//
//    interface IUIAutomationElement
//    {
//        CONST_VTBL struct IUIAutomationElementVtbl *lpVtbl;
//    };
