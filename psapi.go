// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	libpsapi uintptr

	// Functions
	emptyWorkingSet             uintptr
	enumDeviceDrivers           uintptr
	enumPageFiles               uintptr
	enumProcessModules          uintptr
	enumProcessModulesEx        uintptr
	enumProcesses               uintptr
	getDeviceDriverBaseName     uintptr
	getDeviceDriverFileName     uintptr
	getMappedFileName           uintptr
	getModuleBaseName           uintptr
	getModuleFileNameEx         uintptr
	getModuleInformation        uintptr
	getPerformanceInfo          uintptr
	getProcessImageFileName     uintptr
	getProcessMemoryInfo        uintptr
	getWsChanges                uintptr
	getWsChangesEx              uintptr
	initializeProcessForWsWatch uintptr
	queryWorkingSet             uintptr
	queryWorkingSetEx           uintptr
)

func init() {
	// Library
	libpsapi = doLoadLibrary("psapi.dll")

	// Functions
	emptyWorkingSet = doGetProcAddress(libpsapi, "EmptyWorkingSet")
	enumDeviceDrivers = doGetProcAddress(libpsapi, "EnumDeviceDrivers")
	enumPageFiles = doGetProcAddress(libpsapi, "EnumPageFilesW")
	enumProcessModules = doGetProcAddress(libpsapi, "EnumProcessModules")
	enumProcessModulesEx = doGetProcAddress(libpsapi, "EnumProcessModulesEx")
	enumProcesses = doGetProcAddress(libpsapi, "EnumProcesses")
	getDeviceDriverBaseName = doGetProcAddress(libpsapi, "GetDeviceDriverBaseNameW")
	getDeviceDriverFileName = doGetProcAddress(libpsapi, "GetDeviceDriverFileNameW")
	getMappedFileName = doGetProcAddress(libpsapi, "GetMappedFileNameW")
	getModuleBaseName = doGetProcAddress(libpsapi, "GetModuleBaseNameW")
	getModuleFileNameEx = doGetProcAddress(libpsapi, "GetModuleFileNameExW")
	getModuleInformation = doGetProcAddress(libpsapi, "GetModuleInformation")
	getPerformanceInfo = doGetProcAddress(libpsapi, "GetPerformanceInfo")
	getProcessImageFileName = doGetProcAddress(libpsapi, "GetProcessImageFileNameW")
	getProcessMemoryInfo = doGetProcAddress(libpsapi, "GetProcessMemoryInfo")
	getWsChanges = doGetProcAddress(libpsapi, "GetWsChanges")
	getWsChangesEx = doGetProcAddress(libpsapi, "GetWsChangesEx")
	initializeProcessForWsWatch = doGetProcAddress(libpsapi, "InitializeProcessForWsWatch")
	queryWorkingSet = doGetProcAddress(libpsapi, "QueryWorkingSet")
	queryWorkingSetEx = doGetProcAddress(libpsapi, "QueryWorkingSetEx")
}

func EmptyWorkingSet(hProcess HANDLE) bool {
	ret1 := syscall3(emptyWorkingSet, 1,
		uintptr(hProcess),
		0,
		0)
	return ret1 != 0
}

func EnumDeviceDrivers(lpImageBase *LPVOID, cb uint32, lpcbNeeded *uint32) bool {
	ret1 := syscall3(enumDeviceDrivers, 3,
		uintptr(unsafe.Pointer(lpImageBase)),
		uintptr(cb),
		uintptr(unsafe.Pointer(lpcbNeeded)))
	return ret1 != 0
}

func EnumPageFiles(pCallBackRoutine PENUM_PAGE_FILE_CALLBACK, pContext LPVOID) bool {
	pCallBackRoutineCallback := syscall.NewCallback(func(pContextRawArg LPVOID, pPageFileInfoRawArg PENUM_PAGE_FILE_INFORMATION, lpFilenameRawArg /*const*/ *uint16) uintptr {
		lpFilename := stringFromUnicode16(lpFilenameRawArg)
		ret := pCallBackRoutine(pContextRawArg, pPageFileInfoRawArg, lpFilename)
		return uintptr(ret)
	})
	ret1 := syscall3(enumPageFiles, 2,
		pCallBackRoutineCallback,
		uintptr(unsafe.Pointer(pContext)),
		0)
	return ret1 != 0
}

func EnumProcessModules(hProcess HANDLE, lphModule *HMODULE, cb uint32, lpcbNeeded *uint32) bool {
	ret1 := syscall6(enumProcessModules, 4,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lphModule)),
		uintptr(cb),
		uintptr(unsafe.Pointer(lpcbNeeded)),
		0,
		0)
	return ret1 != 0
}

func EnumProcessModulesEx(hProcess HANDLE, lphModule *HMODULE, cb uint32, lpcbNeeded *uint32, dwFilterFlag uint32) bool {
	ret1 := syscall6(enumProcessModulesEx, 5,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lphModule)),
		uintptr(cb),
		uintptr(unsafe.Pointer(lpcbNeeded)),
		uintptr(dwFilterFlag),
		0)
	return ret1 != 0
}

