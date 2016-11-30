package win

//ref DWORD
//ref MIB_TCP6ROW

type MIB_TCP6TABLE struct {
	DwNumEntries DWORD
	Table        [ANY_SIZE]MIB_TCP6ROW
}
