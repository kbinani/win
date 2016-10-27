package win

//ref DWORD
//ref IF_INDEX
type MIB_IPADDRROW_XP struct {
	DwAddr      DWORD
	DwIndex     IF_INDEX
	DwMask      DWORD
	DwBCastAddr DWORD
	DwReasmSize DWORD
	Unused1     uint16
	WType       uint16
}
