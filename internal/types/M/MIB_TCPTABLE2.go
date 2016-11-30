package win

//ref DWORD
//ref MIB_TCPROW2

type MIB_TCPTABLE2 struct {
	DwNumEntries DWORD
	Table        [ANY_SIZE]MIB_TCPROW2
}
