package win

//ref DWORD
//ref MIB_IFROW
type MIB_IFTABLE struct {
	DwNumEntries DWORD
	Table        [ANY_SIZE]MIB_IFROW
}
