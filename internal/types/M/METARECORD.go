package win

//ref DWORD
//ref WORD
type METARECORD struct {
	RdSize     DWORD
	RdFunction WORD
	RdParm     [1]WORD
}
