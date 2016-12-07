package win

//ref RPC_IF_ID

type RPC_IF_ID_VECTOR struct {
	Count uint32
	IfId  [1]*RPC_IF_ID
}
