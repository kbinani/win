package win

//ref PVOID
//ref CHAR

type MIDL_ES_WRITE func(state PVOID, buffer *CHAR, size uint32)
