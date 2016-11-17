// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	libshell32 uintptr

	// Functions
	sHCreatePropSheetExtArray               uintptr
	sHDestroyPropSheetExtArray              uintptr
	callCPLEntry16                          uintptr
	checkEscapes                            uintptr
	control_FillCache_RunDLL                uintptr
	control_RunDLL                          uintptr
	dAD_DragEnterEx                         uintptr
	dAD_DragLeave                           uintptr
	dAD_DragMove                            uintptr
	dAD_SetDragImage                        uintptr
	dAD_ShowDragImage                       uintptr
	doEnvironmentSubst                      uintptr
	dragAcceptFiles                         uintptr
	driveType                               uintptr
	duplicateIcon                           uintptr
	extractAssociatedIconEx                 uintptr
	extractAssociatedIcon                   uintptr
	extractIconEx                           uintptr
	extractIcon                             uintptr
	extractVersionResource16W               uintptr
	findExecutable                          uintptr
	freeIconList                            uintptr
	getCurrentProcessExplicitAppUserModelID uintptr
	iLClone                                 uintptr
	iLCloneFirst                            uintptr
	iLCombine                               uintptr
	iLCreateFromPath                        uintptr
	iLFindChild                             uintptr
	iLFindLastID                            uintptr
	iLFree                                  uintptr
	iLGetNext                               uintptr
	iLGetSize                               uintptr
	iLIsEqual                               uintptr
	iLIsParent                              uintptr
	iLRemoveLastID                          uintptr
	iLSaveToStream                          uintptr
	initNetworkAddressControl               uintptr
	isLFNDrive                              uintptr
	isNetDrive                              uintptr
	isUserAnAdmin                           uintptr
	openAs_RunDLL                           uintptr
	pathCleanupSpec                         uintptr
	pathYetAnotherMakeUniqueName            uintptr
	pickIconDlg                             uintptr
	realDriveType                           uintptr
	regenerateUserEnvironment               uintptr
	restartDialog                           uintptr
	restartDialogEx                         uintptr
	sHAddToRecentDocs                       uintptr
	sHAlloc                                 uintptr
	sHBindToParent                          uintptr
	sHChangeNotification_Unlock             uintptr
	sHChangeNotify                          uintptr
	sHChangeNotifyDeregister                uintptr
	sHCloneSpecialIDList                    uintptr
	sHCoCreateInstance                      uintptr
	sHCreateDirectory                       uintptr
	sHCreateDirectoryEx                     uintptr
	sHCreateFileExtractIconW                uintptr
	sHCreateItemFromParsingName             uintptr
	sHCreateQueryCancelAutoPlayMoniker      uintptr
	sHDefExtractIcon                        uintptr
	sHEmptyRecycleBin                       uintptr
	sHEnumerateUnreadMailAccountsW          uintptr
	sHFindFiles                             uintptr
	sHFind_InitMenuPopup                    uintptr
	sHFlushClipboard                        uintptr
	sHFlushSFCache                          uintptr
	sHFormatDrive                           uintptr
	sHFree                                  uintptr
	sHFreeNameMappings                      uintptr
	sHGetFolderPathAndSubDir                uintptr
	sHGetFolderPath                         uintptr
	sHGetIconOverlayIndex                   uintptr
	sHGetImageList                          uintptr
	sHGetInstanceExplorer                   uintptr
	sHGetItemFromObject                     uintptr
	sHGetLocalizedName                      uintptr
	sHGetNewLinkInfo                        uintptr
	sHGetPathFromIDList                     uintptr
	sHGetPropertyStoreForWindow             uintptr
	sHGetSpecialFolderPath                  uintptr
	sHHandleUpdateImage                     uintptr
	sHHelpShortcuts_RunDLL                  uintptr
	sHIsFileAvailableOffline                uintptr
	sHLoadInProc                            uintptr
	sHLoadNonloadedIconOverlayIdentifiers   uintptr
	sHObjectProperties                      uintptr
	sHPathPrepareForWrite                   uintptr
	sHRemoveLocalizedName                   uintptr
	sHRunControlPanel                       uintptr
	sHSetInstanceExplorer                   uintptr
	sHSetLocalizedName                      uintptr
	sHSetUnreadMailCountW                   uintptr
	sHShellFolderView_Message               uintptr
	sHUpdateImage                           uintptr
	sHUpdateRecycleBinIcon                  uintptr
	sHValidateUNC                           uintptr
	setCurrentProcessExplicitAppUserModelID uintptr
	sheChangeDir                            uintptr
	sheGetDir                               uintptr
	shellAbout                              uintptr
	shellExecute                            uintptr
	shell_GetImageLists                     uintptr
	shell_MergeMenus                        uintptr
	wOWShellExecute                         uintptr
)

