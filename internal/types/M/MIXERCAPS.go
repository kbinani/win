package win

//ref WORD
//ref MMVERSION
//ref WCHAR
//ref DWORD

type MIXERCAPS struct {
	WMid           WORD
	WPid           WORD
	VDriverVersion MMVERSION
	SzPname        [MAXPNAMELEN]WCHAR
	FdwSupport     DWORD
	CDestinations  DWORD
}
