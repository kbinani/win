package win

import "unsafe"

//ref DWORD
//ref TCP_RTO_ALGORITHM

type MIB_TCPSTATS_LH struct {
	storage1       DWORD
	DwRtoMin       DWORD
	DwRtoMax       DWORD
	DwMaxConn      DWORD
	DwActiveOpens  DWORD
	DwPassiveOpens DWORD
	DwAttemptFails DWORD
	DwEstabResets  DWORD
	DwCurrEstab    DWORD
	DwInSegs       DWORD
	DwOutSegs      DWORD
	DwRetransSegs  DWORD
	DwInErrs       DWORD
	DwOutRsts      DWORD
	DwNumConns     DWORD
}

func (this *MIB_TCPSTATS_LH) DwRtoAlgorithm() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage1))
}

func (this *MIB_TCPSTATS_LH) RtoAlgorithm() *TCP_RTO_ALGORITHM {
	return (*TCP_RTO_ALGORITHM)(unsafe.Pointer(&this.storage1))
}
