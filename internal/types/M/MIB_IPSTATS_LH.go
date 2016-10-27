package win

//ref MIB_IPSTATS_FORWARDING
//ref DWORD
type MIB_IPSTATS_LH struct {
	Forwarding        MIB_IPSTATS_FORWARDING
	DwDefaultTTL      DWORD
	DwInReceives      DWORD
	DwInHdrErrors     DWORD
	DwInAddrErrors    DWORD
	DwForwDatagrams   DWORD
	DwInUnknownProtos DWORD
	DwInDiscards      DWORD
	DwInDelivers      DWORD
	DwOutRequests     DWORD
	DwRoutingDiscards DWORD
	DwOutDiscards     DWORD
	DwOutNoRoutes     DWORD
	DwReasmTimeout    DWORD
	DwReasmReqds      DWORD
	DwReasmOks        DWORD
	DwReasmFails      DWORD
	DwFragOks         DWORD
	DwFragFails       DWORD
	DwFragCreates     DWORD
	DwNumIf           DWORD
	DwNumAddr         DWORD
	DwNumRoutes       DWORD
}
