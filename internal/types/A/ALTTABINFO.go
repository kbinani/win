package win

//ref POINT
type ALTTABINFO struct {
	CbSize    uint32
	CItems    int32
	CColumns  int32
	CRows     int32
	IColFocus int32
	IRowFocus int32
	CxItem    int32
	CyItem    int32
	PtStart   POINT
}
