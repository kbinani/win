package win

//ref PVOID
//ref CHAR

type MIDL_ES_READ func(state PVOID, pbuffer **CHAR, psize *uint32)
