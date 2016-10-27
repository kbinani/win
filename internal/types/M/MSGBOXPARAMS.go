package win

//ref HWND
//ref HINSTANCE
//ref MSGBOXCALLBACK
type MSGBOXPARAMS struct {
	CbSize             uint32
	HwndOwner          HWND
	HInstance          HINSTANCE
	LpszText           *uint16 // LPCWSTR
	LpszCaption        *uint16 // LPCWSTR
	DwStyle            uint32
	LpszIcon           *uint16 // LPCWSTR
	DwContextHelpId    *uint32 // DWORD_PTR
	LpfnMsgBoxCallback uintptr // MSGBOXCALLBACK
	DwLanguageId       uint32
}
