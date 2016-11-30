package win

//ref DWORD
//ref MIB_UDPROW

type MIB_UDPTABLE struct {
	DwNumEntries DWORD
	Table        [ANY_SIZE]MIB_UDPROW
}
