package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/mattn/go-ole"
)

type IUIAutomation struct {
	ole.IUnknown
}

type IUIAutomationVtbl struct {
	ole.IUnknownVtbl
	CompareElements                           uintptr
	CompareRuntimeIds                         uintptr
	GetRootElement                            uintptr
	ElementFromHandle                         uintptr
	ElementFromPoint                          uintptr
	GetFocusedElement                         uintptr
	GetRootElementBuildCache                  uintptr
	ElementFromHandleBuildCache               uintptr
	ElementFromPointBuildCache                uintptr
	GetFocusedElementBuildCache               uintptr
	CreateTreeWalker                          uintptr
	Get_ControlViewWalker                     uintptr
	Get_ContentViewWalker                     uintptr
	Get_RawViewWalker                         uintptr
	Get_RawViewCondition                      uintptr
	Get_ControlViewCondition                  uintptr
	Get_ContentViewCondition                  uintptr
	CreateCacheRequest                        uintptr
	CreateTrueCondition                       uintptr
	CreateFalseCondition                      uintptr
	CreatePropertyCondition                   uintptr
	CreatePropertyConditionEx                 uintptr
	CreateAndCondition                        uintptr
	CreateAndConditionFromArray               uintptr
	CreateAndConditionFromNativeArray         uintptr
	CreateOrCondition                         uintptr
	CreateOrConditionFromArray                uintptr
	CreateOrConditionFromNativeArray          uintptr
	CreateNotCondition                        uintptr
	AddAutomationEventHandler                 uintptr
	RemoveAutomationEventHandler              uintptr
	AddPropertyChangedEventHandlerNativeArray uintptr
	AddPropertyChangedEventHandler            uintptr
	RemovePropertyChangedEventHandler         uintptr
	AddStructureChangedEventHandler           uintptr
	RemoveStructureChangedEventHandler        uintptr
	AddFocusChangedEventHandler               uintptr
	RemoveFocusChangedEventHandler            uintptr
	RemoveAllEventHandlers                    uintptr
	IntNativeArrayToSafeArray                 uintptr
	IntSafeArrayToNativeArray                 uintptr
	RectToVariant                             uintptr
	VariantToRect                             uintptr
	SafeArrayToRectNativeArray                uintptr
	CreateProxyFactoryEntry                   uintptr
	Get_ProxyFactoryMapping                   uintptr
	GetPropertyProgrammaticName               uintptr
	GetPatternProgrammaticName                uintptr
	PollForPotentialSupportedPatterns         uintptr
	PollForPotentialSupportedProperties       uintptr
	CheckNotSupported                         uintptr
	Get_ReservedNotSupportedValue             uintptr
	Get_ReservedMixedAttributeValue           uintptr
	ElementFromIAccessible                    uintptr
	ElementFromIAccessibleBuildCache          uintptr
}

// var modole32 = syscall.NewLazyDLL("ole32.dll")
// var procCoCreateInstance = modole32.NewProc("CoCreateInstance")
//
// func CoCreateInstance(clsid *ole.GUID, unkOuter *ole.IUnknown, clsContext uint32, iid *ole.GUID) (unk *ole.IUnknown, err error) {
// 	hr, _, _ := procCoCreateInstance.Call(
// 		uintptr(unsafe.Pointer(clsid)),
// 		uintptr(unsafe.Pointer(unkOuter)),
// 		uintptr(clsContext),
// 		uintptr(unsafe.Pointer(iid)),
// 		uintptr(unsafe.Pointer(&unk)))
// 	if hr != 0 {
// 		err = ole.NewError(hr)
// 	}
// 	return
// }

var CLSID_CUIAutomation = &ole.GUID{0xff48dba4, 0x60ef, 0x4201, [8]byte{0xaa, 0x87, 0x54, 0x10, 0x3e, 0xef, 0x59, 0x4e}}

