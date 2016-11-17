// +build windows

package win

import (
	"unsafe"
)

var (
	// Library
	libwinmm uintptr

	// Functions
	closeDriver                 uintptr
	defDriverProc               uintptr
	drvGetModuleHandle          uintptr
	getDriverModuleHandle       uintptr
	openDriver                  uintptr
	playSound                   uintptr
	sendDriverMessage           uintptr
	auxGetDevCaps               uintptr
	auxGetNumDevs               uintptr
	auxGetVolume                uintptr
	auxOutMessage               uintptr
	auxSetVolume                uintptr
	joyGetNumDevs               uintptr
	joyGetThreshold             uintptr
	joyReleaseCapture           uintptr
	joySetCapture               uintptr
	joySetThreshold             uintptr
	mciGetCreatorTask           uintptr
	mciGetDeviceIDFromElementID uintptr
	mciGetDeviceID              uintptr
	mciGetErrorString           uintptr
	mciSendCommand              uintptr
	mciSendString               uintptr
	midiConnect                 uintptr
	midiDisconnect              uintptr
	midiInAddBuffer             uintptr
	midiInClose                 uintptr
	midiInGetDevCaps            uintptr
	midiInGetErrorText          uintptr
	midiInGetID                 uintptr
	midiInGetNumDevs            uintptr
	midiInMessage               uintptr
	midiInOpen                  uintptr
	midiInPrepareHeader         uintptr
	midiInReset                 uintptr
	midiInStart                 uintptr
	midiInStop                  uintptr
	midiInUnprepareHeader       uintptr
	midiOutCacheDrumPatches     uintptr
	midiOutCachePatches         uintptr
	midiOutClose                uintptr
	midiOutGetErrorText         uintptr
	midiOutGetID                uintptr
	midiOutGetNumDevs           uintptr
	midiOutGetVolume            uintptr
	midiOutLongMsg              uintptr
	midiOutMessage              uintptr
	midiOutOpen                 uintptr
	midiOutPrepareHeader        uintptr
	midiOutReset                uintptr
	midiOutSetVolume            uintptr
	midiOutShortMsg             uintptr
	midiOutUnprepareHeader      uintptr
	midiStreamClose             uintptr
	midiStreamOpen              uintptr
	midiStreamOut               uintptr
	midiStreamPause             uintptr
	midiStreamProperty          uintptr
	midiStreamRestart           uintptr
	midiStreamStop              uintptr
	mixerGetNumDevs             uintptr
	sndPlaySound                uintptr
	timeBeginPeriod             uintptr
	timeEndPeriod               uintptr
	timeGetTime                 uintptr
	timeKillEvent               uintptr
	waveInGetErrorText          uintptr
	waveInGetNumDevs            uintptr
	waveOutGetErrorText         uintptr
	waveOutGetNumDevs           uintptr
	driverCallback              uintptr
	joyConfigChanged            uintptr
	mciDriverNotify             uintptr
	mciDriverYield              uintptr
	mciExecute                  uintptr
	mciFreeCommandResource      uintptr
	mciGetDriverData            uintptr
	mciLoadCommandResource      uintptr
	mciSetDriverData            uintptr
	mmGetCurrentTask            uintptr
	mmTaskBlock                 uintptr
	mmTaskSignal                uintptr
	mmTaskYield                 uintptr
	mmsystemGetVersion          uintptr
)