func init() {
	// Library
	libshell32 = doLoadLibrary("shell32.dll")

	// Functions
	sHCreatePropSheetExtArray = doGetProcAddress(libshell32, "SHCreatePropSheetExtArray")
	sHDestroyPropSheetExtArray = doGetProcAddress(libshell32, "SHDestroyPropSheetExtArray")
	callCPLEntry16 = doGetProcAddress(libshell32, "CallCPLEntry16")
	checkEscapes = doGetProcAddress(libshell32, "CheckEscapesW")
	control_FillCache_RunDLL = doGetProcAddress(libshell32, "Control_FillCache_RunDLLW")
	control_RunDLL = doGetProcAddress(libshell32, "Control_RunDLLW")
	dAD_DragEnterEx = doGetProcAddress(libshell32, "DAD_DragEnterEx")
	dAD_DragLeave = doGetProcAddress(libshell32, "DAD_DragLeave")
	dAD_DragMove = doGetProcAddress(libshell32, "DAD_DragMove")
	dAD_SetDragImage = doGetProcAddress(libshell32, "DAD_SetDragImage")
	dAD_ShowDragImage = doGetProcAddress(libshell32, "DAD_ShowDragImage")
	doEnvironmentSubst = doGetProcAddress(libshell32, "DoEnvironmentSubstW")
	dragAcceptFiles = doGetProcAddress(libshell32, "DragAcceptFiles")
	driveType = doGetProcAddress(libshell32, "DriveType")
	duplicateIcon = doGetProcAddress(libshell32, "DuplicateIcon")
	extractAssociatedIconEx = doGetProcAddress(libshell32, "ExtractAssociatedIconExW")
	extractAssociatedIcon = doGetProcAddress(libshell32, "ExtractAssociatedIconW")
	extractIconEx = doGetProcAddress(libshell32, "ExtractIconExW")
	extractIcon = doGetProcAddress(libshell32, "ExtractIconW")
	extractVersionResource16W = doGetProcAddress(libshell32, "ExtractVersionResource16W")
	findExecutable = doGetProcAddress(libshell32, "FindExecutableW")
	freeIconList = doGetProcAddress(libshell32, "FreeIconList")
	getCurrentProcessExplicitAppUserModelID = doGetProcAddress(libshell32, "GetCurrentProcessExplicitAppUserModelID")
	iLClone = doGetProcAddress(libshell32, "ILClone")
	iLCloneFirst = doGetProcAddress(libshell32, "ILCloneFirst")
	iLCombine = doGetProcAddress(libshell32, "ILCombine")
	iLCreateFromPath = doGetProcAddress(libshell32, "ILCreateFromPathW")
	iLFindChild = doGetProcAddress(libshell32, "ILFindChild")
	iLFindLastID = doGetProcAddress(libshell32, "ILFindLastID")
	iLFree = doGetProcAddress(libshell32, "ILFree")
	iLGetNext = doGetProcAddress(libshell32, "ILGetNext")
	iLGetSize = doGetProcAddress(libshell32, "ILGetSize")
	iLIsEqual = doGetProcAddress(libshell32, "ILIsEqual")
	iLIsParent = doGetProcAddress(libshell32, "ILIsParent")
	iLRemoveLastID = doGetProcAddress(libshell32, "ILRemoveLastID")
	iLSaveToStream = doGetProcAddress(libshell32, "ILSaveToStream")
	initNetworkAddressControl = doGetProcAddress(libshell32, "InitNetworkAddressControl")
	isLFNDrive = doGetProcAddress(libshell32, "IsLFNDriveW")
	isNetDrive = doGetProcAddress(libshell32, "IsNetDrive")
	isUserAnAdmin = doGetProcAddress(libshell32, "IsUserAnAdmin")
	openAs_RunDLL = doGetProcAddress(libshell32, "OpenAs_RunDLLW")
	pathCleanupSpec = doGetProcAddress(libshell32, "PathCleanupSpec")
	pathYetAnotherMakeUniqueName = doGetProcAddress(libshell32, "PathYetAnotherMakeUniqueName")
	pickIconDlg = doGetProcAddress(libshell32, "PickIconDlg")
	realDriveType = doGetProcAddress(libshell32, "RealDriveType")
	regenerateUserEnvironment = doGetProcAddress(libshell32, "RegenerateUserEnvironment")
	restartDialog = doGetProcAddress(libshell32, "RestartDialog")
	restartDialogEx = doGetProcAddress(libshell32, "RestartDialogEx")
	sHAddToRecentDocs = doGetProcAddress(libshell32, "SHAddToRecentDocs")
	sHAlloc = doGetProcAddress(libshell32, "SHAlloc")
	sHBindToParent = doGetProcAddress(libshell32, "SHBindToParent")
	sHChangeNotification_Unlock = doGetProcAddress(libshell32, "SHChangeNotification_Unlock")
	sHChangeNotify = doGetProcAddress(libshell32, "SHChangeNotify")
	sHChangeNotifyDeregister = doGetProcAddress(libshell32, "SHChangeNotifyDeregister")
	sHCloneSpecialIDList = doGetProcAddress(libshell32, "SHCloneSpecialIDList")
	sHCoCreateInstance = doGetProcAddress(libshell32, "SHCoCreateInstance")
	sHCreateDirectory = doGetProcAddress(libshell32, "SHCreateDirectory")
	sHCreateDirectoryEx = doGetProcAddress(libshell32, "SHCreateDirectoryExW")
	sHCreateFileExtractIconW = doGetProcAddress(libshell32, "SHCreateFileExtractIconW")
	sHCreateItemFromParsingName = doGetProcAddress(libshell32, "SHCreateItemFromParsingName")
	sHCreateQueryCancelAutoPlayMoniker = doGetProcAddress(libshell32, "SHCreateQueryCancelAutoPlayMoniker")
	sHDefExtractIcon = doGetProcAddress(libshell32, "SHDefExtractIconW")
	sHEmptyRecycleBin = doGetProcAddress(libshell32, "SHEmptyRecycleBinW")
	sHEnumerateUnreadMailAccountsW = doGetProcAddress(libshell32, "SHEnumerateUnreadMailAccountsW")
	sHFindFiles = doGetProcAddress(libshell32, "SHFindFiles")
	sHFind_InitMenuPopup = doGetProcAddress(libshell32, "SHFind_InitMenuPopup")
	sHFlushClipboard = doGetProcAddress(libshell32, "SHFlushClipboard")
	sHFlushSFCache = doGetProcAddress(libshell32, "SHFlushSFCache")
	sHFormatDrive = doGetProcAddress(libshell32, "SHFormatDrive")
	sHFree = doGetProcAddress(libshell32, "SHFree")
	sHFreeNameMappings = doGetProcAddress(libshell32, "SHFreeNameMappings")
	sHGetFolderPathAndSubDir = doGetProcAddress(libshell32, "SHGetFolderPathAndSubDirW")
	sHGetFolderPath = doGetProcAddress(libshell32, "SHGetFolderPathW")
	sHGetIconOverlayIndex = doGetProcAddress(libshell32, "SHGetIconOverlayIndexW")
	sHGetImageList = doGetProcAddress(libshell32, "SHGetImageList")
	sHGetInstanceExplorer = doGetProcAddress(libshell32, "SHGetInstanceExplorer")
	sHGetItemFromObject = doGetProcAddress(libshell32, "SHGetItemFromObject")
	sHGetLocalizedName = doGetProcAddress(libshell32, "SHGetLocalizedName")
	sHGetNewLinkInfo = doGetProcAddress(libshell32, "SHGetNewLinkInfoW")
	sHGetPathFromIDList = doGetProcAddress(libshell32, "SHGetPathFromIDListW")
	sHGetPropertyStoreForWindow = doGetProcAddress(libshell32, "SHGetPropertyStoreForWindow")
	sHGetSpecialFolderPath = doGetProcAddress(libshell32, "SHGetSpecialFolderPathW")
	sHHandleUpdateImage = doGetProcAddress(libshell32, "SHHandleUpdateImage")
	sHHelpShortcuts_RunDLL = doGetProcAddress(libshell32, "SHHelpShortcuts_RunDLLW")
	sHIsFileAvailableOffline = doGetProcAddress(libshell32, "SHIsFileAvailableOffline")
	sHLoadInProc = doGetProcAddress(libshell32, "SHLoadInProc")
	sHLoadNonloadedIconOverlayIdentifiers = doGetProcAddress(libshell32, "SHLoadNonloadedIconOverlayIdentifiers")
	sHObjectProperties = doGetProcAddress(libshell32, "SHObjectProperties")
	sHPathPrepareForWrite = doGetProcAddress(libshell32, "SHPathPrepareForWriteW")
	sHRemoveLocalizedName = doGetProcAddress(libshell32, "SHRemoveLocalizedName")
	sHRunControlPanel = doGetProcAddress(libshell32, "SHRunControlPanel")
	sHSetInstanceExplorer = doGetProcAddress(libshell32, "SHSetInstanceExplorer")
	sHSetLocalizedName = doGetProcAddress(libshell32, "SHSetLocalizedName")
	sHSetUnreadMailCountW = doGetProcAddress(libshell32, "SHSetUnreadMailCountW")
	sHShellFolderView_Message = doGetProcAddress(libshell32, "SHShellFolderView_Message")
	sHUpdateImage = doGetProcAddress(libshell32, "SHUpdateImageW")
	sHUpdateRecycleBinIcon = doGetProcAddress(libshell32, "SHUpdateRecycleBinIcon")
	sHValidateUNC = doGetProcAddress(libshell32, "SHValidateUNC")
	setCurrentProcessExplicitAppUserModelID = doGetProcAddress(libshell32, "SetCurrentProcessExplicitAppUserModelID")
	sheChangeDir = doGetProcAddress(libshell32, "SheChangeDirW")
	sheGetDir = doGetProcAddress(libshell32, "SheGetDirW")
	shellAbout = doGetProcAddress(libshell32, "ShellAboutW")
	shellExecute = doGetProcAddress(libshell32, "ShellExecuteW")
	shell_GetImageLists = doGetProcAddress(libshell32, "Shell_GetImageLists")
	shell_MergeMenus = doGetProcAddress(libshell32, "Shell_MergeMenus")
	wOWShellExecute = doGetProcAddress(libshell32, "WOWShellExecute")
}

// TODO: Unknown type(s): LPFNADDPROPSHEETPAGE
// func SHAddFromPropSheetExtArray(hpsxa HPSXA, lpfnAddPage LPFNADDPROPSHEETPAGE, lParam LPARAM) uint32

func SHCreatePropSheetExtArray(hKey HKEY, pszSubKey string, max_iface uint32) HPSXA {
	pszSubKeyStr := unicode16FromString(pszSubKey)
	ret1 := syscall3(sHCreatePropSheetExtArray, 3,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&pszSubKeyStr[0])),
		uintptr(max_iface))
	return HPSXA(ret1)
}