var IID_IUIAutomation = &ole.GUID{0x30cbe57d, 0xd9d0, 0x452a, [8]byte{0xab, 0x13, 0x7a, 0xc5, 0xac, 0x48, 0x25, 0xee}}

func (v *IUIAutomation) VTable() *IUIAutomationVtbl {
	return (*IUIAutomationVtbl)(unsafe.Pointer(v.RawVTable))
}

func NewUIAutomation() (*IUIAutomation, error) {
	auto, err := ole.CreateInstance(CLSID_CUIAutomation, IID_IUIAutomation)
	//auto, err := CoCreateInstance(CLSID_CUIAutomation, nil, ole.CLSCTX_INPROC_SERVER, IID_IUIAutomation)
	if err != nil {
		return nil, err
	}
	return (*IUIAutomation)(unsafe.Pointer(auto)), nil
}

func (auto *IUIAutomation) CompareElements(el1, el2 *IUIAutomation) (areSame bool, err error) {
	return compareElements(auto, el1, el2)
}

func (auto *IUIAutomation) GetRootElement() (root *IUIAutomationElement, err error) {
	return getRootElement(auto)
}

func (auto *IUIAutomation) CreateTreeWalker(condition *IUIAutomationCondition) (walker *IUIAutomationTreeWalker, err error) {
	return createTreeWalker(auto, condition)
}

func (auto *IUIAutomation) CreateTrueCondition() (newCondition *IUIAutomationCondition, err error) {
	return createTrueCondition(auto)
}

func (auto *IUIAutomation) CreateAndCondition(condition1, condition2 *IUIAutomationCondition) (newCondition *IUIAutomationCondition, err error) {
	return createAndCondition(auto, condition1, condition2)
}

func (auto *IUIAutomation) CreatePropertyCondition(propertyId PROPERTYID, value ole.VARIANT) (newCondition *IUIAutomationCondition, err error) {
	return createPropertyCondition(auto, propertyId, value)
}

