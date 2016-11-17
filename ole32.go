// +build windows

package win

import (
	"unsafe"
)

var (
	// Library
	libole32 uintptr

	// Functions
	coGetInterceptor                      uintptr
	coGetInterceptorFromTypeInfo          uintptr
	bindMoniker                           uintptr
	cLSIDFromProgIDEx                     uintptr
	cLSIDFromString                       uintptr
	coAddRefServerProcess                 uintptr
	coAllowSetForegroundWindow            uintptr
	coBuildVersion                        uintptr
	coCopyProxy                           uintptr
	coCreateFreeThreadedMarshaler         uintptr
	coCreateGuid                          uintptr
	coDisconnectObject                    uintptr
	coFileTimeNow                         uintptr
	coFreeAllLibraries                    uintptr
	coFreeLibrary                         uintptr
	coGetActivationState                  uintptr
	coGetCallContext                      uintptr
	coGetCallState                        uintptr
	coGetCallerTID                        uintptr
	coGetContextToken                     uintptr
	coGetCurrentLogicalThreadId           uintptr
	coGetCurrentProcess                   uintptr
	coGetInterfaceAndReleaseStream        uintptr
	coGetMarshalSizeMax                   uintptr
	coGetObjectContext                    uintptr
	coGetPSClsid                          uintptr
	coGetState                            uintptr
	coGetTreatAsClass                     uintptr
	coImpersonateClient                   uintptr
	coInitialize                          uintptr
	coInitializeWOW                       uintptr
	coIsHandlerConnected                  uintptr
	coIsOle1Class                         uintptr
	coLoadLibrary                         uintptr
	coLockObjectExternal                  uintptr
	coMarshalHresult                      uintptr
	coMarshalInterThreadInterfaceInStream uintptr
	coMarshalInterface                    uintptr
	coQueryProxyBlanket                   uintptr
	coRegisterClassObject                 uintptr
	coRegisterPSClsid                     uintptr
	coReleaseMarshalData                  uintptr
	coReleaseServerProcess                uintptr
	coResumeClassObjects                  uintptr
	coRevertToSelf                        uintptr
	coRevokeInitializeSpy                 uintptr
	coRevokeMallocSpy                     uintptr
	coSetProxyBlanket                     uintptr
	coSetState                            uintptr
	coSuspendClassObjects                 uintptr
	coSwitchCallContext                   uintptr
	coTaskMemAlloc                        uintptr
	coTaskMemFree                         uintptr
	coTaskMemRealloc                      uintptr
	coTreatAsClass                        uintptr
	coUnmarshalHresult                    uintptr
	coUnmarshalInterface                  uintptr
	coWaitForMultipleHandles              uintptr
	createAntiMoniker                     uintptr
	createClassMoniker                    uintptr
	createDataCache                       uintptr
	createFileMoniker                     uintptr
	createGenericComposite                uintptr
	createItemMoniker                     uintptr
	createPointerMoniker                  uintptr
	createStreamOnHGlobal                 uintptr
	dllDebugObjectRPCHook                 uintptr
	getClassFile                          uintptr
	getHGlobalFromStream                  uintptr
	iIDFromString                         uintptr
	isAccelerator                         uintptr
	isEqualGUID                           uintptr
	isValidInterface                      uintptr
	monikerCommonPrefixWith               uintptr
	oleBuildVersion                       uintptr
	oleCreateDefaultHandler               uintptr
	oleFlushClipboard                     uintptr
	oleGetAutoConvert                     uintptr
	oleGetIconOfClass                     uintptr
	oleInitializeWOW                      uintptr
	oleLockRunning                        uintptr
	oleMetafilePictFromIconAndLabel       uintptr
	oleNoteObjectVisible                  uintptr
	oleRegGetMiscStatus                   uintptr
	oleRegGetUserType                     uintptr
	oleSetAutoConvert                     uintptr
	oleSetContainedObject                 uintptr
	propSysAllocString                    uintptr
	propSysFreeString                     uintptr
	readClassStm                          uintptr
	revokeDragDrop                        uintptr
	stgIsStorageFile                      uintptr
	stgSetTimes                           uintptr
	stringFromCLSID                       uintptr
	stringFromGUID2                       uintptr
	wdtpInterfacePointer_UserFree         uintptr
	writeClassStm                         uintptr
)

