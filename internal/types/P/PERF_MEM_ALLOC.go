package win

//ref LPVOID
//ref SIZE_T

type PERF_MEM_ALLOC func(AllocSize SIZE_T, pContext LPVOID) LPVOID
