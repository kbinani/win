package win

//ref PVOID
//ref BYTE
//ref DWORD
//ref BOOL

type PFN_CMSG_STREAM_OUTPUT func(pvArg /*const*/ PVOID, pbData *BYTE, cbData DWORD, fFinal BOOL) BOOL
