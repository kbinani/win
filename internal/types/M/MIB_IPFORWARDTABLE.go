package win

//ref MIB_IPFORWARDROW
//ref DWORD
type MIB_IPFORWARDTABLE struct {
	DwNumEntries DWORD
	Table        [ANY_SIZE]MIB_IPFORWARDROW
}