func compareElements(auto *IUIAutomation, el1, el2 *IUIAutomation) (areSame bool, err error) {
	hr, _, _ := syscall.Syscall6(
		auto.VTable().CompareElements,
		4,
		uintptr(unsafe.Pointer(auto)),
		uintptr(unsafe.Pointer(el1)),
		uintptr(unsafe.Pointer(el2)),
		uintptr(unsafe.Pointer(&areSame)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getRootElement(auto *IUIAutomation) (root *IUIAutomationElement, err error) {
	hr, _, _ := syscall.Syscall(
		auto.VTable().GetRootElement,
		3,
		uintptr(unsafe.Pointer(auto)),
		uintptr(unsafe.Pointer(&root)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func createTreeWalker(auto *IUIAutomation, condition *IUIAutomationCondition) (walker *IUIAutomationTreeWalker, err error) {
	hr, _, _ := syscall.Syscall(
		auto.VTable().CreateTreeWalker,
		3,
		uintptr(unsafe.Pointer(auto)),
		uintptr(unsafe.Pointer(condition)),
		uintptr(unsafe.Pointer(&walker)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func createTrueCondition(auto *IUIAutomation) (newCondition *IUIAutomationCondition, err error) {
	hr, _, _ := syscall.Syscall(
		auto.VTable().CreateTrueCondition,
		2,
		uintptr(unsafe.Pointer(auto)),
		uintptr(unsafe.Pointer(&newCondition)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func createAndCondition(auto *IUIAutomation, condition1, condition2 *IUIAutomationCondition) (newCondition *IUIAutomationCondition, err error) {
	hr, _, _ := syscall.Syscall6(
		auto.VTable().CreateAndCondition,
		4,
		uintptr(unsafe.Pointer(auto)),
		uintptr(unsafe.Pointer(condition1)),
		uintptr(unsafe.Pointer(condition2)),
		uintptr(unsafe.Pointer(&newCondition)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

//         HRESULT ( STDMETHODCALLTYPE *CompareElements )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationElement *el1,
//             /* [in] */ __RPC__in_opt IUIAutomationElement *el2,
//             /* [retval][out] */ __RPC__out BOOL *areSame);

//         HRESULT ( STDMETHODCALLTYPE *CompareElements )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationElement *el1,
//             /* [in] */ __RPC__in_opt IUIAutomationElement *el2,
//             /* [retval][out] */ __RPC__out BOOL *areSame);
//
//         HRESULT ( STDMETHODCALLTYPE *CompareRuntimeIds )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in SAFEARRAY * runtimeId1,
//             /* [in] */ __RPC__in SAFEARRAY * runtimeId2,
//             /* [retval][out] */ __RPC__out BOOL *areSame);
//
//         HRESULT ( STDMETHODCALLTYPE *GetRootElement )(
//             __RPC__in IUIAutomation * This,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **root);
//
//         HRESULT ( STDMETHODCALLTYPE *ElementFromHandle )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in UIA_HWND hwnd,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **element);
//
//         HRESULT ( STDMETHODCALLTYPE *ElementFromPoint )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ POINT pt,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **element);
//
//         HRESULT ( STDMETHODCALLTYPE *GetFocusedElement )(
//             __RPC__in IUIAutomation * This,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **element);
//
//         HRESULT ( STDMETHODCALLTYPE *GetRootElementBuildCache )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **root);
//
//         HRESULT ( STDMETHODCALLTYPE *ElementFromHandleBuildCache )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in UIA_HWND hwnd,
//             /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **element);
//
//         HRESULT ( STDMETHODCALLTYPE *ElementFromPointBuildCache )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ POINT pt,
//             /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **element);
//
//         HRESULT ( STDMETHODCALLTYPE *GetFocusedElementBuildCache )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **element);
//
//         HRESULT ( STDMETHODCALLTYPE *CreateTreeWalker )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationCondition *pCondition,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationTreeWalker **walker);
//
//         /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_ControlViewWalker )(
//             __RPC__in IUIAutomation * This,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationTreeWalker **walker);
//
//         /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_ContentViewWalker )(
//             __RPC__in IUIAutomation * This,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationTreeWalker **walker);
//
//         /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_RawViewWalker )(
//             __RPC__in IUIAutomation * This,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationTreeWalker **walker);
//
//         /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_RawViewCondition )(
//             __RPC__in IUIAutomation * This,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **condition);
//
//         /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_ControlViewCondition )(
//             __RPC__in IUIAutomation * This,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **condition);
//
//         /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_ContentViewCondition )(
//             __RPC__in IUIAutomation * This,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **condition);
//
//         HRESULT ( STDMETHODCALLTYPE *CreateCacheRequest )(
//             __RPC__in IUIAutomation * This,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCacheRequest **cacheRequest);
//
//         HRESULT ( STDMETHODCALLTYPE *CreateTrueCondition )(
//             __RPC__in IUIAutomation * This,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **newCondition);
//
//         HRESULT ( STDMETHODCALLTYPE *CreateFalseCondition )(
//             __RPC__in IUIAutomation * This,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **newCondition);
//
//         HRESULT ( STDMETHODCALLTYPE *CreatePropertyCondition )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ PROPERTYID propertyId,
//             /* [in] */ VARIANT value,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **newCondition);
//
//         HRESULT ( STDMETHODCALLTYPE *CreatePropertyConditionEx )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ PROPERTYID propertyId,
//             /* [in] */ VARIANT value,
//             /* [in] */ enum PropertyConditionFlags flags,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **newCondition);
//
//         HRESULT ( STDMETHODCALLTYPE *CreateAndCondition )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationCondition *condition1,
//             /* [in] */ __RPC__in_opt IUIAutomationCondition *condition2,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **newCondition);
//
//         HRESULT ( STDMETHODCALLTYPE *CreateAndConditionFromArray )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt SAFEARRAY * conditions,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **newCondition);
//
//         HRESULT ( STDMETHODCALLTYPE *CreateAndConditionFromNativeArray )(
//             __RPC__in IUIAutomation * This,
//             /* [size_is][in] */ __RPC__in_ecount_full(conditionCount) IUIAutomationCondition **conditions,
//             /* [in] */ int conditionCount,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **newCondition);
//
//         HRESULT ( STDMETHODCALLTYPE *CreateOrCondition )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationCondition *condition1,
//             /* [in] */ __RPC__in_opt IUIAutomationCondition *condition2,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **newCondition);
//
//         HRESULT ( STDMETHODCALLTYPE *CreateOrConditionFromArray )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt SAFEARRAY * conditions,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **newCondition);
//
//         HRESULT ( STDMETHODCALLTYPE *CreateOrConditionFromNativeArray )(
//             __RPC__in IUIAutomation * This,
//             /* [size_is][in] */ __RPC__in_ecount_full(conditionCount) IUIAutomationCondition **conditions,
//             /* [in] */ int conditionCount,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **newCondition);
//
//         HRESULT ( STDMETHODCALLTYPE *CreateNotCondition )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationCondition *condition,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationCondition **newCondition);
//
//         HRESULT ( STDMETHODCALLTYPE *AddAutomationEventHandler )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ EVENTID eventId,
//             /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//             /* [in] */ enum TreeScope scope,
//             /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//             /* [in] */ __RPC__in_opt IUIAutomationEventHandler *handler);
//
//         HRESULT ( STDMETHODCALLTYPE *RemoveAutomationEventHandler )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ EVENTID eventId,
//             /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//             /* [in] */ __RPC__in_opt IUIAutomationEventHandler *handler);
//
//         HRESULT ( STDMETHODCALLTYPE *AddPropertyChangedEventHandlerNativeArray )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//             /* [in] */ enum TreeScope scope,
//             /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//             /* [in] */ __RPC__in_opt IUIAutomationPropertyChangedEventHandler *handler,
//             /* [size_is][in] */ __RPC__in_ecount_full(propertyCount) PROPERTYID *propertyArray,
//             /* [in] */ int propertyCount);
//
//         HRESULT ( STDMETHODCALLTYPE *AddPropertyChangedEventHandler )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//             /* [in] */ enum TreeScope scope,
//             /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//             /* [in] */ __RPC__in_opt IUIAutomationPropertyChangedEventHandler *handler,
//             /* [in] */ __RPC__in SAFEARRAY * propertyArray);
//
//         HRESULT ( STDMETHODCALLTYPE *RemovePropertyChangedEventHandler )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//             /* [in] */ __RPC__in_opt IUIAutomationPropertyChangedEventHandler *handler);
//
//         HRESULT ( STDMETHODCALLTYPE *AddStructureChangedEventHandler )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//             /* [in] */ enum TreeScope scope,
//             /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//             /* [in] */ __RPC__in_opt IUIAutomationStructureChangedEventHandler *handler);
//
//         HRESULT ( STDMETHODCALLTYPE *RemoveStructureChangedEventHandler )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationElement *element,
//             /* [in] */ __RPC__in_opt IUIAutomationStructureChangedEventHandler *handler);
//
//         HRESULT ( STDMETHODCALLTYPE *AddFocusChangedEventHandler )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//             /* [in] */ __RPC__in_opt IUIAutomationFocusChangedEventHandler *handler);
//
//         HRESULT ( STDMETHODCALLTYPE *RemoveFocusChangedEventHandler )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationFocusChangedEventHandler *handler);
//
//         HRESULT ( STDMETHODCALLTYPE *RemoveAllEventHandlers )(
//             __RPC__in IUIAutomation * This);
//
//         HRESULT ( STDMETHODCALLTYPE *IntNativeArrayToSafeArray )(
//             __RPC__in IUIAutomation * This,
//             /* [size_is][in] */ __RPC__in_ecount_full(arrayCount) int *array,
//             /* [in] */ int arrayCount,
//             /* [retval][out] */ __RPC__deref_out_opt SAFEARRAY * *safeArray);
//
//         HRESULT ( STDMETHODCALLTYPE *IntSafeArrayToNativeArray )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in SAFEARRAY * intArray,
//             /* [size_is][size_is][out] */ __RPC__deref_out_ecount_full_opt(*arrayCount) int **array,
//             /* [retval][out] */ __RPC__out int *arrayCount);
//
//         HRESULT ( STDMETHODCALLTYPE *RectToVariant )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ RECT rc,
//             /* [retval][out] */ __RPC__out VARIANT *var);
//
//         HRESULT ( STDMETHODCALLTYPE *VariantToRect )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ VARIANT var,
//             /* [retval][out] */ __RPC__out RECT *rc);
//
//         HRESULT ( STDMETHODCALLTYPE *SafeArrayToRectNativeArray )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in SAFEARRAY * rects,
//             /* [size_is][size_is][out] */ __RPC__deref_out_ecount_full_opt(*rectArrayCount) RECT **rectArray,
//             /* [retval][out] */ __RPC__out int *rectArrayCount);
//
//         HRESULT ( STDMETHODCALLTYPE *CreateProxyFactoryEntry )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationProxyFactory *factory,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationProxyFactoryEntry **factoryEntry);
//
//         /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_ProxyFactoryMapping )(
//             __RPC__in IUIAutomation * This,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationProxyFactoryMapping **factoryMapping);
//
//         HRESULT ( STDMETHODCALLTYPE *GetPropertyProgrammaticName )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ PROPERTYID property,
//             /* [retval][out] */ __RPC__deref_out_opt BSTR *name);
//
//         HRESULT ( STDMETHODCALLTYPE *GetPatternProgrammaticName )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ PATTERNID pattern,
//             /* [retval][out] */ __RPC__deref_out_opt BSTR *name);
//
//         HRESULT ( STDMETHODCALLTYPE *PollForPotentialSupportedPatterns )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationElement *pElement,
//             /* [out] */ __RPC__deref_out_opt SAFEARRAY * *patternIds,
//             /* [out] */ __RPC__deref_out_opt SAFEARRAY * *patternNames);
//
//         HRESULT ( STDMETHODCALLTYPE *PollForPotentialSupportedProperties )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IUIAutomationElement *pElement,
//             /* [out] */ __RPC__deref_out_opt SAFEARRAY * *propertyIds,
//             /* [out] */ __RPC__deref_out_opt SAFEARRAY * *propertyNames);
//
//         HRESULT ( STDMETHODCALLTYPE *CheckNotSupported )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ VARIANT value,
//             /* [retval][out] */ __RPC__out BOOL *isNotSupported);
//
//         /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_ReservedNotSupportedValue )(
//             __RPC__in IUIAutomation * This,
//             /* [retval][out] */ __RPC__deref_out_opt IUnknown **notSupportedValue);
//
//         /* [propget] */ HRESULT ( STDMETHODCALLTYPE *get_ReservedMixedAttributeValue )(
//             __RPC__in IUIAutomation * This,
//             /* [retval][out] */ __RPC__deref_out_opt IUnknown **mixedAttributeValue);
//
//         HRESULT ( STDMETHODCALLTYPE *ElementFromIAccessible )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IAccessible *accessible,
//             /* [in] */ int childId,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **element);
//
//         HRESULT ( STDMETHODCALLTYPE *ElementFromIAccessibleBuildCache )(
//             __RPC__in IUIAutomation * This,
//             /* [in] */ __RPC__in_opt IAccessible *accessible,
//             /* [in] */ int childId,
//             /* [in] */ __RPC__in_opt IUIAutomationCacheRequest *cacheRequest,
//             /* [retval][out] */ __RPC__deref_out_opt IUIAutomationElement **element);

// func WaitForToplevelWindowWithName(name string) (IAutomationElement, error) {
// }
