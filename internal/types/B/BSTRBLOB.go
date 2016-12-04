package win

//ref ULONG
//ref BYTE

type BSTRBLOB struct {
	CbSize ULONG
	PData  *BYTE
}
