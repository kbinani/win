package win

//ref ULONGLONG
//ref ULONG

type EVENT_FILTER_DESCRIPTOR struct {
	Ptr  ULONGLONG
	Size ULONG
	Type ULONG
}
