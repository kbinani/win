package win

type TBBUTTON struct {
	IBitmap   int32
	IdCommand int32
	FsState   byte
	FsStyle   byte
	BReserved [6]byte
	DwData    *DWORD
	IString   uintptr
}
