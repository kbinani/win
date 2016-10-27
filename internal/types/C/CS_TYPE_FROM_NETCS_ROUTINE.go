package win

//ref RPC_BINDING_HANDLE
//ref Error_status_t
type CS_TYPE_FROM_NETCS_ROUTINE func(hBinding RPC_BINDING_HANDLE, ulNetworkCodeSet uint32, pNetworkData *byte, ulNetworkDataLength uint32, ulLocalBufferSize uint32, pLocalData uintptr, pulLocalDataLength *uint32, pStatus *Error_status_t)
