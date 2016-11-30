package win

//ref PVOID
//ref PIO_STATUS_BLOCK
//ref ULONG

type PIO_APC_ROUTINE func(ApcContext PVOID, IoStatusBlock PIO_STATUS_BLOCK, Reserved ULONG)