func init() {
	// Library
	libwinmm = doLoadLibrary("winmm.dll")

	// Functions
	closeDriver = doGetProcAddress(libwinmm, "CloseDriver")
	defDriverProc = doGetProcAddress(libwinmm, "DefDriverProc")
	drvGetModuleHandle = doGetProcAddress(libwinmm, "DrvGetModuleHandle")
	getDriverModuleHandle = doGetProcAddress(libwinmm, "GetDriverModuleHandle")
	openDriver = doGetProcAddress(libwinmm, "OpenDriver")
	playSound = doGetProcAddress(libwinmm, "PlaySoundW")
	sendDriverMessage = doGetProcAddress(libwinmm, "SendDriverMessage")
	auxGetDevCaps = doGetProcAddress(libwinmm, "auxGetDevCapsW")
	auxGetNumDevs = doGetProcAddress(libwinmm, "auxGetNumDevs")
	auxGetVolume = doGetProcAddress(libwinmm, "auxGetVolume")
	auxOutMessage = doGetProcAddress(libwinmm, "auxOutMessage")
	auxSetVolume = doGetProcAddress(libwinmm, "auxSetVolume")
	joyGetNumDevs = doGetProcAddress(libwinmm, "joyGetNumDevs")
	joyGetThreshold = doGetProcAddress(libwinmm, "joyGetThreshold")
	joyReleaseCapture = doGetProcAddress(libwinmm, "joyReleaseCapture")
	joySetCapture = doGetProcAddress(libwinmm, "joySetCapture")
	joySetThreshold = doGetProcAddress(libwinmm, "joySetThreshold")
	mciGetCreatorTask = doGetProcAddress(libwinmm, "mciGetCreatorTask")
	mciGetDeviceIDFromElementID = doGetProcAddress(libwinmm, "mciGetDeviceIDFromElementIDW")
	mciGetDeviceID = doGetProcAddress(libwinmm, "mciGetDeviceIDW")
	mciGetErrorString = doGetProcAddress(libwinmm, "mciGetErrorStringW")
	mciSendCommand = doGetProcAddress(libwinmm, "mciSendCommandW")
	mciSendString = doGetProcAddress(libwinmm, "mciSendStringW")
	midiConnect = doGetProcAddress(libwinmm, "midiConnect")
	midiDisconnect = doGetProcAddress(libwinmm, "midiDisconnect")
	midiInAddBuffer = doGetProcAddress(libwinmm, "midiInAddBuffer")
	midiInClose = doGetProcAddress(libwinmm, "midiInClose")
	midiInGetDevCaps = doGetProcAddress(libwinmm, "midiInGetDevCapsW")
	midiInGetErrorText = doGetProcAddress(libwinmm, "midiInGetErrorTextW")
	midiInGetID = doGetProcAddress(libwinmm, "midiInGetID")
	midiInGetNumDevs = doGetProcAddress(libwinmm, "midiInGetNumDevs")
	midiInMessage = doGetProcAddress(libwinmm, "midiInMessage")
	midiInOpen = doGetProcAddress(libwinmm, "midiInOpen")
	midiInPrepareHeader = doGetProcAddress(libwinmm, "midiInPrepareHeader")
	midiInReset = doGetProcAddress(libwinmm, "midiInReset")
	midiInStart = doGetProcAddress(libwinmm, "midiInStart")
	midiInStop = doGetProcAddress(libwinmm, "midiInStop")
	midiInUnprepareHeader = doGetProcAddress(libwinmm, "midiInUnprepareHeader")
	midiOutCacheDrumPatches = doGetProcAddress(libwinmm, "midiOutCacheDrumPatches")
	midiOutCachePatches = doGetProcAddress(libwinmm, "midiOutCachePatches")
	midiOutClose = doGetProcAddress(libwinmm, "midiOutClose")
	midiOutGetErrorText = doGetProcAddress(libwinmm, "midiOutGetErrorTextW")
	midiOutGetID = doGetProcAddress(libwinmm, "midiOutGetID")
	midiOutGetNumDevs = doGetProcAddress(libwinmm, "midiOutGetNumDevs")
	midiOutGetVolume = doGetProcAddress(libwinmm, "midiOutGetVolume")
	midiOutLongMsg = doGetProcAddress(libwinmm, "midiOutLongMsg")
	midiOutMessage = doGetProcAddress(libwinmm, "midiOutMessage")
	midiOutOpen = doGetProcAddress(libwinmm, "midiOutOpen")
	midiOutPrepareHeader = doGetProcAddress(libwinmm, "midiOutPrepareHeader")
	midiOutReset = doGetProcAddress(libwinmm, "midiOutReset")
	midiOutSetVolume = doGetProcAddress(libwinmm, "midiOutSetVolume")
	midiOutShortMsg = doGetProcAddress(libwinmm, "midiOutShortMsg")
	midiOutUnprepareHeader = doGetProcAddress(libwinmm, "midiOutUnprepareHeader")
	midiStreamClose = doGetProcAddress(libwinmm, "midiStreamClose")
	midiStreamOpen = doGetProcAddress(libwinmm, "midiStreamOpen")
	midiStreamOut = doGetProcAddress(libwinmm, "midiStreamOut")
	midiStreamPause = doGetProcAddress(libwinmm, "midiStreamPause")
	midiStreamProperty = doGetProcAddress(libwinmm, "midiStreamProperty")
	midiStreamRestart = doGetProcAddress(libwinmm, "midiStreamRestart")
	midiStreamStop = doGetProcAddress(libwinmm, "midiStreamStop")
	mixerGetNumDevs = doGetProcAddress(libwinmm, "mixerGetNumDevs")
	sndPlaySound = doGetProcAddress(libwinmm, "sndPlaySoundW")
	timeBeginPeriod = doGetProcAddress(libwinmm, "timeBeginPeriod")
	timeEndPeriod = doGetProcAddress(libwinmm, "timeEndPeriod")
	timeGetTime = doGetProcAddress(libwinmm, "timeGetTime")
	timeKillEvent = doGetProcAddress(libwinmm, "timeKillEvent")
	waveInGetErrorText = doGetProcAddress(libwinmm, "waveInGetErrorTextW")
	waveInGetNumDevs = doGetProcAddress(libwinmm, "waveInGetNumDevs")
	waveOutGetErrorText = doGetProcAddress(libwinmm, "waveOutGetErrorTextW")
	waveOutGetNumDevs = doGetProcAddress(libwinmm, "waveOutGetNumDevs")
	driverCallback = doGetProcAddress(libwinmm, "DriverCallback")
	joyConfigChanged = doGetProcAddress(libwinmm, "joyConfigChanged")
	mciDriverNotify = doGetProcAddress(libwinmm, "mciDriverNotify")
	mciDriverYield = doGetProcAddress(libwinmm, "mciDriverYield")
	mciExecute = doGetProcAddress(libwinmm, "mciExecute")
	mciFreeCommandResource = doGetProcAddress(libwinmm, "mciFreeCommandResource")
	mciGetDriverData = doGetProcAddress(libwinmm, "mciGetDriverData")
	mciLoadCommandResource = doGetProcAddress(libwinmm, "mciLoadCommandResource")
	mciSetDriverData = doGetProcAddress(libwinmm, "mciSetDriverData")
	mmGetCurrentTask = doGetProcAddress(libwinmm, "mmGetCurrentTask")
	mmTaskBlock = doGetProcAddress(libwinmm, "mmTaskBlock")
	mmTaskSignal = doGetProcAddress(libwinmm, "mmTaskSignal")
	mmTaskYield = doGetProcAddress(libwinmm, "mmTaskYield")
	mmsystemGetVersion = doGetProcAddress(libwinmm, "mmsystemGetVersion")
}

func CloseDriver(hDriver HDRVR, lParam1 LPARAM, lParam2 LPARAM) LRESULT {
	ret1 := syscall3(closeDriver, 3,
		uintptr(hDriver),
		uintptr(lParam1),
		uintptr(lParam2))
	return LRESULT(ret1)
}

func DefDriverProc(dwDriverIdentifier *uint32, hdrvr HDRVR, uMsg uint32, lParam1 LPARAM, lParam2 LPARAM) LRESULT {
	ret1 := syscall6(defDriverProc, 5,
		uintptr(unsafe.Pointer(dwDriverIdentifier)),
		uintptr(hdrvr),
		uintptr(uMsg),
		uintptr(lParam1),
		uintptr(lParam2),
		0)
	return LRESULT(ret1)
}

func DrvGetModuleHandle(hDriver HDRVR) HMODULE {
	ret1 := syscall3(drvGetModuleHandle, 1,
		uintptr(hDriver),
		0,
		0)
	return HMODULE(ret1)
}

