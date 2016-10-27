package win

//ref RPC_BINDING_HANDLE
//ref Error_status_t
type CS_TYPE_TO_NETCS_ROUTINE func(hBinding RPC_BINDING_HANDLE, ulNetworkCodeSet uint32, pLocalData uintptr, ulLocalDataLength uint32, pNetworkData *byte, pulNetworkDataLength *uint32, pStatus *Error_status_t)
