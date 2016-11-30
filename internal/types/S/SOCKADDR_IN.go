package win

//ref ADDRESS_FAMILY
//ref USHORT
//ref IN_ADDR
//ref CHAR

type SOCKADDR_IN struct {
	Sin_family ADDRESS_FAMILY
	Sin_port   USHORT
	Sin_addr   IN_ADDR
	Sin_zero   [8]CHAR
}
