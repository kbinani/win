package win

//ref INT
//ref ULONG
type NUMPARSE struct {
	CDig       INT
	DwInFlags  ULONG
	DwOutFlags ULONG
	CchUsed    INT
	NBaseShift INT
	NPwr10     INT
}
