package win

//ref RPC_BINDING_HANDLE
//ref IDL_CS_CONVERT
//ref Error_status_t
type CS_TYPE_LOCAL_SIZE_ROUTINE func(hBinding RPC_BINDING_HANDLE, ulNetworkCodeSet uint32, ulNetworkBufferSize uint32, conversionType *IDL_CS_CONVERT, pulLocalBufferSize *uint32, pStatus *Error_status_t)
