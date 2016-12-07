package win

//ref RPC_IF_HANDLE
//ref PVOID
//ref RPC_STATUS

type RPC_IF_CALLBACK_FN func(InterfaceUuid RPC_IF_HANDLE, Context PVOID) RPC_STATUS
