package win

//ref MIB_IPFORWARD_PROTO
//ref DWORD
//ref IF_INDEX
//ref MIB_IPFORWARD_TYPE
type MIB_IPFORWARDROW struct {
	DwForwardDest      DWORD
	DwForwardMask      DWORD
	DwForwardPolicy    DWORD
	DwForwardNextHop   DWORD
	DwForwardIfIndex   IF_INDEX
	ForwardType        MIB_IPFORWARD_TYPE
	ForwardProto       MIB_IPFORWARD_PROTO
	DwForwardAge       DWORD
	DwForwardNextHopAS DWORD
	DwForwardMetric1   DWORD
	DwForwardMetric2   DWORD
	DwForwardMetric3   DWORD
	DwForwardMetric4   DWORD
	DwForwardMetric5   DWORD
}
