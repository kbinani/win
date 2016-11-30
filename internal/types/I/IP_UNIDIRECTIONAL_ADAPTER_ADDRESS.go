package win

//ref ULONG
//ref IPAddr

type IP_UNIDIRECTIONAL_ADAPTER_ADDRESS struct {
	NumAdapters ULONG
	Address     [1]IPAddr
}
