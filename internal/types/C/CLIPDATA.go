package win

//ref ULONG
//ref LONG
//ref BYTE

type CLIPDATA struct {
	CbSize    ULONG
	UlClipFmt LONG
	PClipData *BYTE
}
