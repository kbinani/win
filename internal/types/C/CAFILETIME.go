package win

//ref ULONG
//ref FILETIME

type CAFILETIME struct {
	CElems ULONG
	PElems *FILETIME
}
