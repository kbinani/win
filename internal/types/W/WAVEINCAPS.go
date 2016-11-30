package win

//ref WORD
//ref MMVERSION
//ref WCHAR
//ref DWORD

type WAVEINCAPS struct {
	WMid           WORD
	WPid           WORD
	VDriverVersion MMVERSION
	SzPname        [MAXPNAMELEN]WCHAR
	DwFormats      DWORD
	WChannels      WORD
	WReserved1     WORD
}
