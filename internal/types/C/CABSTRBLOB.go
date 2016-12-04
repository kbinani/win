package win

//ref ULONG
//ref BSTRBLOB

type CABSTRBLOB struct {
	CElems ULONG
	PElems *BSTRBLOB
}