func SHDestroyPropSheetExtArray(hpsxa HPSXA) {
	syscall3(sHDestroyPropSheetExtArray, 1,
		uintptr(hpsxa),
		0,
		0)
}

// TODO: Unknown type(s): LPFNADDPROPSHEETPAGE
// func SHReplaceFromPropSheetExtArray(hpsxa HPSXA, uPageID uint32, lpfnReplaceWith LPFNADDPROPSHEETPAGE, lParam LPARAM) uint32

// TODO: Unknown type(s): IContextMenu * *, IShellFolder *, LPFNDFMCALLBACK
// func CDefFolderMenu_Create2(pidlFolder /*const*/ LPCITEMIDLIST, hwnd HWND, cidl uint32, apidl *LPCITEMIDLIST, psf IShellFolder *, lpfn LPFNDFMCALLBACK, nKeys uint32, ahkeys /*const*/ *HKEY, ppcm IContextMenu * *) HRESULT

// TODO: Unknown type(s): LPDATAOBJECT *
// func CIDLData_CreateFromIDArray(pidlFolder /*const*/ LPCITEMIDLIST, cpidlFiles uint32, lppidlFiles *LPCITEMIDLIST, ppdataObject LPDATAOBJECT *) HRESULT

func CallCPLEntry16(hMod HMODULE, pFunc FARPROC, dw3 uint32, dw4 uint32, dw5 uint32, dw6 uint32) uint32 {
	pFuncCallback := syscall.NewCallback(func() uintptr {
		ret := pFunc()
		return uintptr(unsafe.Pointer(ret))
	})
	ret1 := syscall6(callCPLEntry16, 6,
		uintptr(hMod),
		pFuncCallback,
		uintptr(dw3),
		uintptr(dw4),
		uintptr(dw5),
		uintptr(dw6))
	return uint32(ret1)
}

func CheckEscapes(string LPWSTR, aLen uint32) uint32 {
	ret1 := syscall3(checkEscapes, 2,
		uintptr(unsafe.Pointer(string)),
		uintptr(aLen),
		0)
	return uint32(ret1)
}

func Control_FillCache_RunDLL(hWnd HWND, hModule HANDLE, w uint32, x uint32) HRESULT {
	ret1 := syscall6(control_FillCache_RunDLL, 4,
		uintptr(hWnd),
		uintptr(hModule),
		uintptr(w),
		uintptr(x),
		0,
		0)
	return HRESULT(ret1)
}

func Control_RunDLL(hWnd HWND, hInst HINSTANCE, cmd string, nCmdShow uint32) {
	cmdStr := unicode16FromString(cmd)
	syscall6(control_RunDLL, 4,
		uintptr(hWnd),
		uintptr(hInst),
		uintptr(unsafe.Pointer(&cmdStr[0])),
		uintptr(nCmdShow),
		0,
		0)
}

// TODO: Unknown type(s): AUTO_SCROLL_DATA *
// func DAD_AutoScroll(hwnd HWND, samples AUTO_SCROLL_DATA *, pt *POINT) bool

func DAD_DragEnterEx(hwnd HWND, p POINT) bool {
	ret1 := syscall3(dAD_DragEnterEx, 3,
		uintptr(hwnd),
		uintptr(p.X),
		uintptr(p.Y))
	return ret1 != 0
}

