package win

//ref RPC_BINDING_HANDLE
//ref Error_status_t
type CS_TAG_GETTING_ROUTINE func(hBinding RPC_BINDING_HANDLE, fServerSide int32, pulSendingTag *uint32, pulDesiredReceivingTag *uint32, pulReceivingTag *uint32, pStatus *Error_status_t)
