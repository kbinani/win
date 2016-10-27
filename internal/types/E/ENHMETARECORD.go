package win

//ref DWORD
type ENHMETARECORD struct {
	IType DWORD
	NSize DWORD
	DParm [1]DWORD
}
