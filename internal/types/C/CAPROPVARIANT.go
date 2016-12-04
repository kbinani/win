package win

//ref ULONG
//ref PROPVARIANT

type CAPROPVARIANT struct {
	CElems ULONG
	PElems *PROPVARIANT
}
