package win

//ref SOCKET
//ref SHORT
type WSAPOLLFD struct {
	Fd      SOCKET
	Events  SHORT
	Revents SHORT
}
