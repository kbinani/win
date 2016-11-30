package win

type TBBUTTON struct {
	IBitmap   int32
	IdCommand int32
	FsState   byte
	FsStyle   byte
	BReserved [2 * pad3for64_1for32]byte
	DwData    *DWORD
	IString   uintptr
}
