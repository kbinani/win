package win

//ref DWORD
//ref MIB_TCP6ROW2

type MIB_TCP6TABLE2 struct {
	DwNumEntries DWORD
	Table        [ANY_SIZE]MIB_TCP6ROW2
}
