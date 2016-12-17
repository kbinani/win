package win

//ref DWORD
//ref HWND
//ref HGLOBAL
//ref POINT
//ref RECT
//ref HINSTANCE
//ref LPARAM
//ref LPPAGESETUPHOOK
//ref LPPAGEPAINTHOOK
//ref LPCWSTR
type PAGESETUPDLG struct {
	LStructSize             DWORD
	HwndOwner               HWND
	HDevMode                HGLOBAL
	HDevNames               HGLOBAL
	Flags                   DWORD
	PtPaperSize             POINT
	RtMinMargin             RECT
	RtMargin                RECT
	HInstance               HINSTANCE
	LCustData               LPARAM
	LpfnPageSetupHook       uintptr // LPPAGESETUPHOOK
	LpfnPagePaintHook       uintptr // LPPAGEPAINTHOOK
	LpPageSetupTemplateName LPCWSTR
	HPageSetupTemplate      HGLOBAL
}
