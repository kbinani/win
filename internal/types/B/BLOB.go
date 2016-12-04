package win

//ref ULONG
//ref BYTE

type BLOB struct {
	CbSize    ULONG
	PBlobData *BYTE
}
