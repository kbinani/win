package win

type MALLOC_FREE_STRUCT struct {
	PfnAllocate uintptr // void* (__RPC_USER *pfnAllocate)(size_t)
	PfnFree     uintptr // void (__RPC_USER *pfnFree)(void *)
}