func init() {
	// Library
	libole32 = doLoadLibrary("ole32.dll")

	// Functions
	coGetInterceptor = doGetProcAddress(libole32, "CoGetInterceptor")
	coGetInterceptorFromTypeInfo = doGetProcAddress(libole32, "CoGetInterceptorFromTypeInfo")
	bindMoniker = doGetProcAddress(libole32, "BindMoniker")
	cLSIDFromProgIDEx = doGetProcAddress(libole32, "CLSIDFromProgIDEx")
	cLSIDFromString = doGetProcAddress(libole32, "CLSIDFromString")
	coAddRefServerProcess = doGetProcAddress(libole32, "CoAddRefServerProcess")
	coAllowSetForegroundWindow = doGetProcAddress(libole32, "CoAllowSetForegroundWindow")
	coBuildVersion = doGetProcAddress(libole32, "CoBuildVersion")
	coCopyProxy = doGetProcAddress(libole32, "CoCopyProxy")
	coCreateFreeThreadedMarshaler = doGetProcAddress(libole32, "CoCreateFreeThreadedMarshaler")
	coCreateGuid = doGetProcAddress(libole32, "CoCreateGuid")
	coDisconnectObject = doGetProcAddress(libole32, "CoDisconnectObject")
	coFileTimeNow = doGetProcAddress(libole32, "CoFileTimeNow")
	coFreeAllLibraries = doGetProcAddress(libole32, "CoFreeAllLibraries")
	coFreeLibrary = doGetProcAddress(libole32, "CoFreeLibrary")
	coGetActivationState = doGetProcAddress(libole32, "CoGetActivationState")
	coGetCallContext = doGetProcAddress(libole32, "CoGetCallContext")
	coGetCallState = doGetProcAddress(libole32, "CoGetCallState")
	coGetCallerTID = doGetProcAddress(libole32, "CoGetCallerTID")
	coGetContextToken = doGetProcAddress(libole32, "CoGetContextToken")
	coGetCurrentLogicalThreadId = doGetProcAddress(libole32, "CoGetCurrentLogicalThreadId")
	coGetCurrentProcess = doGetProcAddress(libole32, "CoGetCurrentProcess")
	coGetInterfaceAndReleaseStream = doGetProcAddress(libole32, "CoGetInterfaceAndReleaseStream")
	coGetMarshalSizeMax = doGetProcAddress(libole32, "CoGetMarshalSizeMax")
	coGetObjectContext = doGetProcAddress(libole32, "CoGetObjectContext")
	coGetPSClsid = doGetProcAddress(libole32, "CoGetPSClsid")
	coGetState = doGetProcAddress(libole32, "CoGetState")
	coGetTreatAsClass = doGetProcAddress(libole32, "CoGetTreatAsClass")
	coImpersonateClient = doGetProcAddress(libole32, "CoImpersonateClient")
	coInitialize = doGetProcAddress(libole32, "CoInitialize")
	coInitializeWOW = doGetProcAddress(libole32, "CoInitializeWOW")
	coIsHandlerConnected = doGetProcAddress(libole32, "CoIsHandlerConnected")
	coIsOle1Class = doGetProcAddress(libole32, "CoIsOle1Class")
	coLoadLibrary = doGetProcAddress(libole32, "CoLoadLibrary")
	coLockObjectExternal = doGetProcAddress(libole32, "CoLockObjectExternal")
	coMarshalHresult = doGetProcAddress(libole32, "CoMarshalHresult")
	coMarshalInterThreadInterfaceInStream = doGetProcAddress(libole32, "CoMarshalInterThreadInterfaceInStream")
	coMarshalInterface = doGetProcAddress(libole32, "CoMarshalInterface")
	coQueryProxyBlanket = doGetProcAddress(libole32, "CoQueryProxyBlanket")
	coRegisterClassObject = doGetProcAddress(libole32, "CoRegisterClassObject")
	coRegisterPSClsid = doGetProcAddress(libole32, "CoRegisterPSClsid")
	coReleaseMarshalData = doGetProcAddress(libole32, "CoReleaseMarshalData")
	coReleaseServerProcess = doGetProcAddress(libole32, "CoReleaseServerProcess")
	coResumeClassObjects = doGetProcAddress(libole32, "CoResumeClassObjects")
	coRevertToSelf = doGetProcAddress(libole32, "CoRevertToSelf")
	coRevokeInitializeSpy = doGetProcAddress(libole32, "CoRevokeInitializeSpy")
	coRevokeMallocSpy = doGetProcAddress(libole32, "CoRevokeMallocSpy")
	coSetProxyBlanket = doGetProcAddress(libole32, "CoSetProxyBlanket")
	coSetState = doGetProcAddress(libole32, "CoSetState")
	coSuspendClassObjects = doGetProcAddress(libole32, "CoSuspendClassObjects")
	coSwitchCallContext = doGetProcAddress(libole32, "CoSwitchCallContext")
	coTaskMemAlloc = doGetProcAddress(libole32, "CoTaskMemAlloc")
	coTaskMemFree = doGetProcAddress(libole32, "CoTaskMemFree")
	coTaskMemRealloc = doGetProcAddress(libole32, "CoTaskMemRealloc")
	coTreatAsClass = doGetProcAddress(libole32, "CoTreatAsClass")
	coUnmarshalHresult = doGetProcAddress(libole32, "CoUnmarshalHresult")
	coUnmarshalInterface = doGetProcAddress(libole32, "CoUnmarshalInterface")
	coWaitForMultipleHandles = doGetProcAddress(libole32, "CoWaitForMultipleHandles")
	createAntiMoniker = doGetProcAddress(libole32, "CreateAntiMoniker")
	createClassMoniker = doGetProcAddress(libole32, "CreateClassMoniker")
	createDataCache = doGetProcAddress(libole32, "CreateDataCache")
	createFileMoniker = doGetProcAddress(libole32, "CreateFileMoniker")
	createGenericComposite = doGetProcAddress(libole32, "CreateGenericComposite")
	createItemMoniker = doGetProcAddress(libole32, "CreateItemMoniker")
	createPointerMoniker = doGetProcAddress(libole32, "CreatePointerMoniker")
	createStreamOnHGlobal = doGetProcAddress(libole32, "CreateStreamOnHGlobal")
	dllDebugObjectRPCHook = doGetProcAddress(libole32, "DllDebugObjectRPCHook")
	getClassFile = doGetProcAddress(libole32, "GetClassFile")
	getHGlobalFromStream = doGetProcAddress(libole32, "GetHGlobalFromStream")
	iIDFromString = doGetProcAddress(libole32, "IIDFromString")
	isAccelerator = doGetProcAddress(libole32, "IsAccelerator")
	isEqualGUID = doGetProcAddress(libole32, "IsEqualGUID")
	isValidInterface = doGetProcAddress(libole32, "IsValidInterface")
	monikerCommonPrefixWith = doGetProcAddress(libole32, "MonikerCommonPrefixWith")
	oleBuildVersion = doGetProcAddress(libole32, "OleBuildVersion")
	oleCreateDefaultHandler = doGetProcAddress(libole32, "OleCreateDefaultHandler")
	oleFlushClipboard = doGetProcAddress(libole32, "OleFlushClipboard")
	oleGetAutoConvert = doGetProcAddress(libole32, "OleGetAutoConvert")
	oleGetIconOfClass = doGetProcAddress(libole32, "OleGetIconOfClass")
	oleInitializeWOW = doGetProcAddress(libole32, "OleInitializeWOW")
	oleLockRunning = doGetProcAddress(libole32, "OleLockRunning")
	oleMetafilePictFromIconAndLabel = doGetProcAddress(libole32, "OleMetafilePictFromIconAndLabel")
	oleNoteObjectVisible = doGetProcAddress(libole32, "OleNoteObjectVisible")
	oleRegGetMiscStatus = doGetProcAddress(libole32, "OleRegGetMiscStatus")
	oleRegGetUserType = doGetProcAddress(libole32, "OleRegGetUserType")
	oleSetAutoConvert = doGetProcAddress(libole32, "OleSetAutoConvert")
	oleSetContainedObject = doGetProcAddress(libole32, "OleSetContainedObject")
	propSysAllocString = doGetProcAddress(libole32, "PropSysAllocString")
	propSysFreeString = doGetProcAddress(libole32, "PropSysFreeString")
	readClassStm = doGetProcAddress(libole32, "ReadClassStm")
	revokeDragDrop = doGetProcAddress(libole32, "RevokeDragDrop")
	stgIsStorageFile = doGetProcAddress(libole32, "StgIsStorageFile")
	stgSetTimes = doGetProcAddress(libole32, "StgSetTimes")
	stringFromCLSID = doGetProcAddress(libole32, "StringFromCLSID")
	stringFromGUID2 = doGetProcAddress(libole32, "StringFromGUID2")
	wdtpInterfacePointer_UserFree = doGetProcAddress(libole32, "WdtpInterfacePointer_UserFree")
	writeClassStm = doGetProcAddress(libole32, "WriteClassStm")
}

func CoGetInterceptor(iidIntercepted REFIID, punkOuter *IUnknown, iid REFIID, ppv uintptr) HRESULT {
	ret1 := syscall6(coGetInterceptor, 4,
		uintptr(unsafe.Pointer(iidIntercepted)),
		uintptr(unsafe.Pointer(punkOuter)),
		uintptr(unsafe.Pointer(iid)),
		ppv,
		0,
		0)
	return HRESULT(ret1)
}

