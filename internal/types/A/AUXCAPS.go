package win

//ref WORD
//ref MMVERSION
//ref WCHAR
//ref DWORD
type AUXCAPS struct {
	WMid           WORD
	WPid           WORD
	VDriverVersion MMVERSION
	SzPname        [MAXPNAMELEN]WCHAR
	WTechnology    WORD
	WReserved1     WORD
	DwSupport      DWORD
}
