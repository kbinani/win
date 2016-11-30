package win

//ref WORD
//ref MMVERSION
//ref WCHAR
//ref DWORD

type WAVEOUTCAPS struct {
	WMid           WORD
	WPid           WORD
	VDriverVersion MMVERSION
	SzPname        [MAXPNAMELEN]WCHAR
	DwFormats      DWORD
	WChannels      WORD
	WReserved1     WORD
	DwSupport      DWORD
}
