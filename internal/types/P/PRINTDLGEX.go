package win

//ref DWORD
//ref HWND
//ref HGLOBAL
//ref HDC
//ref LPPRINTPAGERANGE
//ref HINSTANCE
//ref LPCWSTR
//ref LPUNKNOWN
//ref HPROPSHEETPAGE
type PRINTDLGEX struct {
	LStructSize         DWORD
	HwndOwner           HWND
	HDevMode            HGLOBAL
	HDevNames           HGLOBAL
	HDC                 HDC
	Flags               DWORD
	Flags2              DWORD
	ExclusionFlags      DWORD
	NPageRanges         DWORD
	NMaxPageRanges      DWORD
	LpPageRanges        LPPRINTPAGERANGE
	NMinPage            DWORD
	NMaxPage            DWORD
	NCopies             DWORD
	HInstance           HINSTANCE
	LpPrintTemplateName LPCWSTR
	LpCallback          LPUNKNOWN
	NPropertyPages      DWORD
	LphPropertyPages    *HPROPSHEETPAGE
	NStartPage          DWORD
	DwResultAction      DWORD
}