func EnumProcesses(lpidProcess *uint32, cb uint32, cbNeeded *uint32) bool {
	ret1 := syscall3(enumProcesses, 3,
		uintptr(unsafe.Pointer(lpidProcess)),
		uintptr(cb),
		uintptr(unsafe.Pointer(cbNeeded)))
	return ret1 != 0
}

func GetDeviceDriverBaseName(imageBase LPVOID, lpBaseName LPWSTR, nSize uint32) uint32 {
	ret1 := syscall3(getDeviceDriverBaseName, 3,
		uintptr(unsafe.Pointer(imageBase)),
		uintptr(unsafe.Pointer(lpBaseName)),
		uintptr(nSize))
	return uint32(ret1)
}

func GetDeviceDriverFileName(imageBase LPVOID, lpFilename LPWSTR, nSize uint32) uint32 {
	ret1 := syscall3(getDeviceDriverFileName, 3,
		uintptr(unsafe.Pointer(imageBase)),
		uintptr(unsafe.Pointer(lpFilename)),
		uintptr(nSize))
	return uint32(ret1)
}

func GetMappedFileName(hProcess HANDLE, lpv LPVOID, lpFilename LPWSTR, nSize uint32) uint32 {
	ret1 := syscall6(getMappedFileName, 4,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpv)),
		uintptr(unsafe.Pointer(lpFilename)),
		uintptr(nSize),
		0,
		0)
	return uint32(ret1)
}

func GetModuleBaseName(hProcess HANDLE, hModule HMODULE, lpBaseName LPWSTR, nSize uint32) uint32 {
	ret1 := syscall6(getModuleBaseName, 4,
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(unsafe.Pointer(lpBaseName)),
		uintptr(nSize),
		0,
		0)
	return uint32(ret1)
}

func GetModuleFileNameEx(hProcess HANDLE, hModule HMODULE, lpFilename LPWSTR, nSize uint32) uint32 {
	ret1 := syscall6(getModuleFileNameEx, 4,
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(unsafe.Pointer(lpFilename)),
		uintptr(nSize),
		0,
		0)
	return uint32(ret1)
}

func GetModuleInformation(hProcess HANDLE, hModule HMODULE, lpmodinfo *MODULEINFO, cb uint32) bool {
	ret1 := syscall6(getModuleInformation, 4,
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(unsafe.Pointer(lpmodinfo)),
		uintptr(cb),
		0,
		0)
	return ret1 != 0
}

func GetPerformanceInfo(pPerformanceInformation PPERFORMACE_INFORMATION, cb uint32) bool {
	ret1 := syscall3(getPerformanceInfo, 2,
		uintptr(unsafe.Pointer(pPerformanceInformation)),
		uintptr(cb),
		0)
	return ret1 != 0
}

func GetProcessImageFileName(hProcess HANDLE, lpImageFileName LPWSTR, nSize uint32) uint32 {
	ret1 := syscall3(getProcessImageFileName, 3,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpImageFileName)),
		uintptr(nSize))
	return uint32(ret1)
}

func GetProcessMemoryInfo(process HANDLE, ppsmemCounters PPROCESS_MEMORY_COUNTERS, cb uint32) bool {
	ret1 := syscall3(getProcessMemoryInfo, 3,
		uintptr(process),
		uintptr(unsafe.Pointer(ppsmemCounters)),
		uintptr(cb))
	return ret1 != 0
}

func GetWsChanges(hProcess HANDLE, lpWatchInfo PPSAPI_WS_WATCH_INFORMATION, cb uint32) bool {
	ret1 := syscall3(getWsChanges, 3,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpWatchInfo)),
		uintptr(cb))
	return ret1 != 0
}

func GetWsChangesEx(hProcess HANDLE, lpWatchInfoEx PPSAPI_WS_WATCH_INFORMATION_EX, cb uint32) bool {
	ret1 := syscall3(getWsChangesEx, 3,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpWatchInfoEx)),
		uintptr(cb))
	return ret1 != 0
}

func InitializeProcessForWsWatch(hProcess HANDLE) bool {
	ret1 := syscall3(initializeProcessForWsWatch, 1,
		uintptr(hProcess),
		0,
		0)
	return ret1 != 0
}

func QueryWorkingSet(hProcess HANDLE, pv uintptr, cb uint32) bool {
	ret1 := syscall3(queryWorkingSet, 3,
		uintptr(hProcess),
		pv,
		uintptr(cb))
	return ret1 != 0
}

func QueryWorkingSetEx(hProcess HANDLE, pv uintptr, cb uint32) bool {
	ret1 := syscall3(queryWorkingSetEx, 3,
		uintptr(hProcess),
		pv,
		uintptr(cb))
	return ret1 != 0
}
