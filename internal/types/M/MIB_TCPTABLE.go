package win

//ref DWORD
//ref MIB_TCPROW

type MIB_TCPTABLE struct {
	DwNumEntries DWORD
	Table        [ANY_SIZE]MIB_TCPROW
}
