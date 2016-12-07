package win

//ref RPC_ASYNC_EVENT
//ref RPC_NOTIFICATION_TYPES
//ref RPC_ASYNC_NOTIFICATION_INFO
//ref LONG_PTR

type RPC_ASYNC_STATE struct {
	Size             uint32
	Signature        uint32
	Lock             int32
	Flags            uint32
	StubInfo         uintptr
	UserInfo         uintptr
	RuntimeInfo      uintptr
	Event            RPC_ASYNC_EVENT
	NotificationType RPC_NOTIFICATION_TYPES
	U                RPC_ASYNC_NOTIFICATION_INFO
	Reserved         [4]LONG_PTR
}
