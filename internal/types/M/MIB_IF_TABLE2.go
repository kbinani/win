package win

//ref ULONG
//ref MIB_IF_ROW2

type MIB_IF_TABLE2 struct {
	NumEntries ULONG
	padding1   [pad0for64_4for32]byte
	Table      [ANY_SIZE]MIB_IF_ROW2
}
