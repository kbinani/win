package win

//ref DWORD
//ref MIB_UDP6ROW

type MIB_UDP6TABLE struct {
	DwNumEntries DWORD
	Table        [ANY_SIZE]MIB_UDP6ROW
}
