package win

//ref RPC_BINDING_HANDLE

type RPC_BINDING_VECTOR struct {
	Count    uint32
	BindingH [1]RPC_BINDING_HANDLE
}