func GetDriverModuleHandle(hDriver HDRVR) HMODULE {
	ret1 := syscall3(getDriverModuleHandle, 1,
		uintptr(hDriver),
		0,
		0)
	return HMODULE(ret1)
}

func OpenDriver(szDriverName string, szSectionName string, lParam2 LPARAM) HDRVR {
	szDriverNameStr := unicode16FromString(szDriverName)
	szSectionNameStr := unicode16FromString(szSectionName)
	ret1 := syscall3(openDriver, 3,
		uintptr(unsafe.Pointer(&szDriverNameStr[0])),
		uintptr(unsafe.Pointer(&szSectionNameStr[0])),
		uintptr(lParam2))
	return HDRVR(ret1)
}

func PlaySound(pszSound string, hmod HMODULE, fdwSound uint32) bool {
	pszSoundStr := unicode16FromString(pszSound)
	ret1 := syscall3(playSound, 3,
		uintptr(unsafe.Pointer(&pszSoundStr[0])),
		uintptr(hmod),
		uintptr(fdwSound))
	return ret1 != 0
}

func SendDriverMessage(hDriver HDRVR, message uint32, lParam1 LPARAM, lParam2 LPARAM) LRESULT {
	ret1 := syscall6(sendDriverMessage, 4,
		uintptr(hDriver),
		uintptr(message),
		uintptr(lParam1),
		uintptr(lParam2),
		0,
		0)
	return LRESULT(ret1)
}

func AuxGetDevCaps(uDeviceID *uint32, pac *AUXCAPS, cbac uint32) MMRESULT {
	ret1 := syscall3(auxGetDevCaps, 3,
		uintptr(unsafe.Pointer(uDeviceID)),
		uintptr(unsafe.Pointer(pac)),
		uintptr(cbac))
	return MMRESULT(ret1)
}

