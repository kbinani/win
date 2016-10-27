package win

//ref LPSOCKADDR
//ref INT
//ref LPWSABUF
//ref ULONG
type WSAMSG struct {
	Name          LPSOCKADDR
	Namelen       INT
	LpBuffers     LPWSABUF
	DwBufferCount ULONG
	Control       WSABUF
	DwFlags       ULONG
}
