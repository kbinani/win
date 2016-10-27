package win

//ref DWORD
//ref MIB_IPADDRROW
type MIB_IPADDRTABLE struct {
	DwNumEntries DWORD
	Table        [ANY_SIZE]MIB_IPADDRROW
}
