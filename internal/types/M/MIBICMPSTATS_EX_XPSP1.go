package win

//ref DWORD
type MIBICMPSTATS_EX_XPSP1 struct {
	DwMsgs        DWORD
	DwErrors      DWORD
	RgdwTypeCount [256]DWORD
}
