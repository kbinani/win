package win

type SOCKADDR struct {
	Sa_family ADDRESS_FAMILY
	Sa_data   [14]CHAR
}
