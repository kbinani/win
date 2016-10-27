package win

//ref DWORD
//ref HWND
//ref HINSTANCE
//ref LPWSTR
//ref WORD
//ref LPARAM
//ref LPFRHOOKPROC
//ref LPCWSTR
type FINDREPLACE struct {
	LStructSize      DWORD
	HwndOwner        HWND
	HInstance        HINSTANCE
	Flags            DWORD
	LpstrFindWhat    LPWSTR
	LpstrReplaceWith LPWSTR
	WFindWhatLen     WORD
	WReplaceWithLen  WORD
	LCustData        LPARAM
	LpfnHook         uintptr // LPFRHOOKPROC
	LpTemplateName   LPCWSTR
}
