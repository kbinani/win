package win

//ref PVOID
//ref RPCOLEDATAREP
//ref ULONG

type RPCOLEMESSAGE struct {
	Reserved1          PVOID
	DataRepresentation RPCOLEDATAREP
	Buffer             PVOID
	CbBuffer           ULONG
	IMethod            ULONG
	Reserved2          [5]PVOID
	RpcFlags           ULONG
}
