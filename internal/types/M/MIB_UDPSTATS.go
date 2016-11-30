package win

//ref DWORD

type MIB_UDPSTATS struct {
	DwInDatagrams  DWORD
	DwNoPorts      DWORD
	DwInErrors     DWORD
	DwOutDatagrams DWORD
	DwNumAddrs     DWORD
}
