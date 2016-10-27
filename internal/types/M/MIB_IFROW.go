package win

//ref WCHAR
//ref IF_INDEX
//ref IFTYPE
//ref DWORD
//ref UCHAR
//ref INTERNAL_IF_OPER_STATUS
type MIB_IFROW struct {
	WszName           [MAX_INTERFACE_NAME_LEN]WCHAR
	DwIndex           IF_INDEX
	DwType            IFTYPE
	DwMtu             DWORD
	DwSpeed           DWORD
	DwPhysAddrLen     DWORD
	BPhysAddr         [MAXLEN_PHYSADDR]UCHAR
	DwAdminStatus     DWORD
	DwOperStatus      INTERNAL_IF_OPER_STATUS
	DwLastChange      DWORD
	DwInOctets        DWORD
	DwInUcastPkts     DWORD
	DwInNUcastPkts    DWORD
	DwInDiscards      DWORD
	DwInErrors        DWORD
	DwInUnknownProtos DWORD
	DwOutOctets       DWORD
	DwOutUcastPkts    DWORD
	DwOutNUcastPkts   DWORD
	DwOutDiscards     DWORD
	DwOutErrors       DWORD
	DwOutQLen         DWORD
	DwDescrLen        DWORD
	BDescr            [MAXLEN_IFDESCR]UCHAR
}
