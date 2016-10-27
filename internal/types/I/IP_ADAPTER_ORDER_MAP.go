package win

//ref ULONG
type IP_ADAPTER_ORDER_MAP struct {
	NumAdapters  ULONG
	AdapterOrder [1]ULONG
}
