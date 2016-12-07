package win

//ref PVOID
//ref CHAR

type MIDL_ES_ALLOC func(state PVOID, pbuffer **CHAR, psize *uint32)
