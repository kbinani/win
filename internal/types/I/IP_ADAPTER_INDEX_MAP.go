package win

//ref WCHAR
//ref ULONG
type IP_ADAPTER_INDEX_MAP struct {
	Index ULONG
	Name  [MAX_ADAPTER_NAME]WCHAR
}
