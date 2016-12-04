package win

//ref ULONG
//ref LPWSTR

type CALPWSTR struct {
	CElems ULONG
	PElems *LPWSTR
}