func AuxGetNumDevs() uint32 {
	ret1 := syscall3(auxGetNumDevs, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func AuxGetVolume(uDeviceID uint32, pdwVolume *uint32) MMRESULT {
	ret1 := syscall3(auxGetVolume, 2,
		uintptr(uDeviceID),
		uintptr(unsafe.Pointer(pdwVolume)),
		0)
	return MMRESULT(ret1)
}

func AuxOutMessage(uDeviceID uint32, uMsg uint32, dw1 *uint32, dw2 *uint32) MMRESULT {
	ret1 := syscall6(auxOutMessage, 4,
		uintptr(uDeviceID),
		uintptr(uMsg),
		uintptr(unsafe.Pointer(dw1)),
		uintptr(unsafe.Pointer(dw2)),
		0,
		0)
	return MMRESULT(ret1)
}

func AuxSetVolume(uDeviceID uint32, dwVolume uint32) MMRESULT {
	ret1 := syscall3(auxSetVolume, 2,
		uintptr(uDeviceID),
		uintptr(dwVolume),
		0)
	return MMRESULT(ret1)
}

// TODO: Unknown type(s): LPJOYCAPSW
// func JoyGetDevCaps(uJoyID *uint32, pjc LPJOYCAPSW, cbjc uint32) MMRESULT

func JoyGetNumDevs() uint32 {
	ret1 := syscall3(joyGetNumDevs, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): LPJOYINFO
// func JoyGetPos(uJoyID uint32, pji LPJOYINFO) MMRESULT

// TODO: Unknown type(s): LPJOYINFOEX
// func JoyGetPosEx(uJoyID uint32, pji LPJOYINFOEX) MMRESULT

func JoyGetThreshold(uJoyID uint32, puThreshold *UINT) MMRESULT {
	ret1 := syscall3(joyGetThreshold, 2,
		uintptr(uJoyID),
		uintptr(unsafe.Pointer(puThreshold)),
		0)
	return MMRESULT(ret1)
}

func JoyReleaseCapture(uJoyID uint32) MMRESULT {
	ret1 := syscall3(joyReleaseCapture, 1,
		uintptr(uJoyID),
		0,
		0)
	return MMRESULT(ret1)
}

func JoySetCapture(hwnd HWND, uJoyID uint32, uPeriod uint32, fChanged bool) MMRESULT {
	ret1 := syscall6(joySetCapture, 4,
		uintptr(hwnd),
		uintptr(uJoyID),
		uintptr(uPeriod),
		getUintptrFromBool(fChanged),
		0,
		0)
	return MMRESULT(ret1)
}

func JoySetThreshold(uJoyID uint32, uThreshold uint32) MMRESULT {
	ret1 := syscall3(joySetThreshold, 2,
		uintptr(uJoyID),
		uintptr(uThreshold),
		0)
	return MMRESULT(ret1)
}

func MciGetCreatorTask(mciId MCIDEVICEID) HTASK {
	ret1 := syscall3(mciGetCreatorTask, 1,
		uintptr(mciId),
		0,
		0)
	return HTASK(ret1)
}

func MciGetDeviceIDFromElementID(dwElementID uint32, lpstrType string) MCIDEVICEID {
	lpstrTypeStr := unicode16FromString(lpstrType)
	ret1 := syscall3(mciGetDeviceIDFromElementID, 2,
		uintptr(dwElementID),
		uintptr(unsafe.Pointer(&lpstrTypeStr[0])),
		0)
	return MCIDEVICEID(ret1)
}

func MciGetDeviceID(pszDevice string) MCIDEVICEID {
	pszDeviceStr := unicode16FromString(pszDevice)
	ret1 := syscall3(mciGetDeviceID, 1,
		uintptr(unsafe.Pointer(&pszDeviceStr[0])),
		0,
		0)
	return MCIDEVICEID(ret1)
}

func MciGetErrorString(mcierr MCIERROR, pszText LPWSTR, cchText uint32) bool {
	ret1 := syscall3(mciGetErrorString, 3,
		uintptr(mcierr),
		uintptr(unsafe.Pointer(pszText)),
		uintptr(cchText))
	return ret1 != 0
}

// TODO: Unknown type(s): YIELDPROC
// func MciGetYieldProc(mciId MCIDEVICEID, pdwYieldData *uint32) YIELDPROC

func MciSendCommand(mciId MCIDEVICEID, uMsg uint32, dwParam1 *uint32, dwParam2 *uint32) MCIERROR {
	ret1 := syscall6(mciSendCommand, 4,
		uintptr(mciId),
		uintptr(uMsg),
		uintptr(unsafe.Pointer(dwParam1)),
		uintptr(unsafe.Pointer(dwParam2)),
		0,
		0)
	return MCIERROR(ret1)
}

func MciSendString(lpstrCommand string, lpstrReturnString LPWSTR, uReturnLength uint32, hwndCallback HWND) MCIERROR {
	lpstrCommandStr := unicode16FromString(lpstrCommand)
	ret1 := syscall6(mciSendString, 4,
		uintptr(unsafe.Pointer(&lpstrCommandStr[0])),
		uintptr(unsafe.Pointer(lpstrReturnString)),
		uintptr(uReturnLength),
		uintptr(hwndCallback),
		0,
		0)
	return MCIERROR(ret1)
}

// TODO: Unknown type(s): YIELDPROC
// func MciSetYieldProc(mciId MCIDEVICEID, fpYieldProc YIELDPROC, dwYieldData uint32) bool

func MidiConnect(hmi HMIDI, hmo HMIDIOUT, pReserved LPVOID) MMRESULT {
	ret1 := syscall3(midiConnect, 3,
		uintptr(hmi),
		uintptr(hmo),
		uintptr(unsafe.Pointer(pReserved)))
	return MMRESULT(ret1)
}

func MidiDisconnect(hmi HMIDI, hmo HMIDIOUT, pReserved LPVOID) MMRESULT {
	ret1 := syscall3(midiDisconnect, 3,
		uintptr(hmi),
		uintptr(hmo),
		uintptr(unsafe.Pointer(pReserved)))
	return MMRESULT(ret1)
}

func MidiInAddBuffer(hmi HMIDIIN, pmh *MIDIHDR, cbmh uint32) MMRESULT {
	ret1 := syscall3(midiInAddBuffer, 3,
		uintptr(hmi),
		uintptr(unsafe.Pointer(pmh)),
		uintptr(cbmh))
	return MMRESULT(ret1)
}

func MidiInClose(hmi HMIDIIN) MMRESULT {
	ret1 := syscall3(midiInClose, 1,
		uintptr(hmi),
		0,
		0)
	return MMRESULT(ret1)
}

func MidiInGetDevCaps(uDeviceID *uint32, pmic *MIDIINCAPS, cbmic uint32) MMRESULT {
	ret1 := syscall3(midiInGetDevCaps, 3,
		uintptr(unsafe.Pointer(uDeviceID)),
		uintptr(unsafe.Pointer(pmic)),
		uintptr(cbmic))
	return MMRESULT(ret1)
}

func MidiInGetErrorText(mmrError MMRESULT, pszText LPWSTR, cchText uint32) MMRESULT {
	ret1 := syscall3(midiInGetErrorText, 3,
		uintptr(mmrError),
		uintptr(unsafe.Pointer(pszText)),
		uintptr(cchText))
	return MMRESULT(ret1)
}

func MidiInGetID(hmi HMIDIIN, puDeviceID *UINT) MMRESULT {
	ret1 := syscall3(midiInGetID, 2,
		uintptr(hmi),
		uintptr(unsafe.Pointer(puDeviceID)),
		0)
	return MMRESULT(ret1)
}

func MidiInGetNumDevs() uint32 {
	ret1 := syscall3(midiInGetNumDevs, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func MidiInMessage(hmi HMIDIIN, uMsg uint32, dw1 *uint32, dw2 *uint32) MMRESULT {
	ret1 := syscall6(midiInMessage, 4,
		uintptr(hmi),
		uintptr(uMsg),
		uintptr(unsafe.Pointer(dw1)),
		uintptr(unsafe.Pointer(dw2)),
		0,
		0)
	return MMRESULT(ret1)
}

func MidiInOpen(phmi *HMIDIIN, uDeviceID uint32, dwCallback *uint32, dwInstance *uint32, fdwOpen uint32) MMRESULT {
	ret1 := syscall6(midiInOpen, 5,
		uintptr(unsafe.Pointer(phmi)),
		uintptr(uDeviceID),
		uintptr(unsafe.Pointer(dwCallback)),
		uintptr(unsafe.Pointer(dwInstance)),
		uintptr(fdwOpen),
		0)
	return MMRESULT(ret1)
}

func MidiInPrepareHeader(hmi HMIDIIN, pmh *MIDIHDR, cbmh uint32) MMRESULT {
	ret1 := syscall3(midiInPrepareHeader, 3,
		uintptr(hmi),
		uintptr(unsafe.Pointer(pmh)),
		uintptr(cbmh))
	return MMRESULT(ret1)
}

func MidiInReset(hmi HMIDIIN) MMRESULT {
	ret1 := syscall3(midiInReset, 1,
		uintptr(hmi),
		0,
		0)
	return MMRESULT(ret1)
}

func MidiInStart(hmi HMIDIIN) MMRESULT {
	ret1 := syscall3(midiInStart, 1,
		uintptr(hmi),
		0,
		0)
	return MMRESULT(ret1)
}

func MidiInStop(hmi HMIDIIN) MMRESULT {
	ret1 := syscall3(midiInStop, 1,
		uintptr(hmi),
		0,
		0)
	return MMRESULT(ret1)
}

func MidiInUnprepareHeader(hmi HMIDIIN, pmh *MIDIHDR, cbmh uint32) MMRESULT {
	ret1 := syscall3(midiInUnprepareHeader, 3,
		uintptr(hmi),
		uintptr(unsafe.Pointer(pmh)),
		uintptr(cbmh))
	return MMRESULT(ret1)
}

func MidiOutCacheDrumPatches(hmo HMIDIOUT, uPatch uint32, pwkya *uint16, fuCache uint32) MMRESULT {
	ret1 := syscall6(midiOutCacheDrumPatches, 4,
		uintptr(hmo),
		uintptr(uPatch),
		uintptr(unsafe.Pointer(pwkya)),
		uintptr(fuCache),
		0,
		0)
	return MMRESULT(ret1)
}

func MidiOutCachePatches(hmo HMIDIOUT, uBank uint32, pwpa *uint16, fuCache uint32) MMRESULT {
	ret1 := syscall6(midiOutCachePatches, 4,
		uintptr(hmo),
		uintptr(uBank),
		uintptr(unsafe.Pointer(pwpa)),
		uintptr(fuCache),
		0,
		0)
	return MMRESULT(ret1)
}

func MidiOutClose(hmo HMIDIOUT) MMRESULT {
	ret1 := syscall3(midiOutClose, 1,
		uintptr(hmo),
		0,
		0)
	return MMRESULT(ret1)
}

// TODO: Unknown type(s): LPMIDIOUTCAPSW
// func MidiOutGetDevCaps(uDeviceID *uint32, pmoc LPMIDIOUTCAPSW, cbmoc uint32) MMRESULT

func MidiOutGetErrorText(mmrError MMRESULT, pszText LPWSTR, cchText uint32) MMRESULT {
	ret1 := syscall3(midiOutGetErrorText, 3,
		uintptr(mmrError),
		uintptr(unsafe.Pointer(pszText)),
		uintptr(cchText))
	return MMRESULT(ret1)
}

func MidiOutGetID(hmo HMIDIOUT, puDeviceID *UINT) MMRESULT {
	ret1 := syscall3(midiOutGetID, 2,
		uintptr(hmo),
		uintptr(unsafe.Pointer(puDeviceID)),
		0)
	return MMRESULT(ret1)
}

func MidiOutGetNumDevs() uint32 {
	ret1 := syscall3(midiOutGetNumDevs, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func MidiOutGetVolume(hmo HMIDIOUT, pdwVolume *uint32) MMRESULT {
	ret1 := syscall3(midiOutGetVolume, 2,
		uintptr(hmo),
		uintptr(unsafe.Pointer(pdwVolume)),
		0)
	return MMRESULT(ret1)
}

func MidiOutLongMsg(hmo HMIDIOUT, pmh *MIDIHDR, cbmh uint32) MMRESULT {
	ret1 := syscall3(midiOutLongMsg, 3,
		uintptr(hmo),
		uintptr(unsafe.Pointer(pmh)),
		uintptr(cbmh))
	return MMRESULT(ret1)
}

func MidiOutMessage(hmo HMIDIOUT, uMsg uint32, dw1 *uint32, dw2 *uint32) MMRESULT {
	ret1 := syscall6(midiOutMessage, 4,
		uintptr(hmo),
		uintptr(uMsg),
		uintptr(unsafe.Pointer(dw1)),
		uintptr(unsafe.Pointer(dw2)),
		0,
		0)
	return MMRESULT(ret1)
}

func MidiOutOpen(phmo *HMIDIOUT, uDeviceID uint32, dwCallback *uint32, dwInstance *uint32, fdwOpen uint32) MMRESULT {
	ret1 := syscall6(midiOutOpen, 5,
		uintptr(unsafe.Pointer(phmo)),
		uintptr(uDeviceID),
		uintptr(unsafe.Pointer(dwCallback)),
		uintptr(unsafe.Pointer(dwInstance)),
		uintptr(fdwOpen),
		0)
	return MMRESULT(ret1)
}

func MidiOutPrepareHeader(hmo HMIDIOUT, pmh *MIDIHDR, cbmh uint32) MMRESULT {
	ret1 := syscall3(midiOutPrepareHeader, 3,
		uintptr(hmo),
		uintptr(unsafe.Pointer(pmh)),
		uintptr(cbmh))
	return MMRESULT(ret1)
}

func MidiOutReset(hmo HMIDIOUT) MMRESULT {
	ret1 := syscall3(midiOutReset, 1,
		uintptr(hmo),
		0,
		0)
	return MMRESULT(ret1)
}

func MidiOutSetVolume(hmo HMIDIOUT, dwVolume uint32) MMRESULT {
	ret1 := syscall3(midiOutSetVolume, 2,
		uintptr(hmo),
		uintptr(dwVolume),
		0)
	return MMRESULT(ret1)
}

func MidiOutShortMsg(hmo HMIDIOUT, dwMsg uint32) MMRESULT {
	ret1 := syscall3(midiOutShortMsg, 2,
		uintptr(hmo),
		uintptr(dwMsg),
		0)
	return MMRESULT(ret1)
}

func MidiOutUnprepareHeader(hmo HMIDIOUT, pmh *MIDIHDR, cbmh uint32) MMRESULT {
	ret1 := syscall3(midiOutUnprepareHeader, 3,
		uintptr(hmo),
		uintptr(unsafe.Pointer(pmh)),
		uintptr(cbmh))
	return MMRESULT(ret1)
}

func MidiStreamClose(hms HMIDISTRM) MMRESULT {
	ret1 := syscall3(midiStreamClose, 1,
		uintptr(hms),
		0,
		0)
	return MMRESULT(ret1)
}

func MidiStreamOpen(phms *HMIDISTRM, puDeviceID *UINT, cMidi uint32, dwCallback *uint32, dwInstance *uint32, fdwOpen uint32) MMRESULT {
	ret1 := syscall6(midiStreamOpen, 6,
		uintptr(unsafe.Pointer(phms)),
		uintptr(unsafe.Pointer(puDeviceID)),
		uintptr(cMidi),
		uintptr(unsafe.Pointer(dwCallback)),
		uintptr(unsafe.Pointer(dwInstance)),
		uintptr(fdwOpen))
	return MMRESULT(ret1)
}

func MidiStreamOut(hms HMIDISTRM, pmh *MIDIHDR, cbmh uint32) MMRESULT {
	ret1 := syscall3(midiStreamOut, 3,
		uintptr(hms),
		uintptr(unsafe.Pointer(pmh)),
		uintptr(cbmh))
	return MMRESULT(ret1)
}

func MidiStreamPause(hms HMIDISTRM) MMRESULT {
	ret1 := syscall3(midiStreamPause, 1,
		uintptr(hms),
		0,
		0)
	return MMRESULT(ret1)
}

// TODO: Unknown type(s): LPMMTIME
// func MidiStreamPosition(hms HMIDISTRM, lpmmt LPMMTIME, cbmmt uint32) MMRESULT

func MidiStreamProperty(hms HMIDISTRM, lppropdata *byte, dwProperty uint32) MMRESULT {
	ret1 := syscall3(midiStreamProperty, 3,
		uintptr(hms),
		uintptr(unsafe.Pointer(lppropdata)),
		uintptr(dwProperty))
	return MMRESULT(ret1)
}

func MidiStreamRestart(hms HMIDISTRM) MMRESULT {
	ret1 := syscall3(midiStreamRestart, 1,
		uintptr(hms),
		0,
		0)
	return MMRESULT(ret1)
}

func MidiStreamStop(hms HMIDISTRM) MMRESULT {
	ret1 := syscall3(midiStreamStop, 1,
		uintptr(hms),
		0,
		0)
	return MMRESULT(ret1)
}

// TODO: Unknown type(s): HMIXER
// func MixerClose(hmx HMIXER) MMRESULT

// TODO: Unknown type(s): HMIXEROBJ, LPMIXERCONTROLDETAILS
// func MixerGetControlDetails(hmxobj HMIXEROBJ, pmxcd LPMIXERCONTROLDETAILS, fdwDetails uint32) MMRESULT

// TODO: Unknown type(s): LPMIXERCAPSW
// func MixerGetDevCaps(uMxId *uint32, pmxcaps LPMIXERCAPSW, cbmxcaps uint32) MMRESULT

// TODO: Unknown type(s): HMIXEROBJ
// func MixerGetID(hmxobj HMIXEROBJ, puMxId *UINT, fdwId uint32) MMRESULT

// TODO: Unknown type(s): HMIXEROBJ, LPMIXERLINECONTROLSW
// func MixerGetLineControls(hmxobj HMIXEROBJ, pmxlc LPMIXERLINECONTROLSW, fdwControls uint32) MMRESULT

// TODO: Unknown type(s): HMIXEROBJ, LPMIXERLINEW
// func MixerGetLineInfo(hmxobj HMIXEROBJ, pmxl LPMIXERLINEW, fdwInfo uint32) MMRESULT

func MixerGetNumDevs() uint32 {
	ret1 := syscall3(mixerGetNumDevs, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): HMIXER
// func MixerMessage(hmx HMIXER, uMsg uint32, dwParam1 *uint32, dwParam2 *uint32) uint32

// TODO: Unknown type(s): LPHMIXER
// func MixerOpen(phmx LPHMIXER, uMxId uint32, dwCallback *uint32, dwInstance *uint32, fdwOpen uint32) MMRESULT

// TODO: Unknown type(s): HMIXEROBJ, LPMIXERCONTROLDETAILS
// func MixerSetControlDetails(hmxobj HMIXEROBJ, pmxcd LPMIXERCONTROLDETAILS, fdwDetails uint32) MMRESULT

// TODO: Unknown type(s): HMMIO, LPMMIOINFO
// func MmioAdvance(hmmio HMMIO, pmmioinfo LPMMIOINFO, fuAdvance uint32) MMRESULT

// TODO: Unknown type(s): HMMIO, LPMMCKINFO
// func MmioAscend(hmmio HMMIO, pmmcki LPMMCKINFO, fuAscend uint32) MMRESULT

// TODO: Unknown type(s): HMMIO
// func MmioClose(hmmio HMMIO, fuClose uint32) MMRESULT

// TODO: Unknown type(s): HMMIO, LPMMCKINFO
// func MmioCreateChunk(hmmio HMMIO, pmmcki LPMMCKINFO, fuCreate uint32) MMRESULT

// TODO: Unknown type(s): HMMIO, LPMMCKINFO, const MMCKINFO *
// func MmioDescend(hmmio HMMIO, pmmcki LPMMCKINFO, pmmckiParent /*const*/ const MMCKINFO *, fuDescend uint32) MMRESULT

// TODO: Unknown type(s): HMMIO
// func MmioFlush(hmmio HMMIO, fuFlush uint32) MMRESULT

// TODO: Unknown type(s): HMMIO, LPMMIOINFO
// func MmioGetInfo(hmmio HMMIO, pmmioinfo LPMMIOINFO, fuInfo uint32) MMRESULT

// TODO: Unknown type(s): FOURCC, LPMMIOPROC
// func MmioInstallIOProc(fccIOProc FOURCC, pIOProc LPMMIOPROC, dwFlags uint32) LPMMIOPROC

// TODO: Unknown type(s): HMMIO, LPMMIOINFO
// func MmioOpen(pszFileName LPWSTR, pmmioinfo LPMMIOINFO, fdwOpen uint32) HMMIO

// TODO: Unknown type(s): HMMIO, HPSTR
// func MmioRead(hmmio HMMIO, pch HPSTR, cch LONG) LONG

// TODO: Unknown type(s): LPCMMIOINFO
// func MmioRename(pszFileName string, pszNewFileName string, pmmioinfo LPCMMIOINFO, fdwRename uint32) MMRESULT

// TODO: Unknown type(s): HMMIO
// func MmioSeek(hmmio HMMIO, lOffset LONG, iOrigin int32) LONG

// TODO: Unknown type(s): HMMIO
// func MmioSendMessage(hmmio HMMIO, uMsg uint32, lParam1 LPARAM, lParam2 LPARAM) LRESULT

// TODO: Unknown type(s): HMMIO
// func MmioSetBuffer(hmmio HMMIO, pchBuffer LPSTR, cchBuffer LONG, fuBuffer uint32) MMRESULT

// TODO: Unknown type(s): HMMIO, LPCMMIOINFO
// func MmioSetInfo(hmmio HMMIO, pmmioinfo LPCMMIOINFO, fuInfo uint32) MMRESULT

// TODO: Unknown type(s): FOURCC
// func MmioStringToFOURCC(sz string, uFlags uint32) FOURCC

// TODO: Unknown type(s): HMMIO, const char _huge *
// func MmioWrite(hmmio HMMIO, pch /*const*/ const char _huge *, cch LONG) LONG

func SndPlaySound(pszSound string, fuSound uint32) bool {
	pszSoundStr := unicode16FromString(pszSound)
	ret1 := syscall3(sndPlaySound, 2,
		uintptr(unsafe.Pointer(&pszSoundStr[0])),
		uintptr(fuSound),
		0)
	return ret1 != 0
}

func TimeBeginPeriod(uPeriod uint32) MMRESULT {
	ret1 := syscall3(timeBeginPeriod, 1,
		uintptr(uPeriod),
		0,
		0)
	return MMRESULT(ret1)
}

func TimeEndPeriod(uPeriod uint32) MMRESULT {
	ret1 := syscall3(timeEndPeriod, 1,
		uintptr(uPeriod),
		0,
		0)
	return MMRESULT(ret1)
}

// TODO: Unknown type(s): LPTIMECAPS
// func TimeGetDevCaps(ptc LPTIMECAPS, cbtc uint32) MMRESULT

// TODO: Unknown type(s): LPMMTIME
// func TimeGetSystemTime(pmmt LPMMTIME, cbmmt uint32) MMRESULT

func TimeGetTime() uint32 {
	ret1 := syscall3(timeGetTime, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func TimeKillEvent(uTimerID uint32) MMRESULT {
	ret1 := syscall3(timeKillEvent, 1,
		uintptr(uTimerID),
		0,
		0)
	return MMRESULT(ret1)
}

// TODO: Unknown type(s): LPTIMECALLBACK
// func TimeSetEvent(uDelay uint32, uResolution uint32, fptc LPTIMECALLBACK, dwUser *uint32, fuEvent uint32) MMRESULT

// TODO: Unknown type(s): HWAVEIN, LPWAVEHDR
// func WaveInAddBuffer(hwi HWAVEIN, pwh LPWAVEHDR, cbwh uint32) MMRESULT

// TODO: Unknown type(s): HWAVEIN
// func WaveInClose(hwi HWAVEIN) MMRESULT

// TODO: Unknown type(s): LPWAVEINCAPSW
// func WaveInGetDevCaps(uDeviceID *uint32, pwic LPWAVEINCAPSW, cbwic uint32) MMRESULT

func WaveInGetErrorText(mmrError MMRESULT, pszText LPWSTR, cchText uint32) MMRESULT {
	ret1 := syscall3(waveInGetErrorText, 3,
		uintptr(mmrError),
		uintptr(unsafe.Pointer(pszText)),
		uintptr(cchText))
	return MMRESULT(ret1)
}

// TODO: Unknown type(s): HWAVEIN
// func WaveInGetID(hwi HWAVEIN, puDeviceID *UINT) MMRESULT

func WaveInGetNumDevs() uint32 {
	ret1 := syscall3(waveInGetNumDevs, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): HWAVEIN, LPMMTIME
// func WaveInGetPosition(hwi HWAVEIN, pmmt LPMMTIME, cbmmt uint32) MMRESULT

// TODO: Unknown type(s): HWAVEIN
// func WaveInMessage(hwi HWAVEIN, uMsg uint32, dw1 *uint32, dw2 *uint32) MMRESULT

// TODO: Unknown type(s): LPCWAVEFORMATEX, LPHWAVEIN
// func WaveInOpen(phwi LPHWAVEIN, uDeviceID uint32, pwfx LPCWAVEFORMATEX, dwCallback *uint32, dwInstance *uint32, fdwOpen uint32) MMRESULT

// TODO: Unknown type(s): HWAVEIN, LPWAVEHDR
// func WaveInPrepareHeader(hwi HWAVEIN, pwh LPWAVEHDR, cbwh uint32) MMRESULT

// TODO: Unknown type(s): HWAVEIN
// func WaveInReset(hwi HWAVEIN) MMRESULT

// TODO: Unknown type(s): HWAVEIN
// func WaveInStart(hwi HWAVEIN) MMRESULT

// TODO: Unknown type(s): HWAVEIN
// func WaveInStop(hwi HWAVEIN) MMRESULT

// TODO: Unknown type(s): HWAVEIN, LPWAVEHDR
// func WaveInUnprepareHeader(hwi HWAVEIN, pwh LPWAVEHDR, cbwh uint32) MMRESULT

// TODO: Unknown type(s): HWAVEOUT
// func WaveOutBreakLoop(hwo HWAVEOUT) MMRESULT

// TODO: Unknown type(s): HWAVEOUT
// func WaveOutClose(hwo HWAVEOUT) MMRESULT

// TODO: Unknown type(s): LPWAVEOUTCAPSW
// func WaveOutGetDevCaps(uDeviceID *uint32, pwoc LPWAVEOUTCAPSW, cbwoc uint32) MMRESULT

func WaveOutGetErrorText(mmrError MMRESULT, pszText LPWSTR, cchText uint32) MMRESULT {
	ret1 := syscall3(waveOutGetErrorText, 3,
		uintptr(mmrError),
		uintptr(unsafe.Pointer(pszText)),
		uintptr(cchText))
	return MMRESULT(ret1)
}

// TODO: Unknown type(s): HWAVEOUT
// func WaveOutGetID(hwo HWAVEOUT, puDeviceID *UINT) MMRESULT

func WaveOutGetNumDevs() uint32 {
	ret1 := syscall3(waveOutGetNumDevs, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): HWAVEOUT
// func WaveOutGetPitch(hwo HWAVEOUT, pdwPitch *uint32) MMRESULT

// TODO: Unknown type(s): HWAVEOUT
// func WaveOutGetPlaybackRate(hwo HWAVEOUT, pdwRate *uint32) MMRESULT

// TODO: Unknown type(s): HWAVEOUT, LPMMTIME
// func WaveOutGetPosition(hwo HWAVEOUT, pmmt LPMMTIME, cbmmt uint32) MMRESULT

// TODO: Unknown type(s): HWAVEOUT
// func WaveOutGetVolume(hwo HWAVEOUT, pdwVolume *uint32) MMRESULT

// TODO: Unknown type(s): HWAVEOUT
// func WaveOutMessage(hwo HWAVEOUT, uMsg uint32, dw1 *uint32, dw2 *uint32) MMRESULT

// TODO: Unknown type(s): LPCWAVEFORMATEX, LPHWAVEOUT
// func WaveOutOpen(phwo LPHWAVEOUT, uDeviceID uint32, pwfx LPCWAVEFORMATEX, dwCallback *uint32, dwInstance *uint32, fdwOpen uint32) MMRESULT

// TODO: Unknown type(s): HWAVEOUT
// func WaveOutPause(hwo HWAVEOUT) MMRESULT

// TODO: Unknown type(s): HWAVEOUT, LPWAVEHDR
// func WaveOutPrepareHeader(hwo HWAVEOUT, pwh LPWAVEHDR, cbwh uint32) MMRESULT

// TODO: Unknown type(s): HWAVEOUT
// func WaveOutReset(hwo HWAVEOUT) MMRESULT

// TODO: Unknown type(s): HWAVEOUT
// func WaveOutRestart(hwo HWAVEOUT) MMRESULT

// TODO: Unknown type(s): HWAVEOUT
// func WaveOutSetPitch(hwo HWAVEOUT, dwPitch uint32) MMRESULT

// TODO: Unknown type(s): HWAVEOUT
// func WaveOutSetPlaybackRate(hwo HWAVEOUT, dwRate uint32) MMRESULT

// TODO: Unknown type(s): HWAVEOUT
// func WaveOutSetVolume(hwo HWAVEOUT, dwVolume uint32) MMRESULT

// TODO: Unknown type(s): HWAVEOUT, LPWAVEHDR
// func WaveOutUnprepareHeader(hwo HWAVEOUT, pwh LPWAVEHDR, cbwh uint32) MMRESULT

// TODO: Unknown type(s): HWAVEOUT, LPWAVEHDR
// func WaveOutWrite(hwo HWAVEOUT, pwh LPWAVEHDR, cbwh uint32) MMRESULT

func DriverCallback(dwCallBack *uint32, uFlags uint32, hDev HDRVR, wMsg uint32, dwUser *uint32, dwParam1 *uint32, dwParam2 *uint32) bool {
	ret1 := syscall9(driverCallback, 7,
		uintptr(unsafe.Pointer(dwCallBack)),
		uintptr(uFlags),
		uintptr(hDev),
		uintptr(wMsg),
		uintptr(unsafe.Pointer(dwUser)),
		uintptr(unsafe.Pointer(dwParam1)),
		uintptr(unsafe.Pointer(dwParam2)),
		0,
		0)
	return ret1 != 0
}

func JoyConfigChanged(flags uint32) MMRESULT {
	ret1 := syscall3(joyConfigChanged, 1,
		uintptr(flags),
		0,
		0)
	return MMRESULT(ret1)
}

func MciDriverNotify(hWndCallBack HWND, wDevID MCIDEVICEID, wStatus uint32) bool {
	ret1 := syscall3(mciDriverNotify, 3,
		uintptr(hWndCallBack),
		uintptr(wDevID),
		uintptr(wStatus))
	return ret1 != 0
}

func MciDriverYield(uDeviceID MCIDEVICEID) uint32 {
	ret1 := syscall3(mciDriverYield, 1,
		uintptr(uDeviceID),
		0,
		0)
	return uint32(ret1)
}

func MciExecute(lpstrCommand /*const*/ LPCSTR) bool {
	ret1 := syscall3(mciExecute, 1,
		uintptr(unsafe.Pointer(lpstrCommand)),
		0,
		0)
	return ret1 != 0
}

func MciFreeCommandResource(uTable uint32) bool {
	ret1 := syscall3(mciFreeCommandResource, 1,
		uintptr(uTable),
		0,
		0)
	return ret1 != 0
}

func MciGetDriverData(uDeviceID MCIDEVICEID) *uint32 {
	ret1 := syscall3(mciGetDriverData, 1,
		uintptr(uDeviceID),
		0,
		0)
	return (*uint32)(unsafe.Pointer(ret1))
}

func MciLoadCommandResource(hInst HINSTANCE, resNameW string, aType uint32) uint32 {
	resNameWStr := unicode16FromString(resNameW)
	ret1 := syscall3(mciLoadCommandResource, 3,
		uintptr(hInst),
		uintptr(unsafe.Pointer(&resNameWStr[0])),
		uintptr(aType))
	return uint32(ret1)
}

func MciSetDriverData(uDeviceID MCIDEVICEID, data *uint32) bool {
	ret1 := syscall3(mciSetDriverData, 2,
		uintptr(uDeviceID),
		uintptr(unsafe.Pointer(data)),
		0)
	return ret1 != 0
}

func MmGetCurrentTask() uint32 {
	ret1 := syscall3(mmGetCurrentTask, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func MmTaskBlock(tid uint32) {
	syscall3(mmTaskBlock, 1,
		uintptr(tid),
		0,
		0)
}

// TODO: Unknown type(s): LPTASKCALLBACK
// func MmTaskCreate(cb LPTASKCALLBACK, ph *HANDLE, client *uint32) uint32

func MmTaskSignal(tid uint32) bool {
	ret1 := syscall3(mmTaskSignal, 1,
		uintptr(tid),
		0,
		0)
	return ret1 != 0
}

func MmTaskYield() {
	syscall3(mmTaskYield, 0,
		0,
		0,
		0)
}

func MmsystemGetVersion() uint32 {
	ret1 := syscall3(mmsystemGetVersion, 0,
		0,
		0,
		0)
	return uint32(ret1)
}
