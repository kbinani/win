package win

//ref DWORD
//ref HWND
//ref HINSTANCE
//ref LPCWSTR
//ref LPWSTR
//ref LPARAM
//ref LPOFNHOOKPROC
type OPENFILENAME struct {
	LStructSize       DWORD
	HwndOwner         HWND
	HInstance         HINSTANCE
	LpstrFilter       LPCWSTR
	LpstrCustomFilter LPWSTR
	NMaxCustFilter    DWORD
	NFilterIndex      DWORD
	LpstrFile         LPWSTR
	NMaxFile          DWORD
	LpstrFileTitle    LPWSTR
	NMaxFileTitle     DWORD
	LpstrInitialDir   LPCWSTR
	LpstrTitle        LPCWSTR
	Flags             DWORD
	NFileOffset       WORD
	NFileExtension    WORD
	LpstrDefExt       LPCWSTR
	LCustData         LPARAM
	LpfnHook          uintptr // LPOFNHOOKPROC
	LpTemplateName    LPCWSTR
	PvReserved        uintptr
	DwReserved        DWORD
	FlagsEx           DWORD
}
