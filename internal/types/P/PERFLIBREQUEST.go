package win

//ref ULONG
//ref PVOID

type PERFLIBREQUEST func(RequestCode ULONG, Buffer PVOID, BufferSize ULONG) ULONG
