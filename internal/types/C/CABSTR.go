package win

//ref ULONG
//ref BSTR

type CABSTR struct {
	CElems ULONG
	PElems *BSTR
}
