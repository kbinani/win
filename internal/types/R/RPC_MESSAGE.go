package win

//ref PRPC_SYNTAX_IDENTIFIER
//ref RPC_BINDING_HANDLE
type RPC_MESSAGE struct {
	Handle                  RPC_BINDING_HANDLE
	DataRepresentation      uint32
	Buffer                  uintptr
	BufferLength            uint32
	ProcNum                 uint32
	TransferSyntax          PRPC_SYNTAX_IDENTIFIER
	RpcInterfaceInformation uintptr
	ReservedForRuntime      uintptr
	ManagerEpv              uintptr
	ImportContext           uintptr
	RpcFlags                uint32
}
