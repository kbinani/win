package win

//ref IF_INDEX
//ref DWORD
//ref UCHAR
//ref MIB_IPNET_TYPE
type MIB_IPNETROW_LH struct {
	DwIndex       IF_INDEX
	DwPhysAddrLen DWORD
	BPhysAddr     [MAXLEN_PHYSADDR]UCHAR
	DwAddr        DWORD
	Type          MIB_IPNET_TYPE
}
