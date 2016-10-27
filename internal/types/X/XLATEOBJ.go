package win

//ref ULONG
//ref FLONG
//ref USHORT
type XLATEOBJ struct {
	IUniq    ULONG
	FlXlate  FLONG
	ISrcType USHORT
	IDstType USHORT
	CEntries ULONG
	PulXlate *ULONG
}
