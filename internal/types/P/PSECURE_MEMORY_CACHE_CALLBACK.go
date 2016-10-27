package win

//ref PVOID
//ref SIZE_T
//ref BOOLEAN
type PSECURE_MEMORY_CACHE_CALLBACK func(Addr PVOID, Range SIZE_T) BOOLEAN