func DAD_DragLeave() bool {
	ret1 := syscall3(dAD_DragLeave, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func DAD_DragMove(p POINT) bool {
	ret1 := syscall3(dAD_DragMove, 2,
		uintptr(p.X),
		uintptr(p.Y),
		0)
	return ret1 != 0
}

func DAD_SetDragImage(himlTrack HIMAGELIST, lppt *POINT) bool {
	ret1 := syscall3(dAD_SetDragImage, 2,
		uintptr(himlTrack),
		uintptr(unsafe.Pointer(lppt)),
		0)
	return ret1 != 0
}

func DAD_ShowDragImage(bShow bool) bool {
	ret1 := syscall3(dAD_ShowDragImage, 1,
		getUintptrFromBool(bShow),
		0,
		0)
	return ret1 != 0
}

func DoEnvironmentSubst(pszString LPWSTR, cchString uint32) uint32 {
	ret1 := syscall3(doEnvironmentSubst, 2,
		uintptr(unsafe.Pointer(pszString)),
		uintptr(cchString),
		0)
	return uint32(ret1)
}

func DragAcceptFiles(hWnd HWND, b bool) {
	syscall3(dragAcceptFiles, 2,
		uintptr(hWnd),
		getUintptrFromBool(b),
		0)
}

// TODO: Unknown type(s): HDROP
// func DragFinish(h HDROP)

// TODO: Unknown type(s): HDROP
// func DragQueryFile(hDrop HDROP, lFile uint32, lpszwFile LPWSTR, lLength uint32) uint32

// TODO: Unknown type(s): HDROP
// func DragQueryPoint(hDrop HDROP, p *POINT) bool

func DriveType(u int32) int32 {
	ret1 := syscall3(driveType, 1,
		uintptr(u),
		0,
		0)
	return int32(ret1)
}

func DuplicateIcon(hInstance HINSTANCE, hIcon HICON) HICON {
	ret1 := syscall3(duplicateIcon, 2,
		uintptr(hInstance),
		uintptr(hIcon),
		0)
	return HICON(ret1)
}

func ExtractAssociatedIconEx(hInst HINSTANCE, lpIconPath LPWSTR, lpiIconIdx *uint16, lpiIconId *uint16) HICON {
	ret1 := syscall6(extractAssociatedIconEx, 4,
		uintptr(hInst),
		uintptr(unsafe.Pointer(lpIconPath)),
		uintptr(unsafe.Pointer(lpiIconIdx)),
		uintptr(unsafe.Pointer(lpiIconId)),
		0,
		0)
	return HICON(ret1)
}

func ExtractAssociatedIcon(hInst HINSTANCE, lpIconPath LPWSTR, lpiIcon *uint16) HICON {
	ret1 := syscall3(extractAssociatedIcon, 3,
		uintptr(hInst),
		uintptr(unsafe.Pointer(lpIconPath)),
		uintptr(unsafe.Pointer(lpiIcon)))
	return HICON(ret1)
}

func ExtractIconEx(lpszFile string, nIconIndex int32, phiconLarge *HICON, phiconSmall *HICON, nIcons uint32) uint32 {
	lpszFileStr := unicode16FromString(lpszFile)
	ret1 := syscall6(extractIconEx, 5,
		uintptr(unsafe.Pointer(&lpszFileStr[0])),
		uintptr(nIconIndex),
		uintptr(unsafe.Pointer(phiconLarge)),
		uintptr(unsafe.Pointer(phiconSmall)),
		uintptr(nIcons),
		0)
	return uint32(ret1)
}

func ExtractIcon(hInstance HINSTANCE, lpszFile string, nIconIndex uint32) HICON {
	lpszFileStr := unicode16FromString(lpszFile)
	ret1 := syscall3(extractIcon, 3,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(&lpszFileStr[0])),
		uintptr(nIconIndex))
	return HICON(ret1)
}

func ExtractVersionResource16W(s LPWSTR, d uint32) bool {
	ret1 := syscall3(extractVersionResource16W, 2,
		uintptr(unsafe.Pointer(s)),
		uintptr(d),
		0)
	return ret1 != 0
}

func FindExecutable(lpFile string, lpDirectory string, lpResult LPWSTR) HINSTANCE {
	lpFileStr := unicode16FromString(lpFile)
	lpDirectoryStr := unicode16FromString(lpDirectory)
	ret1 := syscall3(findExecutable, 3,
		uintptr(unsafe.Pointer(&lpFileStr[0])),
		uintptr(unsafe.Pointer(&lpDirectoryStr[0])),
		uintptr(unsafe.Pointer(lpResult)))
	return HINSTANCE(ret1)
}

func FreeIconList(dw uint32) {
	syscall3(freeIconList, 1,
		uintptr(dw),
		0,
		0)
}

func GetCurrentProcessExplicitAppUserModelID(appid *PWSTR) HRESULT {
	ret1 := syscall3(getCurrentProcessExplicitAppUserModelID, 1,
		uintptr(unsafe.Pointer(appid)),
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): LPCSHITEMID
// func ILAppendID(pidl *ITEMIDLIST, item LPCSHITEMID, bEnd bool) *ITEMIDLIST

func ILClone(pidl /*const*/ LPCITEMIDLIST) *ITEMIDLIST {
	ret1 := syscall3(iLClone, 1,
		uintptr(unsafe.Pointer(pidl)),
		0,
		0)
	return (*ITEMIDLIST)(unsafe.Pointer(ret1))
}

func ILCloneFirst(pidl /*const*/ LPCITEMIDLIST) *ITEMIDLIST {
	ret1 := syscall3(iLCloneFirst, 1,
		uintptr(unsafe.Pointer(pidl)),
		0,
		0)
	return (*ITEMIDLIST)(unsafe.Pointer(ret1))
}

func ILCombine(pidl1 /*const*/ LPCITEMIDLIST, pidl2 /*const*/ LPCITEMIDLIST) *ITEMIDLIST {
	ret1 := syscall3(iLCombine, 2,
		uintptr(unsafe.Pointer(pidl1)),
		uintptr(unsafe.Pointer(pidl2)),
		0)
	return (*ITEMIDLIST)(unsafe.Pointer(ret1))
}

func ILCreateFromPath(path string) *ITEMIDLIST {
	pathStr := unicode16FromString(path)
	ret1 := syscall3(iLCreateFromPath, 1,
		uintptr(unsafe.Pointer(&pathStr[0])),
		0,
		0)
	return (*ITEMIDLIST)(unsafe.Pointer(ret1))
}

func ILFindChild(pidl1 /*const*/ LPCITEMIDLIST, pidl2 /*const*/ LPCITEMIDLIST) *ITEMIDLIST {
	ret1 := syscall3(iLFindChild, 2,
		uintptr(unsafe.Pointer(pidl1)),
		uintptr(unsafe.Pointer(pidl2)),
		0)
	return (*ITEMIDLIST)(unsafe.Pointer(ret1))
}

func ILFindLastID(pidl /*const*/ LPCITEMIDLIST) *ITEMIDLIST {
	ret1 := syscall3(iLFindLastID, 1,
		uintptr(unsafe.Pointer(pidl)),
		0,
		0)
	return (*ITEMIDLIST)(unsafe.Pointer(ret1))
}

func ILFree(pidl *ITEMIDLIST) {
	syscall3(iLFree, 1,
		uintptr(unsafe.Pointer(pidl)),
		0,
		0)
}

func ILGetNext(pidl /*const*/ LPCITEMIDLIST) *ITEMIDLIST {
	ret1 := syscall3(iLGetNext, 1,
		uintptr(unsafe.Pointer(pidl)),
		0,
		0)
	return (*ITEMIDLIST)(unsafe.Pointer(ret1))
}

func ILGetSize(pidl /*const*/ LPCITEMIDLIST) uint32 {
	ret1 := syscall3(iLGetSize, 1,
		uintptr(unsafe.Pointer(pidl)),
		0,
		0)
	return uint32(ret1)
}

func ILIsEqual(pidl1 /*const*/ LPCITEMIDLIST, pidl2 /*const*/ LPCITEMIDLIST) bool {
	ret1 := syscall3(iLIsEqual, 2,
		uintptr(unsafe.Pointer(pidl1)),
		uintptr(unsafe.Pointer(pidl2)),
		0)
	return ret1 != 0
}

func ILIsParent(pidlParent /*const*/ LPCITEMIDLIST, pidlChild /*const*/ LPCITEMIDLIST, bImmediate bool) bool {
	ret1 := syscall3(iLIsParent, 3,
		uintptr(unsafe.Pointer(pidlParent)),
		uintptr(unsafe.Pointer(pidlChild)),
		getUintptrFromBool(bImmediate))
	return ret1 != 0
}

// TODO: Unknown type(s): LPITEMIDLIST *
// func ILLoadFromStream(pStream *IStream, ppPidl LPITEMIDLIST *) HRESULT

func ILRemoveLastID(pidl *ITEMIDLIST) bool {
	ret1 := syscall3(iLRemoveLastID, 1,
		uintptr(unsafe.Pointer(pidl)),
		0,
		0)
	return ret1 != 0
}

func ILSaveToStream(pStream *IStream, pPidl /*const*/ LPCITEMIDLIST) HRESULT {
	ret1 := syscall3(iLSaveToStream, 2,
		uintptr(unsafe.Pointer(pStream)),
		uintptr(unsafe.Pointer(pPidl)),
		0)
	return HRESULT(ret1)
}

func InitNetworkAddressControl() bool {
	ret1 := syscall3(initNetworkAddressControl, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func IsLFNDrive(lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(isLFNDrive, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func IsNetDrive(drive int32) int32 {
	ret1 := syscall3(isNetDrive, 1,
		uintptr(drive),
		0,
		0)
	return int32(ret1)
}

func IsUserAnAdmin() bool {
	ret1 := syscall3(isUserAnAdmin, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func OpenAs_RunDLL(hwnd HWND, hinst HINSTANCE, cmdline string, cmdshow int32) {
	cmdlineStr := unicode16FromString(cmdline)
	syscall6(openAs_RunDLL, 4,
		uintptr(hwnd),
		uintptr(hinst),
		uintptr(unsafe.Pointer(&cmdlineStr[0])),
		uintptr(cmdshow),
		0,
		0)
}

func PathCleanupSpec(lpszPathW string, lpszFileW LPWSTR) int32 {
	lpszPathWStr := unicode16FromString(lpszPathW)
	ret1 := syscall3(pathCleanupSpec, 2,
		uintptr(unsafe.Pointer(&lpszPathWStr[0])),
		uintptr(unsafe.Pointer(lpszFileW)),
		0)
	return int32(ret1)
}

func PathYetAnotherMakeUniqueName(buffer LPWSTR, path string, shortname string, longname string) bool {
	pathStr := unicode16FromString(path)
	shortnameStr := unicode16FromString(shortname)
	longnameStr := unicode16FromString(longname)
	ret1 := syscall6(pathYetAnotherMakeUniqueName, 4,
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(&pathStr[0])),
		uintptr(unsafe.Pointer(&shortnameStr[0])),
		uintptr(unsafe.Pointer(&longnameStr[0])),
		0,
		0)
	return ret1 != 0
}

func PickIconDlg(hwndOwner HWND, lpstrFile LPSTR, nMaxFile uint32, lpdwIconIndex *uint32) int32 {
	ret1 := syscall6(pickIconDlg, 4,
		uintptr(hwndOwner),
		uintptr(unsafe.Pointer(lpstrFile)),
		uintptr(nMaxFile),
		uintptr(unsafe.Pointer(lpdwIconIndex)),
		0,
		0)
	return int32(ret1)
}

// TODO: Unknown type(s): CABINETSTATE *
// func ReadCabinetState(cs CABINETSTATE *, length int32) bool

func RealDriveType(drive int32, bQueryNet bool) int32 {
	ret1 := syscall3(realDriveType, 2,
		uintptr(drive),
		getUintptrFromBool(bQueryNet),
		0)
	return int32(ret1)
}

func RegenerateUserEnvironment(wunknown *WCHAR, bunknown bool) bool {
	ret1 := syscall3(regenerateUserEnvironment, 2,
		uintptr(unsafe.Pointer(wunknown)),
		getUintptrFromBool(bunknown),
		0)
	return ret1 != 0
}

func RestartDialog(hWndOwner HWND, lpstrReason string, uFlags uint32) int32 {
	lpstrReasonStr := unicode16FromString(lpstrReason)
	ret1 := syscall3(restartDialog, 3,
		uintptr(hWndOwner),
		uintptr(unsafe.Pointer(&lpstrReasonStr[0])),
		uintptr(uFlags))
	return int32(ret1)
}

func RestartDialogEx(hWndOwner HWND, lpwstrReason string, uFlags uint32, uReason uint32) int32 {
	lpwstrReasonStr := unicode16FromString(lpwstrReason)
	ret1 := syscall6(restartDialogEx, 4,
		uintptr(hWndOwner),
		uintptr(unsafe.Pointer(&lpwstrReasonStr[0])),
		uintptr(uFlags),
		uintptr(uReason),
		0,
		0)
	return int32(ret1)
}

func SHAddToRecentDocs(uFlags uint32, pv /*const*/ uintptr) {
	syscall3(sHAddToRecentDocs, 2,
		uintptr(uFlags),
		pv,
		0)
}

func SHAlloc(aLen uint32) LPVOID {
	ret1 := syscall3(sHAlloc, 1,
		uintptr(aLen),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): PAPPBARDATA
// func SHAppBarMessage(msg uint32, data PAPPBARDATA) *uint32

// TODO: Unknown type(s): ASSOC_FILTER, IEnumAssocHandlers * *
// func SHAssocEnumHandlers(extra /*const*/ *WCHAR, filter ASSOC_FILTER, enumhandlers IEnumAssocHandlers * *) HRESULT

func SHBindToParent(pidl /*const*/ LPCITEMIDLIST, riid REFIID, ppv *LPVOID, ppidlLast *LPCITEMIDLIST) HRESULT {
	ret1 := syscall6(sHBindToParent, 4,
		uintptr(unsafe.Pointer(pidl)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppv)),
		uintptr(unsafe.Pointer(ppidlLast)),
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): LPBROWSEINFOW
// func SHBrowseForFolder(lpbi LPBROWSEINFOW) *ITEMIDLIST

// TODO: Unknown type(s): LPITEMIDLIST * *
// func SHChangeNotification_Lock(hChange HANDLE, dwProcessId uint32, lppidls LPITEMIDLIST * *, lpwEventId *LONG) HANDLE

func SHChangeNotification_Unlock(hLock HANDLE) bool {
	ret1 := syscall3(sHChangeNotification_Unlock, 1,
		uintptr(hLock),
		0,
		0)
	return ret1 != 0
}

func SHChangeNotify(wEventId LONG, uFlags uint32, dwItem1 /*const*/ uintptr, dwItem2 /*const*/ uintptr) {
	syscall6(sHChangeNotify, 4,
		uintptr(wEventId),
		uintptr(uFlags),
		dwItem1,
		dwItem2,
		0,
		0)
}

func SHChangeNotifyDeregister(hNotify ULONG) bool {
	ret1 := syscall3(sHChangeNotifyDeregister, 1,
		uintptr(hNotify),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): SHChangeNotifyEntry *
// func SHChangeNotifyRegister(hwnd HWND, fSources int32, wEventMask LONG, uMsg uint32, cItems int32, lpItems SHChangeNotifyEntry *) ULONG

func SHCloneSpecialIDList(hwndOwner HWND, nFolder uint32, fCreate bool) *ITEMIDLIST {
	ret1 := syscall3(sHCloneSpecialIDList, 3,
		uintptr(hwndOwner),
		uintptr(nFolder),
		getUintptrFromBool(fCreate))
	return (*ITEMIDLIST)(unsafe.Pointer(ret1))
}

func SHCoCreateInstance(aclsid string, clsid /*const*/ *CLSID, pUnkOuter LPUNKNOWN, refiid REFIID, ppv *LPVOID) HRESULT {
	aclsidStr := unicode16FromString(aclsid)
	ret1 := syscall6(sHCoCreateInstance, 5,
		uintptr(unsafe.Pointer(&aclsidStr[0])),
		uintptr(unsafe.Pointer(clsid)),
		uintptr(unsafe.Pointer(pUnkOuter)),
		uintptr(unsafe.Pointer(refiid)),
		uintptr(unsafe.Pointer(ppv)),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): const DEFCONTEXTMENU *
// func SHCreateDefaultContextMenu(pdcm /*const*/ const DEFCONTEXTMENU *, riid REFIID, ppv uintptr) HRESULT

func SHCreateDirectory(hWnd HWND, path /*const*/ uintptr) uint32 {
	ret1 := syscall3(sHCreateDirectory, 2,
		uintptr(hWnd),
		path,
		0)
	return uint32(ret1)
}

func SHCreateDirectoryEx(hWnd HWND, path string, sec *SECURITY_ATTRIBUTES) int32 {
	pathStr := unicode16FromString(path)
	ret1 := syscall3(sHCreateDirectoryEx, 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&pathStr[0])),
		uintptr(unsafe.Pointer(sec)))
	return int32(ret1)
}

func SHCreateFileExtractIconW(file string, attribs uint32, riid REFIID, ppv uintptr) HRESULT {
	fileStr := unicode16FromString(file)
	ret1 := syscall6(sHCreateFileExtractIconW, 4,
		uintptr(unsafe.Pointer(&fileStr[0])),
		uintptr(attribs),
		uintptr(unsafe.Pointer(riid)),
		ppv,
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): PCIDLIST_ABSOLUTE
// func SHCreateItemFromIDList(pidl PCIDLIST_ABSOLUTE, riid REFIID, ppv uintptr) HRESULT

func SHCreateItemFromParsingName(pszPath string, pbc *IBindCtx, riid REFIID, ppv uintptr) HRESULT {
	pszPathStr := unicode16FromString(pszPath)
	ret1 := syscall6(sHCreateItemFromParsingName, 4,
		uintptr(unsafe.Pointer(&pszPathStr[0])),
		uintptr(unsafe.Pointer(pbc)),
		uintptr(unsafe.Pointer(riid)),
		ppv,
		0,
		0)
	return HRESULT(ret1)
}

func SHCreateQueryCancelAutoPlayMoniker(moniker **IMoniker) HRESULT {
	ret1 := syscall3(sHCreateQueryCancelAutoPlayMoniker, 1,
		uintptr(unsafe.Pointer(moniker)),
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IShellView * *, const SFV_CREATE *
// func SHCreateShellFolderView(pcsfv /*const*/ const SFV_CREATE *, ppsv IShellView * *) HRESULT

// TODO: Unknown type(s): IShellView * *, LPCSFV
// func SHCreateShellFolderViewEx(psvcbi LPCSFV, ppv IShellView * *) HRESULT

// TODO: Unknown type(s): IShellFolder *, IShellItem * *
// func SHCreateShellItem(pidlParent /*const*/ LPCITEMIDLIST, psfParent IShellFolder *, pidl /*const*/ LPCITEMIDLIST, ppsi IShellItem * *) HRESULT

// TODO: Unknown type(s): IShellFolder *, IShellItemArray * *, PCIDLIST_ABSOLUTE, PCUITEMID_CHILD_ARRAY
// func SHCreateShellItemArray(pidlParent PCIDLIST_ABSOLUTE, psf IShellFolder *, cidl uint32, ppidl PCUITEMID_CHILD_ARRAY, ppsiItemArray IShellItemArray * *) HRESULT

// TODO: Unknown type(s): IDataObject *
// func SHCreateShellItemArrayFromDataObject(pdo IDataObject *, riid REFIID, ppv uintptr) HRESULT

// TODO: Unknown type(s): IShellItemArray * *, PCIDLIST_ABSOLUTE_ARRAY
// func SHCreateShellItemArrayFromIDLists(cidl uint32, pidl_array PCIDLIST_ABSOLUTE_ARRAY, psia IShellItemArray * *) HRESULT

// TODO: Unknown type(s): IShellItem *
// func SHCreateShellItemArrayFromShellItem(item IShellItem *, riid REFIID, ppv uintptr) HRESULT

// TODO: Unknown type(s): LPENUMFORMATETC *, const FORMATETC *
// func SHCreateStdEnumFmtEtc(cFormats uint32, lpFormats /*const*/ const FORMATETC *, ppenumFormatetc LPENUMFORMATETC *) HRESULT

func SHDefExtractIcon(pszIconFile string, iIndex int32, uFlags uint32, phiconLarge *HICON, phiconSmall *HICON, nIconSize uint32) HRESULT {
	pszIconFileStr := unicode16FromString(pszIconFile)
	ret1 := syscall6(sHDefExtractIcon, 6,
		uintptr(unsafe.Pointer(&pszIconFileStr[0])),
		uintptr(iIndex),
		uintptr(uFlags),
		uintptr(unsafe.Pointer(phiconLarge)),
		uintptr(unsafe.Pointer(phiconSmall)),
		uintptr(nIconSize))
	return HRESULT(ret1)
}

// TODO: Unknown type(s): LPDATAOBJECT, LPDROPSOURCE
// func SHDoDragDrop(hWnd HWND, lpDataObject LPDATAOBJECT, lpDropSource LPDROPSOURCE, dwOKEffect uint32, pdwEffect *uint32) HRESULT

func SHEmptyRecycleBin(hwnd HWND, pszRootPath string, dwFlags uint32) HRESULT {
	pszRootPathStr := unicode16FromString(pszRootPath)
	ret1 := syscall3(sHEmptyRecycleBin, 3,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pszRootPathStr[0])),
		uintptr(dwFlags))
	return HRESULT(ret1)
}

func SHEnumerateUnreadMailAccountsW(user HKEY, idx uint32, mailaddress LPWSTR, mailaddresslen int32) HRESULT {
	ret1 := syscall6(sHEnumerateUnreadMailAccountsW, 4,
		uintptr(user),
		uintptr(idx),
		uintptr(unsafe.Pointer(mailaddress)),
		uintptr(mailaddresslen),
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): LPSHFILEOPSTRUCTW
// func SHFileOperation(lpFileOp LPSHFILEOPSTRUCTW) int32

func SHFindFiles(pidlFolder /*const*/ LPCITEMIDLIST, pidlSaveFile /*const*/ LPCITEMIDLIST) bool {
	ret1 := syscall3(sHFindFiles, 2,
		uintptr(unsafe.Pointer(pidlFolder)),
		uintptr(unsafe.Pointer(pidlSaveFile)),
		0)
	return ret1 != 0
}

func SHFind_InitMenuPopup(hMenu HMENU, hWndParent HWND, w uint32, x uint32) LPVOID {
	ret1 := syscall6(sHFind_InitMenuPopup, 4,
		uintptr(hMenu),
		uintptr(hWndParent),
		uintptr(w),
		uintptr(x),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func SHFlushClipboard() HRESULT {
	ret1 := syscall3(sHFlushClipboard, 0,
		0,
		0,
		0)
	return HRESULT(ret1)
}

func SHFlushSFCache() {
	syscall3(sHFlushSFCache, 0,
		0,
		0,
		0)
}

func SHFormatDrive(hwnd HWND, drive uint32, fmtID uint32, options uint32) uint32 {
	ret1 := syscall6(sHFormatDrive, 4,
		uintptr(hwnd),
		uintptr(drive),
		uintptr(fmtID),
		uintptr(options),
		0,
		0)
	return uint32(ret1)
}

func SHFree(pv LPVOID) {
	syscall3(sHFree, 1,
		uintptr(unsafe.Pointer(pv)),
		0,
		0)
}

func SHFreeNameMappings(hNameMapping HANDLE) {
	syscall3(sHFreeNameMappings, 1,
		uintptr(hNameMapping),
		0,
		0)
}

// TODO: Unknown type(s): LPSHELLFOLDER
// func SHGetDataFromIDList(psf LPSHELLFOLDER, pidl /*const*/ LPCITEMIDLIST, nFormat int32, dest LPVOID, aLen int32) HRESULT

// TODO: Unknown type(s): IShellFolder * *
// func SHGetDesktopFolder(psf IShellFolder * *) HRESULT

// TODO: Unknown type(s): SHFILEINFOW *
// func SHGetFileInfo(path string, dwFileAttributes uint32, psfi SHFILEINFOW *, sizeofpsfi uint32, flags uint32) *uint32

// TODO: Unknown type(s): LPITEMIDLIST *
// func SHGetFolderLocation(hwndOwner HWND, nFolder int32, hToken HANDLE, dwReserved uint32, ppidl LPITEMIDLIST *) HRESULT

func SHGetFolderPathAndSubDir(hwndOwner HWND, nFolder int32, hToken HANDLE, dwFlags uint32, pszSubPath string, pszPath LPWSTR) HRESULT {
	pszSubPathStr := unicode16FromString(pszSubPath)
	ret1 := syscall6(sHGetFolderPathAndSubDir, 6,
		uintptr(hwndOwner),
		uintptr(nFolder),
		uintptr(hToken),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(&pszSubPathStr[0])),
		uintptr(unsafe.Pointer(pszPath)))
	return HRESULT(ret1)
}

// TODO: Unknown type(s): REFKNOWNFOLDERID
// func SHGetFolderPathEx(rfid REFKNOWNFOLDERID, flags uint32, token HANDLE, path LPWSTR, aLen uint32) HRESULT

func SHGetFolderPath(hwndOwner HWND, nFolder int32, hToken HANDLE, dwFlags uint32, pszPath LPWSTR) HRESULT {
	ret1 := syscall6(sHGetFolderPath, 5,
		uintptr(hwndOwner),
		uintptr(nFolder),
		uintptr(hToken),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pszPath)),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): PIDLIST_ABSOLUTE *
// func SHGetIDListFromObject(punk *IUnknown, ppidl PIDLIST_ABSOLUTE *) HRESULT

func SHGetIconOverlayIndex(pszIconPath string, iIconIndex int32) int32 {
	pszIconPathStr := unicode16FromString(pszIconPath)
	ret1 := syscall3(sHGetIconOverlayIndex, 2,
		uintptr(unsafe.Pointer(&pszIconPathStr[0])),
		uintptr(iIconIndex),
		0)
	return int32(ret1)
}

func SHGetImageList(iImageList int32, riid REFIID, ppv uintptr) HRESULT {
	ret1 := syscall3(sHGetImageList, 3,
		uintptr(iImageList),
		uintptr(unsafe.Pointer(riid)),
		ppv)
	return HRESULT(ret1)
}

func SHGetInstanceExplorer(lpUnknown **IUnknown) HRESULT {
	ret1 := syscall3(sHGetInstanceExplorer, 1,
		uintptr(unsafe.Pointer(lpUnknown)),
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): DATAOBJ_GET_ITEM_FLAGS, IDataObject *
// func SHGetItemFromDataObject(pdtobj IDataObject *, dwFlags DATAOBJ_GET_ITEM_FLAGS, riid REFIID, ppv uintptr) HRESULT

func SHGetItemFromObject(punk *IUnknown, riid REFIID, ppv uintptr) HRESULT {
	ret1 := syscall3(sHGetItemFromObject, 3,
		uintptr(unsafe.Pointer(punk)),
		uintptr(unsafe.Pointer(riid)),
		ppv)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): PIDLIST_ABSOLUTE *, REFKNOWNFOLDERID
// func SHGetKnownFolderIDList(rfid REFKNOWNFOLDERID, flags uint32, token HANDLE, pidl PIDLIST_ABSOLUTE *) HRESULT

// TODO: Unknown type(s): KNOWN_FOLDER_FLAG, REFKNOWNFOLDERID
// func SHGetKnownFolderItem(rfid REFKNOWNFOLDERID, flags KNOWN_FOLDER_FLAG, hToken HANDLE, riid REFIID, ppv uintptr) HRESULT

// TODO: Unknown type(s): REFKNOWNFOLDERID
// func SHGetKnownFolderPath(rfid REFKNOWNFOLDERID, flags uint32, token HANDLE, path *PWSTR) HRESULT

func SHGetLocalizedName(path string, module LPWSTR, size uint32, res *int32) HRESULT {
	pathStr := unicode16FromString(path)
	ret1 := syscall6(sHGetLocalizedName, 4,
		uintptr(unsafe.Pointer(&pathStr[0])),
		uintptr(unsafe.Pointer(module)),
		uintptr(size),
		uintptr(unsafe.Pointer(res)),
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): LPMALLOC *
// func SHGetMalloc(lpmal LPMALLOC *) HRESULT

// TODO: Unknown type(s): PCIDLIST_ABSOLUTE, SIGDN
// func SHGetNameFromIDList(pidl PCIDLIST_ABSOLUTE, sigdnName SIGDN, ppszName *PWSTR) HRESULT

func SHGetNewLinkInfo(pszLinkTo string, pszDir string, pszName LPWSTR, pfMustCopy *BOOL, uFlags uint32) bool {
	pszLinkToStr := unicode16FromString(pszLinkTo)
	pszDirStr := unicode16FromString(pszDir)
	ret1 := syscall6(sHGetNewLinkInfo, 5,
		uintptr(unsafe.Pointer(&pszLinkToStr[0])),
		uintptr(unsafe.Pointer(&pszDirStr[0])),
		uintptr(unsafe.Pointer(pszName)),
		uintptr(unsafe.Pointer(pfMustCopy)),
		uintptr(uFlags),
		0)
	return ret1 != 0
}

func SHGetPathFromIDList(pidl /*const*/ LPCITEMIDLIST, pszPath LPWSTR) bool {
	ret1 := syscall3(sHGetPathFromIDList, 2,
		uintptr(unsafe.Pointer(pidl)),
		uintptr(unsafe.Pointer(pszPath)),
		0)
	return ret1 != 0
}

func SHGetPropertyStoreForWindow(hwnd HWND, riid REFIID, ppv uintptr) HRESULT {
	ret1 := syscall3(sHGetPropertyStoreForWindow, 3,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(riid)),
		ppv)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): GETPROPERTYSTOREFLAGS
// func SHGetPropertyStoreFromParsingName(pszPath string, pbc *IBindCtx, flags GETPROPERTYSTOREFLAGS, riid REFIID, ppv uintptr) HRESULT

// TODO: Unknown type(s): LPITEMIDLIST *, LPSHELLFOLDER
// func SHGetRealIDL(lpsf LPSHELLFOLDER, pidlSimple /*const*/ LPCITEMIDLIST, pidlReal LPITEMIDLIST *) HRESULT

// TODO: Unknown type(s): LPSHELLSTATE
// func SHGetSetSettings(lpss LPSHELLSTATE, dwMask uint32, bSet bool)

// TODO: Unknown type(s): LPSHELLFLAGSTATE
// func SHGetSettings(lpsfs LPSHELLFLAGSTATE, dwMask uint32)

// TODO: Unknown type(s): LPITEMIDLIST *
// func SHGetSpecialFolderLocation(hwndOwner HWND, nFolder int32, ppidl LPITEMIDLIST *) HRESULT

func SHGetSpecialFolderPath(hwndOwner HWND, szPath LPWSTR, nFolder int32, bCreate bool) bool {
	ret1 := syscall6(sHGetSpecialFolderPath, 4,
		uintptr(hwndOwner),
		uintptr(unsafe.Pointer(szPath)),
		uintptr(nFolder),
		getUintptrFromBool(bCreate),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): SHSTOCKICONID, SHSTOCKICONINFO *
// func SHGetStockIconInfo(id SHSTOCKICONID, flags uint32, sii SHSTOCKICONINFO *) HRESULT

func SHHandleUpdateImage(pidlExtra /*const*/ LPCITEMIDLIST) int32 {
	ret1 := syscall3(sHHandleUpdateImage, 1,
		uintptr(unsafe.Pointer(pidlExtra)),
		0,
		0)
	return int32(ret1)
}

func SHHelpShortcuts_RunDLL(dwArg1 uint32, dwArg2 uint32, dwArg3 uint32, dwArg4 uint32) uint32 {
	ret1 := syscall6(sHHelpShortcuts_RunDLL, 4,
		uintptr(dwArg1),
		uintptr(dwArg2),
		uintptr(dwArg3),
		uintptr(dwArg4),
		0,
		0)
	return uint32(ret1)
}

func SHIsFileAvailableOffline(path string, status *uint32) HRESULT {
	pathStr := unicode16FromString(path)
	ret1 := syscall3(sHIsFileAvailableOffline, 2,
		uintptr(unsafe.Pointer(&pathStr[0])),
		uintptr(unsafe.Pointer(status)),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IShellFolder *
// func SHLimitInputEdit(textbox HWND, folder IShellFolder *) HRESULT

func SHLoadInProc(rclsid /*const*/ REFCLSID) HRESULT {
	ret1 := syscall3(sHLoadInProc, 1,
		uintptr(unsafe.Pointer(rclsid)),
		0,
		0)
	return HRESULT(ret1)
}

func SHLoadNonloadedIconOverlayIdentifiers() HRESULT {
	ret1 := syscall3(sHLoadNonloadedIconOverlayIdentifiers, 0,
		0,
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IShellFolder *
// func SHMapIDListToImageListIndexAsync(pts *IUnknown, psf IShellFolder *, pidl /*const*/ LPCITEMIDLIST, flags uint32, pfn uintptr, pvData uintptr, pvHint uintptr, piIndex *int, piIndexSel *int) HRESULT

// TODO: Unknown type(s): IShellFolder *
// func SHMapPIDLToSystemImageListIndex(sh IShellFolder *, pidl /*const*/ LPCITEMIDLIST, pIndex *int) int32

func SHObjectProperties(hwnd HWND, dwType uint32, szObject string, szPage string) bool {
	szObjectStr := unicode16FromString(szObject)
	szPageStr := unicode16FromString(szPage)
	ret1 := syscall6(sHObjectProperties, 4,
		uintptr(hwnd),
		uintptr(dwType),
		uintptr(unsafe.Pointer(&szObjectStr[0])),
		uintptr(unsafe.Pointer(&szPageStr[0])),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCIDLIST_ABSOLUTE, PCUITEMID_CHILD_ARRAY *
// func SHOpenFolderAndSelectItems(pidlFolder PCIDLIST_ABSOLUTE, cidl uint32, apidl PCUITEMID_CHILD_ARRAY *, flags uint32) HRESULT

// TODO: Unknown type(s): LPITEMIDLIST *, SFGAOF, SFGAOF *
// func SHParseDisplayName(name string, bindctx *IBindCtx, pidlist LPITEMIDLIST *, attr_in SFGAOF, attr_out SFGAOF *) HRESULT

func SHPathPrepareForWrite(hwnd HWND, modless *IUnknown, path string, flags uint32) HRESULT {
	pathStr := unicode16FromString(path)
	ret1 := syscall6(sHPathPrepareForWrite, 4,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(modless)),
		uintptr(unsafe.Pointer(&pathStr[0])),
		uintptr(flags),
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IPropertySetStorage *, IPropertyStorage * *, REFFMTID
// func SHPropStgCreate(psstg IPropertySetStorage *, fmtid REFFMTID, pclsid /*const*/ *CLSID, grfFlags uint32, grfMode uint32, dwDisposition uint32, ppstg IPropertyStorage * *, puCodePage *UINT) HRESULT

// TODO: Unknown type(s): IPropertyStorage *, PROPVARIANT *, const PROPSPEC *
// func SHPropStgReadMultiple(pps IPropertyStorage *, uCodePage uint32, cpspec ULONG, rgpspec /*const*/ const PROPSPEC *, rgvar PROPVARIANT *) HRESULT

// TODO: Unknown type(s): IPropertyStorage *, PROPID, PROPVARIANT *, const PROPSPEC *
// func SHPropStgWriteMultiple(pps IPropertyStorage *, uCodePage *UINT, cpspec ULONG, rgpspec /*const*/ const PROPSPEC *, rgvar PROPVARIANT *, propidNameFirst PROPID) HRESULT

// TODO: Unknown type(s): LPSHQUERYRBINFO
// func SHQueryRecycleBin(pszRootPath string, pSHQueryRBInfo LPSHQUERYRBINFO) HRESULT

// TODO: Unknown type(s): QUERY_USER_NOTIFICATION_STATE *
// func SHQueryUserNotificationState(state QUERY_USER_NOTIFICATION_STATE *) HRESULT

func SHRemoveLocalizedName(path /*const*/ *WCHAR) HRESULT {
	ret1 := syscall3(sHRemoveLocalizedName, 1,
		uintptr(unsafe.Pointer(path)),
		0,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): RESTRICTIONS
// func SHRestricted(policy RESTRICTIONS) uint32

func SHRunControlPanel(commandLine string, parent HWND) bool {
	commandLineStr := unicode16FromString(commandLine)
	ret1 := syscall3(sHRunControlPanel, 2,
		uintptr(unsafe.Pointer(&commandLineStr[0])),
		uintptr(parent),
		0)
	return ret1 != 0
}

func SHSetInstanceExplorer(lpUnknown LPUNKNOWN) {
	syscall3(sHSetInstanceExplorer, 1,
		uintptr(unsafe.Pointer(lpUnknown)),
		0,
		0)
}

func SHSetLocalizedName(pszPath LPWSTR, pszResModule string, idsRes int32) HRESULT {
	pszResModuleStr := unicode16FromString(pszResModule)
	ret1 := syscall3(sHSetLocalizedName, 3,
		uintptr(unsafe.Pointer(pszPath)),
		uintptr(unsafe.Pointer(&pszResModuleStr[0])),
		uintptr(idsRes))
	return HRESULT(ret1)
}

func SHSetUnreadMailCountW(mailaddress string, count uint32, executecommand string) HRESULT {
	mailaddressStr := unicode16FromString(mailaddress)
	executecommandStr := unicode16FromString(executecommand)
	ret1 := syscall3(sHSetUnreadMailCountW, 3,
		uintptr(unsafe.Pointer(&mailaddressStr[0])),
		uintptr(count),
		uintptr(unsafe.Pointer(&executecommandStr[0])))
	return HRESULT(ret1)
}

func SHShellFolderView_Message(hwndCabinet HWND, uMessage uint32, lParam LPARAM) LRESULT {
	ret1 := syscall3(sHShellFolderView_Message, 3,
		uintptr(hwndCabinet),
		uintptr(uMessage),
		uintptr(lParam))
	return LRESULT(ret1)
}

func SHUpdateImage(pszHashItem string, iIndex int32, uFlags uint32, iImageIndex int32) {
	pszHashItemStr := unicode16FromString(pszHashItem)
	syscall6(sHUpdateImage, 4,
		uintptr(unsafe.Pointer(&pszHashItemStr[0])),
		uintptr(iIndex),
		uintptr(uFlags),
		uintptr(iImageIndex),
		0,
		0)
}

func SHUpdateRecycleBinIcon() HRESULT {
	ret1 := syscall3(sHUpdateRecycleBinIcon, 0,
		0,
		0,
		0)
	return HRESULT(ret1)
}

func SHValidateUNC(hwndOwner HWND, pszFile PWSTR, fConnect uint32) bool {
	ret1 := syscall3(sHValidateUNC, 3,
		uintptr(hwndOwner),
		uintptr(unsafe.Pointer(pszFile)),
		uintptr(fConnect))
	return ret1 != 0
}

func SetCurrentProcessExplicitAppUserModelID(appid string) HRESULT {
	appidStr := unicode16FromString(appid)
	ret1 := syscall3(setCurrentProcessExplicitAppUserModelID, 1,
		uintptr(unsafe.Pointer(&appidStr[0])),
		0,
		0)
	return HRESULT(ret1)
}

func SheChangeDir(path LPWSTR) uint32 {
	ret1 := syscall3(sheChangeDir, 1,
		uintptr(unsafe.Pointer(path)),
		0,
		0)
	return uint32(ret1)
}

func SheGetDir(drive uint32, buffer LPWSTR) uint32 {
	ret1 := syscall3(sheGetDir, 2,
		uintptr(drive),
		uintptr(unsafe.Pointer(buffer)),
		0)
	return uint32(ret1)
}

func ShellAbout(hWnd HWND, szApp string, szOtherStuff string, hIcon HICON) bool {
	szAppStr := unicode16FromString(szApp)
	szOtherStuffStr := unicode16FromString(szOtherStuff)
	ret1 := syscall6(shellAbout, 4,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&szAppStr[0])),
		uintptr(unsafe.Pointer(&szOtherStuffStr[0])),
		uintptr(hIcon),
		0,
		0)
	return ret1 != 0
}

func ShellExecute(hwnd HWND, lpVerb string, lpFile string, lpParameters string, lpDirectory string, nShowCmd int32) HINSTANCE {
	lpVerbStr := unicode16FromString(lpVerb)
	lpFileStr := unicode16FromString(lpFile)
	lpParametersStr := unicode16FromString(lpParameters)
	lpDirectoryStr := unicode16FromString(lpDirectory)
	ret1 := syscall6(shellExecute, 6,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&lpVerbStr[0])),
		uintptr(unsafe.Pointer(&lpFileStr[0])),
		uintptr(unsafe.Pointer(&lpParametersStr[0])),
		uintptr(unsafe.Pointer(&lpDirectoryStr[0])),
		uintptr(nShowCmd))
	return HINSTANCE(ret1)
}

func Shell_GetImageLists(lpBigList *HIMAGELIST, lpSmallList *HIMAGELIST) bool {
	ret1 := syscall3(shell_GetImageLists, 2,
		uintptr(unsafe.Pointer(lpBigList)),
		uintptr(unsafe.Pointer(lpSmallList)),
		0)
	return ret1 != 0
}

func Shell_MergeMenus(hmDst HMENU, hmSrc HMENU, uInsert uint32, uIDAdjust uint32, uIDAdjustMax uint32, uFlags ULONG) uint32 {
	ret1 := syscall6(shell_MergeMenus, 6,
		uintptr(hmDst),
		uintptr(hmSrc),
		uintptr(uInsert),
		uintptr(uIDAdjust),
		uintptr(uIDAdjustMax),
		uintptr(uFlags))
	return uint32(ret1)
}

// TODO: Unknown type(s): PNOTIFYICONDATAW
// func Shell_NotifyIcon(dwMessage uint32, nid PNOTIFYICONDATAW) bool

// TODO: Unknown type(s): PCIDLIST_ABSOLUTE
// func SignalFileOpen(pidl PCIDLIST_ABSOLUTE) bool

func WOWShellExecute(hWnd HWND, lpOperation /*const*/ LPCSTR, lpFile /*const*/ LPCSTR, lpParameters /*const*/ LPCSTR, lpDirectory /*const*/ LPCSTR, iShowCmd int32, callback uintptr) HINSTANCE {
	ret1 := syscall9(wOWShellExecute, 7,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpOperation)),
		uintptr(unsafe.Pointer(lpFile)),
		uintptr(unsafe.Pointer(lpParameters)),
		uintptr(unsafe.Pointer(lpDirectory)),
		uintptr(iShowCmd),
		callback,
		0,
		0)
	return HINSTANCE(ret1)
}

// TODO: Unknown type(s): CABINETSTATE *
// func WriteCabinetState(cs CABINETSTATE *) bool
