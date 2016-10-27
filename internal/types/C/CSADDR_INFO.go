package win

//ref SOCKET_ADDRESS
//ref INT
type CSADDR_INFO struct {
	LocalAddr   SOCKET_ADDRESS
	RemoteAddr  SOCKET_ADDRESS
	ISocketType INT
	IProtocol   INT
}
