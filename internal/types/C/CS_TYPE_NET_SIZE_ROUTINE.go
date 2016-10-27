package win

//ref RPC_BINDING_HANDLE
//ref IDL_CS_CONVERT
//ref Error_status_t
type CS_TYPE_NET_SIZE_ROUTINE func(hBinding RPC_BINDING_HANDLE, ulNetworkCodeSet uint32, ulLocalBufferSize uint32, conversionType *IDL_CS_CONVERT, pulNetworkBufferSize *uint32, pStatus *Error_status_t)
