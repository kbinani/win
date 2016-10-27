package win

//ref NDR_RUNDOWN
//ref GENERIC_BINDING_ROUTINE_PAIR
//ref EXPR_EVAL
//ref XMIT_ROUTINE_QUINTUPLE
//ref MALLOC_FREE_STRUCT
//ref USER_MARSHAL_ROUTINE_QUADRUPLE
//ref NDR_CS_ROUTINES
//ref PGENERIC_BINDING_INFO
//ref Handle_t
//ref COMM_FAULT_OFFSETS
//ref ULONG_PTR
//ref NDR_EXPR_DESC
//import unsafe
type MIDL_STUB_DESC struct {
	RpcInterfaceInformation     uintptr
	PfnAllocate                 uintptr // void* (__RPC_API *pfnAllocate)(size_t)
	PfnFree                     uintptr // void (__RPC_API *pfnFree)(void *)
	IMPLICIT_HANDLE_INFO        uintptr
	ApfnNdrRundownRoutines      uintptr // const NDR_RUNDOWN*
	AGenericBindingRoutinePairs uintptr // const GENERIC_BINDING_ROUTINE_PAIR*
	ApfnExprEval                uintptr // const EXPR_EVAL*
	AXmitQuintuple              uintptr // const XMIT_ROUTINE_QUINTUPLE*
	PFormatTypes/*const*/ *byte
	FCheckBounds      int32
	Version           uint32
	PMallocFreeStruct uintptr // MALLOC_FREE_STRUCT*
	MIDLVersion       int32
	CommFaultOffsets/*const*/ *COMM_FAULT_OFFSETS
	AUserMarshalQuadruple uintptr // const USER_MARSHAL_ROUTINE_QUADRUPLE*
	NotifyRoutineTable    uintptr // const NDR_NOTIFY_ROUTINE*
	MFlags                ULONG_PTR
	CsRoutineTables/*const*/ *NDR_CS_ROUTINES
	ProxyServerInfo uintptr
	PExprInfo/*const*/ *NDR_EXPR_DESC
}

func (this MIDL_STUB_DESC) PAutoHandle() *Handle_t {
	return (*Handle_t)(unsafe.Pointer(&this.IMPLICIT_HANDLE_INFO))
}

func (this MIDL_STUB_DESC) PPrimitiveHandle() *Handle_t {
	return (*Handle_t)(unsafe.Pointer(&this.IMPLICIT_HANDLE_INFO))
}

func (this MIDL_STUB_DESC) PGenericBindingInfo() PGENERIC_BINDING_INFO {
	return (PGENERIC_BINDING_INFO)(unsafe.Pointer(&this.IMPLICIT_HANDLE_INFO))
}
