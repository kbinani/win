package win

//ref DWORD
//ref HWND
//ref LPCWSTR

type CRYPTPROTECT_PROMPTSTRUCT struct {
	CbSize        DWORD
	DwPromptFlags DWORD
	HwndApp       HWND
	SzPrompt      LPCWSTR
}
