package win

//ref CHAR
type Sockaddr struct {
	Sa_family uint16
	Sa_data   [14]CHAR
}