func CoGetInterceptorFromTypeInfo(iidIntercepted REFIID, punkOuter *IUnknown, typeInfo *ITypeInfo, iid REFIID, ppv uintptr) HRESULT {
	ret1 := syscall6(coGetInterceptorFromTypeInfo, 5,
		uintptr(unsafe.Pointer(iidIntercepted)),
		uintptr(unsafe.Pointer(punkOuter)),
		uintptr(unsafe.Pointer(typeInfo)),
		uintptr(unsafe.Pointer(iid)),
		ppv,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): LHCLIENTDOC, LPOLECLIENT, LPOLEOBJECT *, OLECLIPFORMAT, OLEOPT_RENDER, OLESTATUS
// func OleCreate(unnamed0 /*const*/ LPCSTR, unnamed1 LPOLECLIENT, unnamed2 /*const*/ LPCSTR, unnamed3 LHCLIENTDOC, unnamed4 /*const*/ LPCSTR, unnamed5 LPOLEOBJECT *, unnamed6 OLEOPT_RENDER, unnamed7 OLECLIPFORMAT) OLESTATUS

// TODO: Unknown type(s): LHCLIENTDOC, LPOLECLIENT, LPOLEOBJECT *, OLECLIPFORMAT, OLEOPT_RENDER, OLESTATUS
// func OleCreateFromFile(unnamed0 /*const*/ LPCSTR, unnamed1 LPOLECLIENT, unnamed2 /*const*/ LPCSTR, unnamed3 /*const*/ LPCSTR, unnamed4 LHCLIENTDOC, unnamed5 /*const*/ LPCSTR, unnamed6 LPOLEOBJECT *, unnamed7 OLEOPT_RENDER, unnamed8 OLECLIPFORMAT) OLESTATUS

// TODO: Unknown type(s): LPOLEOBJECT, OLESTATUS
// func OleDraw(unnamed0 LPOLEOBJECT, unnamed1 HDC, unnamed2 /*const*/ *RECT, unnamed3 /*const*/ *RECT, unnamed4 HDC) OLESTATUS

// TODO: Unknown type(s): LHCLIENTDOC, LPOLECLIENT, LPOLEOBJECT *, LPOLESTREAM, OLESTATUS
// func OleLoadFromStream(unnamed0 LPOLESTREAM, unnamed1 /*const*/ LPCSTR, unnamed2 LPOLECLIENT, unnamed3 LHCLIENTDOC, unnamed4 /*const*/ LPCSTR, unnamed5 LPOLEOBJECT *) OLESTATUS

// TODO: Unknown type(s): LPOLEOBJECT, LPOLESTREAM, OLESTATUS
// func OleSaveToStream(unnamed0 LPOLEOBJECT, unnamed1 LPOLESTREAM) OLESTATUS

// TODO: Unknown type(s): PROPVARIANT *, PROPVAR_CHANGE_FLAGS, REFPROPVARIANT
// func PropVariantChangeType(ppropvarDest PROPVARIANT *, propvarSrc REFPROPVARIANT, flags PROPVAR_CHANGE_FLAGS, vt VARTYPE) HRESULT

func BindMoniker(pmk LPMONIKER, grfOpt uint32, riid REFIID, ppvResult *LPVOID) HRESULT {
	ret1 := syscall6(bindMoniker, 4,
		uintptr(unsafe.Pointer(pmk)),
		uintptr(grfOpt),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvResult)),
		0,
		0)
	return HRESULT(ret1)
}

func CLSIDFromProgIDEx(progid /*const*/ LPCOLESTR, clsid *CLSID) HRESULT {
	ret1 := syscall3(cLSIDFromProgIDEx, 2,
		uintptr(unsafe.Pointer(progid)),
		uintptr(unsafe.Pointer(clsid)),
		0)
	return HRESULT(ret1)
}

func CLSIDFromString(idstr /*const*/ LPCOLESTR, id *CLSID) HRESULT {
	ret1 := syscall3(cLSIDFromString, 2,
		uintptr(unsafe.Pointer(idstr)),
		uintptr(unsafe.Pointer(id)),
		0)
	return HRESULT(ret1)
}

func CoAddRefServerProcess() ULONG {
	ret1 := syscall3(coAddRefServerProcess, 0,
		0,
		0,
		0)
	return ULONG(ret1)
}

func CoAllowSetForegroundWindow(pUnk *IUnknown, pvReserved uintptr) HRESULT {
	ret1 := syscall3(coAllowSetForegroundWindow, 2,
		uintptr(unsafe.Pointer(pUnk)),
		pvReserved,
		0)
	return HRESULT(ret1)
}

