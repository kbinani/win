package win

//ref DWORD
//ref HWND
//ref HDC
//ref LPLOGFONT
//ref INT
//ref COLORREF
//ref LPCFHOOKPROC
//ref LPARAM
//ref LPCWSTR
//ref HINSTANCE
//ref WORD
type CHOOSEFONT struct {
	LStructSize            DWORD
	HwndOwner              HWND
	HDC                    HDC
	LpLogFont              LPLOGFONT
	IPointSize             INT
	Flags                  DWORD
	RgbColors              COLORREF
	LCustData              LPARAM
	LpfnHook               uintptr // LPCFHOOKPROC
	LpTemplateName         LPCWSTR
	HInstance              HINSTANCE
	LpszStyle              LPWSTR
	NFontType              WORD
	___MISSING_ALIGNMENT__ WORD
	NSizeMin               INT
	NSizeMax               INT
}
