package win

//ref WORD
//ref WCHAR
//ref MMVERSION
//ref DWORD

type MIDIOUTCAPS struct {
	WMid           WORD
	WPid           WORD
	VDriverVersion MMVERSION
	SzPname        [MAXPNAMELEN]WCHAR
	WTechnology    WORD
	WVoices        WORD
	WNotes         WORD
	WChannelMask   WORD
	DwSupport      DWORD
}
