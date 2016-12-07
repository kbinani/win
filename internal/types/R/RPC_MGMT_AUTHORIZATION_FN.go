package win

//ref RPC_BINDING_HANDLE
//ref RPC_STATUS

type RPC_MGMT_AUTHORIZATION_FN func(ClientBinding RPC_BINDING_HANDLE, RequestedMgmtOperation uint32, Status *RPC_STATUS) int32
