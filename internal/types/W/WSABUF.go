package win

//ref ULONG
//ref CHAR
type WSABUF struct {
	Len ULONG
	Buf *CHAR
}