func CoBuildVersion() uint32 {
	ret1 := syscall3(coBuildVersion, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func CoCopyProxy(pProxy *IUnknown, ppCopy **IUnknown) HRESULT {
	ret1 := syscall3(coCopyProxy, 2,
		uintptr(unsafe.Pointer(pProxy)),
		uintptr(unsafe.Pointer(ppCopy)),
		0)
	return HRESULT(ret1)
}

func CoCreateFreeThreadedMarshaler(punkOuter LPUNKNOWN, ppunkMarshal *LPUNKNOWN) HRESULT {
	ret1 := syscall3(coCreateFreeThreadedMarshaler, 2,
		uintptr(unsafe.Pointer(punkOuter)),
		uintptr(unsafe.Pointer(ppunkMarshal)),
		0)
	return HRESULT(ret1)
}

func CoCreateGuid(pguid *GUID) HRESULT {
	ret1 := syscall3(coCreateGuid, 1,
		uintptr(unsafe.Pointer(pguid)),
		0,
		0)
	return HRESULT(ret1)
}

func CoDisconnectObject(lpUnk LPUNKNOWN, reserved uint32) HRESULT {
	ret1 := syscall3(coDisconnectObject, 2,
		uintptr(unsafe.Pointer(lpUnk)),
		uintptr(reserved),
		0)
	return HRESULT(ret1)
}

func CoFileTimeNow(lpFileTime *FILETIME) HRESULT {
	ret1 := syscall3(coFileTimeNow, 1,
		uintptr(unsafe.Pointer(lpFileTime)),
		0,
		0)
	return HRESULT(ret1)
}

func CoFreeAllLibraries() {
	syscall3(coFreeAllLibraries, 0,
		0,
		0,
		0)
}

func CoFreeLibrary(hLibrary HINSTANCE) {
	syscall3(coFreeLibrary, 1,
		uintptr(hLibrary),
		0,
		0)
}

func CoGetActivationState(guid GUID, unknown uint32, unknown2 *uint32) HRESULT {
	ret1 := syscall6(coGetActivationState, 6,
		uintptr(guid.Data1),
		uintptr((uint32(guid.Data2)<<16)|uint32(guid.Data3)),
		uintptr((uint32(guid.Data4[0])<<24)|(uint32(guid.Data4[1]<<16))|(uint32(guid.Data4[2]<<8))|uint32(guid.Data4[3])),
		uintptr((uint32(guid.Data4[4])<<24)|(uint32(guid.Data4[5]<<16))|(uint32(guid.Data4[6]<<8))|uint32(guid.Data4[7])),
		uintptr(unknown),
		uintptr(unsafe.Pointer(unknown2)))
	return HRESULT(ret1)
}

// TODO: Unknown type(s): APTTYPE *, APTTYPEQUALIFIER *
// func CoGetApartmentType(aType APTTYPE *, qualifier APTTYPEQUALIFIER *) HRESULT

func CoGetCallContext(riid REFIID, ppv uintptr) HRESULT {
	ret1 := syscall3(coGetCallContext, 2,
		uintptr(unsafe.Pointer(riid)),
		ppv,
		0)
	return HRESULT(ret1)
}

func CoGetCallState(unknown int32, unknown2 *uint32) HRESULT {
	ret1 := syscall3(coGetCallState, 2,
		uintptr(unknown),
		uintptr(unsafe.Pointer(unknown2)),
		0)
	return HRESULT(ret1)
}

func CoGetCallerTID(lpdwTID *uint32) HRESULT {
	ret1 := syscall3(coGetCallerTID, 1,
		uintptr(unsafe.Pointer(lpdwTID)),
		0,
		0)
	return HRESULT(ret1)
}

func CoGetContextToken(token *ULONG_PTR) HRESULT {
	ret1 := syscall3(coGetContextToken, 1,
		uintptr(unsafe.Pointer(token)),
		0,
		0)
	return HRESULT(ret1)
}

func CoGetCurrentLogicalThreadId(id *GUID) HRESULT {
	ret1 := syscall3(coGetCurrentLogicalThreadId, 1,
		uintptr(unsafe.Pointer(id)),
		0,
		0)
	return HRESULT(ret1)
}

func CoGetCurrentProcess() uint32 {
	ret1 := syscall3(coGetCurrentProcess, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): APTTYPE
// func CoGetDefaultContext(aType APTTYPE, riid REFIID, ppv *LPVOID) HRESULT

// TODO: Unknown type(s): COSERVERINFO *, IStorage *, MULTI_QI *
// func CoGetInstanceFromIStorage(server_info COSERVERINFO *, rclsid *CLSID, outer *IUnknown, cls_context uint32, storage IStorage *, count uint32, results MULTI_QI *) HRESULT

func CoGetInterfaceAndReleaseStream(pStm LPSTREAM, riid REFIID, ppv *LPVOID) HRESULT {
	ret1 := syscall3(coGetInterfaceAndReleaseStream, 3,
		uintptr(unsafe.Pointer(pStm)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IMalloc * *
// func CoGetMalloc(context uint32, imalloc IMalloc * *) HRESULT

func CoGetMarshalSizeMax(pulSize *ULONG, riid REFIID, pUnk *IUnknown, dwDestContext uint32, pvDestContext uintptr, mshlFlags uint32) HRESULT {
	ret1 := syscall6(coGetMarshalSizeMax, 6,
		uintptr(unsafe.Pointer(pulSize)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(pUnk)),
		uintptr(dwDestContext),
		pvDestContext,
		uintptr(mshlFlags))
	return HRESULT(ret1)
}

// TODO: Unknown type(s): BIND_OPTS *
// func CoGetObject(pszName string, pBindOptions BIND_OPTS *, riid REFIID, ppv uintptr) HRESULT

func CoGetObjectContext(riid REFIID, ppv uintptr) HRESULT {
	ret1 := syscall3(coGetObjectContext, 2,
		uintptr(unsafe.Pointer(riid)),
		ppv,
		0)
	return HRESULT(ret1)
}

func CoGetPSClsid(riid REFIID, pclsid *CLSID) HRESULT {
	ret1 := syscall3(coGetPSClsid, 2,
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(pclsid)),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): LPMARSHAL *
// func CoGetStandardMarshal(riid REFIID, pUnk *IUnknown, dwDestContext uint32, pvDestContext LPVOID, mshlflags uint32, ppMarshal LPMARSHAL *) HRESULT

func CoGetState(ppv **IUnknown) HRESULT {
	ret1 := syscall3(coGetState, 1,
		uintptr(unsafe.Pointer(ppv)),
		0,
		0)
	return HRESULT(ret1)
}

func CoGetTreatAsClass(clsidOld /*const*/ REFCLSID, clsidNew *CLSID) HRESULT {
	ret1 := syscall3(coGetTreatAsClass, 2,
		uintptr(unsafe.Pointer(clsidOld)),
		uintptr(unsafe.Pointer(clsidNew)),
		0)
	return HRESULT(ret1)
}

func CoImpersonateClient() HRESULT {
	ret1 := syscall3(coImpersonateClient, 0,
		0,
		0,
		0)
	return HRESULT(ret1)
}

func CoInitialize(lpReserved LPVOID) HRESULT {
	ret1 := syscall3(coInitialize, 1,
		uintptr(unsafe.Pointer(lpReserved)),
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): SOLE_AUTHENTICATION_SERVICE *
// func CoInitializeSecurity(pSecDesc PSECURITY_DESCRIPTOR, cAuthSvc LONG, asAuthSvc SOLE_AUTHENTICATION_SERVICE *, pReserved1 uintptr, dwAuthnLevel uint32, dwImpLevel uint32, pReserved2 uintptr, dwCapabilities uint32, pReserved3 uintptr) HRESULT

func CoInitializeWOW(x uint32, y uint32) HRESULT {
	ret1 := syscall3(coInitializeWOW, 2,
		uintptr(x),
		uintptr(y),
		0)
	return HRESULT(ret1)
}

func CoIsHandlerConnected(pUnk *IUnknown) bool {
	ret1 := syscall3(coIsHandlerConnected, 1,
		uintptr(unsafe.Pointer(pUnk)),
		0,
		0)
	return ret1 != 0
}

func CoIsOle1Class(clsid /*const*/ REFCLSID) bool {
	ret1 := syscall3(coIsOle1Class, 1,
		uintptr(unsafe.Pointer(clsid)),
		0,
		0)
	return ret1 != 0
}

func CoLoadLibrary(lpszLibName LPOLESTR, bAutoFree bool) HINSTANCE {
	ret1 := syscall3(coLoadLibrary, 2,
		uintptr(unsafe.Pointer(lpszLibName)),
		getUintptrFromBool(bAutoFree),
		0)
	return HINSTANCE(ret1)
}

func CoLockObjectExternal(pUnk LPUNKNOWN, fLock bool, fLastUnlockReleases bool) HRESULT {
	ret1 := syscall3(coLockObjectExternal, 3,
		uintptr(unsafe.Pointer(pUnk)),
		getUintptrFromBool(fLock),
		getUintptrFromBool(fLastUnlockReleases))
	return HRESULT(ret1)
}

func CoMarshalHresult(pStm LPSTREAM, hresult HRESULT) HRESULT {
	ret1 := syscall3(coMarshalHresult, 2,
		uintptr(unsafe.Pointer(pStm)),
		uintptr(hresult),
		0)
	return HRESULT(ret1)
}

func CoMarshalInterThreadInterfaceInStream(riid REFIID, pUnk LPUNKNOWN, ppStm *LPSTREAM) HRESULT {
	ret1 := syscall3(coMarshalInterThreadInterfaceInStream, 3,
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(pUnk)),
		uintptr(unsafe.Pointer(ppStm)))
	return HRESULT(ret1)
}

func CoMarshalInterface(pStream *IStream, riid REFIID, pUnk *IUnknown, dwDestContext uint32, pvDestContext uintptr, mshlFlags uint32) HRESULT {
	ret1 := syscall6(coMarshalInterface, 6,
		uintptr(unsafe.Pointer(pStream)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(pUnk)),
		uintptr(dwDestContext),
		pvDestContext,
		uintptr(mshlFlags))
	return HRESULT(ret1)
}

// TODO: Unknown type(s): RPC_AUTHZ_HANDLE *
// func CoQueryClientBlanket(pAuthnSvc *uint32, pAuthzSvc *uint32, pServerPrincName **OLECHAR, pAuthnLevel *uint32, pImpLevel *uint32, pPrivs RPC_AUTHZ_HANDLE *, pCapabilities *uint32) HRESULT

func CoQueryProxyBlanket(pProxy *IUnknown, pAuthnSvc *uint32, pAuthzSvc *uint32, ppServerPrincName **OLECHAR, pAuthnLevel *uint32, pImpLevel *uint32, ppAuthInfo uintptr, pCapabilities *uint32) HRESULT {
	ret1 := syscall9(coQueryProxyBlanket, 8,
		uintptr(unsafe.Pointer(pProxy)),
		uintptr(unsafe.Pointer(pAuthnSvc)),
		uintptr(unsafe.Pointer(pAuthzSvc)),
		uintptr(unsafe.Pointer(ppServerPrincName)),
		uintptr(unsafe.Pointer(pAuthnLevel)),
		uintptr(unsafe.Pointer(pImpLevel)),
		ppAuthInfo,
		uintptr(unsafe.Pointer(pCapabilities)),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IChannelHook *
// func CoRegisterChannelHook(guidExtension REFGUID, pChannelHook IChannelHook *) HRESULT

func CoRegisterClassObject(rclsid /*const*/ REFCLSID, pUnk LPUNKNOWN, dwClsContext uint32, flags uint32, lpdwRegister *uint32) HRESULT {
	ret1 := syscall6(coRegisterClassObject, 5,
		uintptr(unsafe.Pointer(rclsid)),
		uintptr(unsafe.Pointer(pUnk)),
		uintptr(dwClsContext),
		uintptr(flags),
		uintptr(unsafe.Pointer(lpdwRegister)),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IInitializeSpy *
// func CoRegisterInitializeSpy(spy IInitializeSpy *, cookie *ULARGE_INTEGER) HRESULT

// TODO: Unknown type(s): LPMALLOCSPY
// func CoRegisterMallocSpy(pMallocSpy LPMALLOCSPY) HRESULT

// TODO: Unknown type(s): LPMESSAGEFILTER, LPMESSAGEFILTER *
// func CoRegisterMessageFilter(lpMessageFilter LPMESSAGEFILTER, lplpMessageFilter LPMESSAGEFILTER *) HRESULT

func CoRegisterPSClsid(riid REFIID, rclsid /*const*/ REFCLSID) HRESULT {
	ret1 := syscall3(coRegisterPSClsid, 2,
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(rclsid)),
		0)
	return HRESULT(ret1)
}

func CoReleaseMarshalData(pStream *IStream) HRESULT {
	ret1 := syscall3(coReleaseMarshalData, 1,
		uintptr(unsafe.Pointer(pStream)),
		0,
		0)
	return HRESULT(ret1)
}

func CoReleaseServerProcess() ULONG {
	ret1 := syscall3(coReleaseServerProcess, 0,
		0,
		0,
		0)
	return ULONG(ret1)
}

func CoResumeClassObjects() HRESULT {
	ret1 := syscall3(coResumeClassObjects, 0,
		0,
		0,
		0)
	return HRESULT(ret1)
}

func CoRevertToSelf() HRESULT {
	ret1 := syscall3(coRevertToSelf, 0,
		0,
		0,
		0)
	return HRESULT(ret1)
}

func CoRevokeInitializeSpy(cookie ULARGE_INTEGER) HRESULT {
	ret1 := syscall3(coRevokeInitializeSpy, 2,
		uintptr(*(*uint32)(unsafe.Pointer(&cookie))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cookie)) + uintptr(4)))),
		0)
	return HRESULT(ret1)
}

func CoRevokeMallocSpy() HRESULT {
	ret1 := syscall3(coRevokeMallocSpy, 0,
		0,
		0,
		0)
	return HRESULT(ret1)
}

func CoSetProxyBlanket(pProxy *IUnknown, authnSvc uint32, authzSvc uint32, pServerPrincName *OLECHAR, authnLevel uint32, impLevel uint32, pAuthInfo uintptr, capabilities uint32) HRESULT {
	ret1 := syscall9(coSetProxyBlanket, 8,
		uintptr(unsafe.Pointer(pProxy)),
		uintptr(authnSvc),
		uintptr(authzSvc),
		uintptr(unsafe.Pointer(pServerPrincName)),
		uintptr(authnLevel),
		uintptr(impLevel),
		pAuthInfo,
		uintptr(capabilities),
		0)
	return HRESULT(ret1)
}

func CoSetState(pv *IUnknown) HRESULT {
	ret1 := syscall3(coSetState, 1,
		uintptr(unsafe.Pointer(pv)),
		0,
		0)
	return HRESULT(ret1)
}

func CoSuspendClassObjects() HRESULT {
	ret1 := syscall3(coSuspendClassObjects, 0,
		0,
		0,
		0)
	return HRESULT(ret1)
}

func CoSwitchCallContext(pObject *IUnknown, ppOldObject **IUnknown) HRESULT {
	ret1 := syscall3(coSwitchCallContext, 2,
		uintptr(unsafe.Pointer(pObject)),
		uintptr(unsafe.Pointer(ppOldObject)),
		0)
	return HRESULT(ret1)
}

func CoTaskMemAlloc(size SIZE_T) LPVOID {
	ret1 := syscall3(coTaskMemAlloc, 1,
		uintptr(size),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func CoTaskMemFree(ptr LPVOID) {
	syscall3(coTaskMemFree, 1,
		uintptr(unsafe.Pointer(ptr)),
		0,
		0)
}

func CoTaskMemRealloc(pvOld LPVOID, size SIZE_T) LPVOID {
	ret1 := syscall3(coTaskMemRealloc, 2,
		uintptr(unsafe.Pointer(pvOld)),
		uintptr(size),
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func CoTreatAsClass(clsidOld /*const*/ REFCLSID, clsidNew /*const*/ REFCLSID) HRESULT {
	ret1 := syscall3(coTreatAsClass, 2,
		uintptr(unsafe.Pointer(clsidOld)),
		uintptr(unsafe.Pointer(clsidNew)),
		0)
	return HRESULT(ret1)
}

func CoUnmarshalHresult(pStm LPSTREAM, phresult *HRESULT) HRESULT {
	ret1 := syscall3(coUnmarshalHresult, 2,
		uintptr(unsafe.Pointer(pStm)),
		uintptr(unsafe.Pointer(phresult)),
		0)
	return HRESULT(ret1)
}

func CoUnmarshalInterface(pStream *IStream, riid REFIID, ppv *LPVOID) HRESULT {
	ret1 := syscall3(coUnmarshalInterface, 3,
		uintptr(unsafe.Pointer(pStream)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret1)
}

func CoWaitForMultipleHandles(dwFlags uint32, dwTimeout uint32, cHandles ULONG, pHandles *HANDLE, lpdwindex *uint32) HRESULT {
	ret1 := syscall6(coWaitForMultipleHandles, 5,
		uintptr(dwFlags),
		uintptr(dwTimeout),
		uintptr(cHandles),
		uintptr(unsafe.Pointer(pHandles)),
		uintptr(unsafe.Pointer(lpdwindex)),
		0)
	return HRESULT(ret1)
}

func CreateAntiMoniker(ppmk **IMoniker) HRESULT {
	ret1 := syscall3(createAntiMoniker, 1,
		uintptr(unsafe.Pointer(ppmk)),
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): LPBC *
// func CreateBindCtx(reserved uint32, ppbc LPBC *) HRESULT

func CreateClassMoniker(rclsid /*const*/ REFCLSID, ppmk **IMoniker) HRESULT {
	ret1 := syscall3(createClassMoniker, 2,
		uintptr(unsafe.Pointer(rclsid)),
		uintptr(unsafe.Pointer(ppmk)),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IDataAdviseHolder * *
// func CreateDataAdviseHolder(ppDAHolder IDataAdviseHolder * *) HRESULT

func CreateDataCache(pUnkOuter LPUNKNOWN, rclsid /*const*/ REFCLSID, riid REFIID, ppvObj *LPVOID) HRESULT {
	ret1 := syscall6(createDataCache, 4,
		uintptr(unsafe.Pointer(pUnkOuter)),
		uintptr(unsafe.Pointer(rclsid)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObj)),
		0,
		0)
	return HRESULT(ret1)
}

func CreateFileMoniker(lpszPathName /*const*/ LPCOLESTR, ppmk **IMoniker) HRESULT {
	ret1 := syscall3(createFileMoniker, 2,
		uintptr(unsafe.Pointer(lpszPathName)),
		uintptr(unsafe.Pointer(ppmk)),
		0)
	return HRESULT(ret1)
}

func CreateGenericComposite(pmkFirst *IMoniker, pmkRest *IMoniker, ppmkComposite **IMoniker) HRESULT {
	ret1 := syscall3(createGenericComposite, 3,
		uintptr(unsafe.Pointer(pmkFirst)),
		uintptr(unsafe.Pointer(pmkRest)),
		uintptr(unsafe.Pointer(ppmkComposite)))
	return HRESULT(ret1)
}

// TODO: Unknown type(s): ILockBytes * *
// func CreateILockBytesOnHGlobal(global HGLOBAL, delete_on_release bool, ret ILockBytes * *) HRESULT

func CreateItemMoniker(lpszDelim /*const*/ LPCOLESTR, lpszItem /*const*/ LPCOLESTR, ppmk **IMoniker) HRESULT {
	ret1 := syscall3(createItemMoniker, 3,
		uintptr(unsafe.Pointer(lpszDelim)),
		uintptr(unsafe.Pointer(lpszItem)),
		uintptr(unsafe.Pointer(ppmk)))
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IOleAdviseHolder * *
// func CreateOleAdviseHolder(ppOAHolder IOleAdviseHolder * *) HRESULT

func CreatePointerMoniker(punk LPUNKNOWN, ppmk *LPMONIKER) HRESULT {
	ret1 := syscall3(createPointerMoniker, 2,
		uintptr(unsafe.Pointer(punk)),
		uintptr(unsafe.Pointer(ppmk)),
		0)
	return HRESULT(ret1)
}

func CreateStreamOnHGlobal(hGlobal HGLOBAL, fDeleteOnRelease bool, ppstm *LPSTREAM) HRESULT {
	ret1 := syscall3(createStreamOnHGlobal, 3,
		uintptr(hGlobal),
		getUintptrFromBool(fDeleteOnRelease),
		uintptr(unsafe.Pointer(ppstm)))
	return HRESULT(ret1)
}

func DllDebugObjectRPCHook(b bool, dummy uintptr) bool {
	ret1 := syscall3(dllDebugObjectRPCHook, 2,
		getUintptrFromBool(b),
		dummy,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): IDataObject *, IDropSource *
// func DoDragDrop(pDataObject IDataObject *, pDropSource IDropSource *, dwOKEffect uint32, pdwEffect *uint32) HRESULT

// TODO: Unknown type(s): const FMTID *
// func FmtIdToPropStgName(rfmtid /*const*/ const FMTID *, str LPOLESTR) HRESULT

// TODO: Unknown type(s): PROPVARIANT *
// func FreePropVariantArray(cVariants ULONG, rgvars PROPVARIANT *) HRESULT

func GetClassFile(filePathName /*const*/ LPCOLESTR, pclsid *CLSID) HRESULT {
	ret1 := syscall3(getClassFile, 2,
		uintptr(unsafe.Pointer(filePathName)),
		uintptr(unsafe.Pointer(pclsid)),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IStorage *
// func GetConvertStg(stg IStorage *) HRESULT

// TODO: Unknown type(s): ILockBytes *
// func GetHGlobalFromILockBytes(iface ILockBytes *, phglobal *HGLOBAL) HRESULT

func GetHGlobalFromStream(pstm *IStream, phglobal *HGLOBAL) HRESULT {
	ret1 := syscall3(getHGlobalFromStream, 2,
		uintptr(unsafe.Pointer(pstm)),
		uintptr(unsafe.Pointer(phglobal)),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): LPRUNNINGOBJECTTABLE *
// func GetRunningObjectTable(reserved uint32, pprot LPRUNNINGOBJECTTABLE *) HRESULT

func IIDFromString(s /*const*/ LPCOLESTR, iid *IID) HRESULT {
	ret1 := syscall3(iIDFromString, 2,
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(iid)),
		0)
	return HRESULT(ret1)
}

func IsAccelerator(hAccel HACCEL, cAccelEntries int32, lpMsg *MSG, lpwCmd *WORD) bool {
	ret1 := syscall6(isAccelerator, 4,
		uintptr(hAccel),
		uintptr(cAccelEntries),
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(unsafe.Pointer(lpwCmd)),
		0,
		0)
	return ret1 != 0
}

func IsEqualGUID(rguid1 REFGUID, rguid2 REFGUID) bool {
	ret1 := syscall3(isEqualGUID, 2,
		uintptr(unsafe.Pointer(rguid1)),
		uintptr(unsafe.Pointer(rguid2)),
		0)
	return ret1 != 0
}

func IsValidInterface(punk LPUNKNOWN) bool {
	ret1 := syscall3(isValidInterface, 1,
		uintptr(unsafe.Pointer(punk)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPBC
// func MkParseDisplayName(pbc LPBC, szDisplayName /*const*/ LPCOLESTR, pchEaten *uint32, ppmk *LPMONIKER) HRESULT

func MonikerCommonPrefixWith(pmkThis *IMoniker, pmkOther *IMoniker, ppmkCommon **IMoniker) HRESULT {
	ret1 := syscall3(monikerCommonPrefixWith, 3,
		uintptr(unsafe.Pointer(pmkThis)),
		uintptr(unsafe.Pointer(pmkOther)),
		uintptr(unsafe.Pointer(ppmkCommon)))
	return HRESULT(ret1)
}

func OleBuildVersion() uint32 {
	ret1 := syscall3(oleBuildVersion, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): LPOLESTREAM, LPSTORAGE
// func OleConvertIStorageToOLESTREAM(pstg LPSTORAGE, pOleStream LPOLESTREAM) HRESULT

// TODO: Unknown type(s): LPOLESTREAM, LPSTORAGE, const DVTARGETDEVICE *
// func OleConvertOLESTREAMToIStorage(pOleStream LPOLESTREAM, pstg LPSTORAGE, ptd /*const*/ const DVTARGETDEVICE *) HRESULT

func OleCreateDefaultHandler(clsid /*const*/ REFCLSID, pUnkOuter LPUNKNOWN, riid REFIID, ppvObj *LPVOID) HRESULT {
	ret1 := syscall6(oleCreateDefaultHandler, 4,
		uintptr(unsafe.Pointer(clsid)),
		uintptr(unsafe.Pointer(pUnkOuter)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObj)),
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IClassFactory *
// func OleCreateEmbeddingHelper(clsid /*const*/ REFCLSID, pUnkOuter LPUNKNOWN, flags uint32, pCF IClassFactory *, riid REFIID, ppvObj *LPVOID) HRESULT

// TODO: Unknown type(s): LPDATAOBJECT, LPFORMATETC, LPOLECLIENTSITE, LPSTORAGE
// func OleCreateFromData(data LPDATAOBJECT, iid REFIID, renderopt uint32, fmt LPFORMATETC, client_site LPOLECLIENTSITE, stg LPSTORAGE, obj *LPVOID) HRESULT

// TODO: Unknown type(s): FORMATETC *, IAdviseSink *, IDataObject *, IOleClientSite *, IStorage *
// func OleCreateFromDataEx(data IDataObject *, iid REFIID, flags uint32, renderopt uint32, num_cache_fmts ULONG, adv_flags *uint32, cache_fmts FORMATETC *, sink IAdviseSink *, conns *uint32, client_site IOleClientSite *, stg IStorage *, obj uintptr) HRESULT

// TODO: Unknown type(s): FORMATETC *, IAdviseSink *, IOleClientSite *, IStorage *
// func OleCreateFromFileEx(clsid /*const*/ REFCLSID, filename /*const*/ *OLECHAR, iid REFIID, flags uint32, renderopt uint32, num_fmts ULONG, adv_flags *uint32, fmts FORMATETC *, sink IAdviseSink *, conns *uint32, client_site IOleClientSite *, stg IStorage *, obj uintptr) HRESULT

// TODO: Unknown type(s): LPFORMATETC, LPOLECLIENTSITE, LPSTORAGE
// func OleCreateLink(pmkLinkSrc LPMONIKER, riid REFIID, renderopt uint32, lpFormatEtc LPFORMATETC, pClientSite LPOLECLIENTSITE, pStg LPSTORAGE, ppvObj *LPVOID) HRESULT

// TODO: Unknown type(s): FORMATETC *, IDataObject *, IOleClientSite *, IStorage *
// func OleCreateLinkFromData(data IDataObject *, iid REFIID, renderopt uint32, fmt FORMATETC *, client_site IOleClientSite *, stg IStorage *, obj uintptr) HRESULT

// TODO: Unknown type(s): LPFORMATETC, LPOLECLIENTSITE, LPSTORAGE
// func OleCreateLinkToFile(lpszFileName /*const*/ LPCOLESTR, riid REFIID, renderopt uint32, lpFormatEtc LPFORMATETC, pClientSite LPOLECLIENTSITE, pStg LPSTORAGE, ppvObj *LPVOID) HRESULT

// TODO: Unknown type(s): HOLEMENU, LPOLEMENUGROUPWIDTHS
// func OleCreateMenuDescriptor(hmenuCombined HMENU, lpMenuWidths LPOLEMENUGROUPWIDTHS) HOLEMENU

// TODO: Unknown type(s): FORMATETC *, IDataObject *, IOleClientSite *, IStorage *
// func OleCreateStaticFromData(data IDataObject *, iid REFIID, renderopt uint32, fmt FORMATETC *, client_site IOleClientSite *, stg IStorage *, obj uintptr) HRESULT

// TODO: Unknown type(s): HOLEMENU
// func OleDestroyMenuDescriptor(hmenuDescriptor HOLEMENU) HRESULT

// TODO: Unknown type(s): LPSTORAGE
// func OleDoAutoConvert(pStg LPSTORAGE, pClsidNew *CLSID) HRESULT

// TODO: Unknown type(s): CLIPFORMAT
// func OleDuplicateData(hSrc HANDLE, cfFormat CLIPFORMAT, uiFlags uint32) HANDLE

func OleFlushClipboard() HRESULT {
	ret1 := syscall3(oleFlushClipboard, 0,
		0,
		0,
		0)
	return HRESULT(ret1)
}

func OleGetAutoConvert(clsidOld /*const*/ REFCLSID, pClsidNew *CLSID) HRESULT {
	ret1 := syscall3(oleGetAutoConvert, 2,
		uintptr(unsafe.Pointer(clsidOld)),
		uintptr(unsafe.Pointer(pClsidNew)),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IDataObject * *
// func OleGetClipboard(obj IDataObject * *) HRESULT

func OleGetIconOfClass(rclsid /*const*/ REFCLSID, lpszLabel LPOLESTR, fUseTypeAsLabel bool) HGLOBAL {
	ret1 := syscall3(oleGetIconOfClass, 3,
		uintptr(unsafe.Pointer(rclsid)),
		uintptr(unsafe.Pointer(lpszLabel)),
		getUintptrFromBool(fUseTypeAsLabel))
	return HGLOBAL(ret1)
}

func OleInitializeWOW(x uint32, y uint32) HRESULT {
	ret1 := syscall3(oleInitializeWOW, 2,
		uintptr(x),
		uintptr(y),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IDataObject *
// func OleIsCurrentClipboard(data IDataObject *) HRESULT

// TODO: Unknown type(s): LPOLEOBJECT
// func OleIsRunning(object LPOLEOBJECT) bool

// TODO: Unknown type(s): LPOLECLIENTSITE, LPSTORAGE
// func OleLoad(pStg LPSTORAGE, riid REFIID, pClientSite LPOLECLIENTSITE, ppvObj *LPVOID) HRESULT

func OleLockRunning(pUnknown LPUNKNOWN, fLock bool, fLastUnlockCloses bool) HRESULT {
	ret1 := syscall3(oleLockRunning, 3,
		uintptr(unsafe.Pointer(pUnknown)),
		getUintptrFromBool(fLock),
		getUintptrFromBool(fLastUnlockCloses))
	return HRESULT(ret1)
}

func OleMetafilePictFromIconAndLabel(hIcon HICON, lpszLabel LPOLESTR, lpszSourceFile LPOLESTR, iIconIndex uint32) HGLOBAL {
	ret1 := syscall6(oleMetafilePictFromIconAndLabel, 4,
		uintptr(hIcon),
		uintptr(unsafe.Pointer(lpszLabel)),
		uintptr(unsafe.Pointer(lpszSourceFile)),
		uintptr(iIconIndex),
		0,
		0)
	return HGLOBAL(ret1)
}

func OleNoteObjectVisible(pUnknown LPUNKNOWN, bVisible bool) HRESULT {
	ret1 := syscall3(oleNoteObjectVisible, 2,
		uintptr(unsafe.Pointer(pUnknown)),
		getUintptrFromBool(bVisible),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IDataObject *
// func OleQueryCreateFromData(data IDataObject *) HRESULT

// TODO: Unknown type(s): IDataObject *
// func OleQueryLinkFromData(pSrcDataObject IDataObject *) HRESULT

// TODO: Unknown type(s): LPENUMOLEVERB *
// func OleRegEnumVerbs(clsid /*const*/ REFCLSID, ppenum LPENUMOLEVERB *) HRESULT

func OleRegGetMiscStatus(clsid /*const*/ REFCLSID, dwAspect uint32, pdwStatus *uint32) HRESULT {
	ret1 := syscall3(oleRegGetMiscStatus, 3,
		uintptr(unsafe.Pointer(clsid)),
		uintptr(dwAspect),
		uintptr(unsafe.Pointer(pdwStatus)))
	return HRESULT(ret1)
}

func OleRegGetUserType(clsid /*const*/ REFCLSID, form uint32, usertype *LPOLESTR) HRESULT {
	ret1 := syscall3(oleRegGetUserType, 3,
		uintptr(unsafe.Pointer(clsid)),
		uintptr(form),
		uintptr(unsafe.Pointer(usertype)))
	return HRESULT(ret1)
}

// TODO: Unknown type(s): LPPERSISTSTORAGE, LPSTORAGE
// func OleSave(pPS LPPERSISTSTORAGE, pStg LPSTORAGE, fSameAsLoad bool) HRESULT

func OleSetAutoConvert(clsidOld /*const*/ REFCLSID, clsidNew /*const*/ REFCLSID) HRESULT {
	ret1 := syscall3(oleSetAutoConvert, 2,
		uintptr(unsafe.Pointer(clsidOld)),
		uintptr(unsafe.Pointer(clsidNew)),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IDataObject *
// func OleSetClipboard(data IDataObject *) HRESULT

func OleSetContainedObject(pUnknown LPUNKNOWN, fContained bool) HRESULT {
	ret1 := syscall3(oleSetContainedObject, 2,
		uintptr(unsafe.Pointer(pUnknown)),
		getUintptrFromBool(fContained),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): HOLEMENU, LPOLEINPLACEACTIVEOBJECT, LPOLEINPLACEFRAME
// func OleSetMenuDescriptor(hOleMenu HOLEMENU, hwndFrame HWND, hwndActiveObject HWND, lpFrame LPOLEINPLACEFRAME, lpActiveObject LPOLEINPLACEACTIVEOBJECT) HRESULT

// TODO: Unknown type(s): LPOLEINPLACEFRAME, LPOLEINPLACEFRAMEINFO
// func OleTranslateAccelerator(lpFrame LPOLEINPLACEFRAME, lpFrameInfo LPOLEINPLACEFRAMEINFO, lpmsg *MSG) HRESULT

// TODO: Unknown type(s): FMTID *
// func PropStgNameToFmtId(str /*const*/ LPOLESTR, rfmtid FMTID *) HRESULT

func PropSysAllocString(str /*const*/ LPCOLESTR) BSTR {
	ret1 := syscall3(propSysAllocString, 1,
		uintptr(unsafe.Pointer(str)),
		0,
		0)
	return (BSTR)(unsafe.Pointer(ret1))
}

func PropSysFreeString(str LPOLESTR) {
	syscall3(propSysFreeString, 1,
		uintptr(unsafe.Pointer(str)),
		0,
		0)
}

// TODO: Unknown type(s): PROPVARIANT *
// func PropVariantClear(pvar PROPVARIANT *) HRESULT

// TODO: Unknown type(s): PROPVARIANT *, const PROPVARIANT *
// func PropVariantCopy(pvarDest PROPVARIANT *, pvarSrc /*const*/ const PROPVARIANT *) HRESULT

// TODO: Unknown type(s): IStorage *
// func ReadClassStg(pstg IStorage *, pclsid *CLSID) HRESULT

func ReadClassStm(pStm *IStream, pclsid *CLSID) HRESULT {
	ret1 := syscall3(readClassStm, 2,
		uintptr(unsafe.Pointer(pStm)),
		uintptr(unsafe.Pointer(pclsid)),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): CLIPFORMAT *, LPSTORAGE
// func ReadFmtUserTypeStg(pstg LPSTORAGE, pcf CLIPFORMAT *, lplpszUserType *LPOLESTR) HRESULT

// TODO: Unknown type(s): LPDROPTARGET
// func RegisterDragDrop(hwnd HWND, pDropTarget LPDROPTARGET) HRESULT

// TODO: Unknown type(s): STGMEDIUM *
// func ReleaseStgMedium(pmedium STGMEDIUM *)

func RevokeDragDrop(hwnd HWND) HRESULT {
	ret1 := syscall3(revokeDragDrop, 1,
		uintptr(hwnd),
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IStorage *
// func SetConvertStg(storage IStorage *, convert bool) HRESULT

// TODO: Unknown type(s): PROPVARIANT *, const SERIALIZEDPROPERTYVALUE *
// func StgConvertPropertyToVariant(prop /*const*/ const SERIALIZEDPROPERTYVALUE *, codePage USHORT, pvar PROPVARIANT *, pma uintptr) BOOLEAN

// TODO: Unknown type(s): IStorage * *
// func StgCreateDocfile(pwcsName /*const*/ LPCOLESTR, grfMode uint32, reserved uint32, ppstgOpen IStorage * *) HRESULT

// TODO: Unknown type(s): ILockBytes *, IStorage * *
// func StgCreateDocfileOnILockBytes(plkbyt ILockBytes *, grfMode uint32, reserved uint32, ppstgOpen IStorage * *) HRESULT

// TODO: Unknown type(s): IPropertySetStorage * *, IStorage *
// func StgCreatePropSetStg(pstg IStorage *, reserved uint32, propset IPropertySetStorage * *) HRESULT

// TODO: Unknown type(s): IPropertyStorage * *, REFFMTID
// func StgCreatePropStg(unk *IUnknown, fmt REFFMTID, clsid /*const*/ *CLSID, flags uint32, reserved uint32, prop_stg IPropertyStorage * *) HRESULT

// TODO: Unknown type(s): STGOPTIONS *
// func StgCreateStorageEx(pwcsName /*const*/ *WCHAR, grfMode uint32, stgfmt uint32, grfAttrs uint32, pStgOptions STGOPTIONS *, reserved uintptr, riid REFIID, ppObjectOpen uintptr) HRESULT

func StgIsStorageFile(fn /*const*/ LPCOLESTR) HRESULT {
	ret1 := syscall3(stgIsStorageFile, 1,
		uintptr(unsafe.Pointer(fn)),
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): ILockBytes *
// func StgIsStorageILockBytes(plkbyt ILockBytes *) HRESULT

// TODO: Unknown type(s): IPropertyStorage * *, REFFMTID
// func StgOpenPropStg(unk *IUnknown, fmt REFFMTID, flags uint32, reserved uint32, prop_stg IPropertyStorage * *) HRESULT

// TODO: Unknown type(s): IStorage *, IStorage * *, SNB
// func StgOpenStorage(pwcsName /*const*/ *OLECHAR, pstgPriority IStorage *, grfMode uint32, snbExclude SNB, reserved uint32, ppstgOpen IStorage * *) HRESULT

// TODO: Unknown type(s): STGOPTIONS *
// func StgOpenStorageEx(pwcsName /*const*/ *WCHAR, grfMode uint32, stgfmt uint32, grfAttrs uint32, pStgOptions STGOPTIONS *, reserved uintptr, riid REFIID, ppObjectOpen uintptr) HRESULT

// TODO: Unknown type(s): ILockBytes *, IStorage *, IStorage * *, SNB
// func StgOpenStorageOnILockBytes(plkbyt ILockBytes *, pstgPriority IStorage *, grfMode uint32, snbExclude SNB, reserved uint32, ppstgOpen IStorage * *) HRESULT

func StgSetTimes(str /*const*/ *OLECHAR, pctime /*const*/ *FILETIME, patime /*const*/ *FILETIME, pmtime /*const*/ *FILETIME) HRESULT {
	ret1 := syscall6(stgSetTimes, 4,
		uintptr(unsafe.Pointer(str)),
		uintptr(unsafe.Pointer(pctime)),
		uintptr(unsafe.Pointer(patime)),
		uintptr(unsafe.Pointer(pmtime)),
		0,
		0)
	return HRESULT(ret1)
}

func StringFromCLSID(id /*const*/ REFCLSID, idstr *LPOLESTR) HRESULT {
	ret1 := syscall3(stringFromCLSID, 2,
		uintptr(unsafe.Pointer(id)),
		uintptr(unsafe.Pointer(idstr)),
		0)
	return HRESULT(ret1)
}

func StringFromGUID2(id REFGUID, str LPOLESTR, cmax int32) int32 {
	ret1 := syscall3(stringFromGUID2, 3,
		uintptr(unsafe.Pointer(id)),
		uintptr(unsafe.Pointer(str)),
		uintptr(cmax))
	return int32(ret1)
}

func WdtpInterfacePointer_UserFree(punk *IUnknown) {
	syscall3(wdtpInterfacePointer_UserFree, 1,
		uintptr(unsafe.Pointer(punk)),
		0,
		0)
}

// TODO: Unknown type(s): IStorage *
// func WriteClassStg(pStg IStorage *, rclsid /*const*/ REFCLSID) HRESULT

func WriteClassStm(pStm *IStream, rclsid /*const*/ REFCLSID) HRESULT {
	ret1 := syscall3(writeClassStm, 2,
		uintptr(unsafe.Pointer(pStm)),
		uintptr(unsafe.Pointer(rclsid)),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): CLIPFORMAT, LPSTORAGE
// func WriteFmtUserTypeStg(pstg LPSTORAGE, cf CLIPFORMAT, lpszUserType LPOLESTR) HRESULT
