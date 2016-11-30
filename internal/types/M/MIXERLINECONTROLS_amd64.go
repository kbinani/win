package win

type MIXERLINECONTROLS struct {
	CbStruct  DWORD
	DwLineID  DWORD
	union1    DWORD
	CControls DWORD
	Cbmxctrl  DWORD
	storage1  [4]byte
	storage2  [4]byte
}
