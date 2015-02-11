// +build 386
package w32uiautomation

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/mattn/go-ole"
)

func createPropertyCondition(aut *IUIAutomation, propertyId PROPERTYID, value ole.VARIANT) (newCondition *IUIAutomationCondition, err error) {
	fmt.Printf("createPropertyCondition value=%v, value.Val=%x\n", value, value.Val)
	fmt.Printf("createPropertyCondition propertyId=%d\n", propertyId)
	v := VariantToUintptrArray(value)
	fmt.Printf("v[0]=%x, v[1]=%x, v[2]=%x, v[3]=%x, len(v)=%d\n", v[0], v[1], v[2], v[3], len(v))
	hr, r2, e := syscall.Syscall9(
		aut.VTable().CreatePropertyCondition,
		7,
		uintptr(unsafe.Pointer(aut)),
		uintptr(propertyId),
		v[0],
		v[1],
		v[2],
		v[3],
		uintptr(unsafe.Pointer(&newCondition)),
		0,
		0)
	fmt.Printf("createPropertyCondition. hr=%x, r2=%x, e=%v\n", hr, r2, e)
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
