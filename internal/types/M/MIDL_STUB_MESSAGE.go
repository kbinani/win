package win

//ref ULONG_PTR
//ref PRPC_MESSAGE
//ref Handle_t
//ref MIDL_STUB_DESC
//ref IRpcChannelBuffer
//ref INT_PTR
//ref ULONG_PTR
//ref NDR_SCONTEXT
//ref PARRAY_INFO
//ref FULL_PTR_XLAT_TABLES
type MIDL_STUB_MESSAGE struct {
	RpcMsg                 PRPC_MESSAGE
	Buffer                 *byte
	BufferStart            *byte
	BufferEnd              *byte
	BufferMark             *byte
	BufferLength           uint32
	MemorySize             uint32
	Memory                 *byte
	IsClient               byte
	Pad                    byte
	UFlags2                uint16
	ReuseBuffer            int32
	PAllocAllNodesContext  uintptr // struct NDR_ALLOC_ALL_NODES_CONTEXT*
	PPointerQueueState     uintptr // struct NDR_POINTER_QUEUE_STATE*
	IgnoreEmbeddedPointers int32
	PointerBufferMark      *byte
	CorrDespIncrement      byte
	uFlags                 byte
	UniquePtrCount         uint16
	MaxCount               ULONG_PTR
	Offset                 uint32
	ActualCount            uint32
	PfnAllocate            uintptr // void*(__RPC_API *pfnAllocate)(size_t)
	PfnFree                uintptr // void(__RPC_API *pfnFree)(void*)
	StackTop               *byte
	PPresentedType         *byte
	PTransmitType          *byte
	SavedHandle            Handle_t
	StubDesc/*const*/ *MIDL_STUB_DESC
	FullPtrXlatTables *FULL_PTR_XLAT_TABLES
	FullPtrRefId      uint32
	PointerLength     uint32
	fBitField32       uint32
	/*
	   int                             fInDontFree       :1;
	   int                             fDontCallFreeInst :1;
	   int                             fInOnlyParam      :1;
	   int                             fHasReturn        :1;
	   int                             fHasExtensions    :1;
	   int                             fHasNewCorrDesc   :1;
	   int                             fIsIn             :1;
	   int                             fIsOut            :1;
	   int                             fIsOicf           :1;
	   int                             fBufferValid      :1;
	   int                             fHasMemoryValidateCallback: 1;
	   int                             fInFree             :1;
	   int                             fNeedMCCP         :1;
	   int                             fUnused           :3;
	   int                             fUnused2          :16;
	*/
	DwDestContext       uint32
	PvDestContext       uintptr
	SavedContextHandles *NDR_SCONTEXT
	ParamNumber         int32
	PRpcChannelBuffer   *IRpcChannelBuffer
	PArrayInfo          PARRAY_INFO
	SizePtrCountArray   *uint32
	SizePtrOffsetArray  *uint32
	SizePtrLengthArray  *uint32
	PArgQueue           uintptr
	DwStubPhase         uint32
	LowStackMark        uintptr
	PAsyncMsg           uintptr // PNDR_ASYNC_MESSAGE
	PCorrInfo           uintptr // PNDR_CORRELATION_INFO
	PCorrMemory         *byte
	PMemoryList         uintptr
	PCSInfo             INT_PTR
	ConformanceMark     *byte
	VarianceMark        *byte
	Unused              INT_PTR
	PContext            uintptr // struct _NDR_PROC_CONTEXT*
	ContextHandleHash   uintptr
	PUserMarshalList    uintptr
	Reserved51_3        INT_PTR
	Reserved51_4        INT_PTR
	Reserved51_5        INT_PTR
}
