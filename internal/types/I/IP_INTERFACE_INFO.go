package win

//ref IP_ADAPTER_INDEX_MAP
//ref LONG
type IP_INTERFACE_INFO struct {
	NumAdapters LONG
	Adapter     [1]IP_ADAPTER_INDEX_MAP
}
