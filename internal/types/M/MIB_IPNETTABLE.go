package win

//ref DWORD
//ref MIB_IPNETROW
type MIB_IPNETTABLE struct {
	DwNumEntries DWORD
	Table        [ANY_SIZE]MIB_IPNETROW
}
