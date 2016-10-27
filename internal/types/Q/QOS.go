package win

//ref FLOWSPEC
//ref WSABUF
type QOS struct {
	SendingFlowspec   FLOWSPEC
	ReceivingFlowspec FLOWSPEC
	ProviderSpecific  WSABUF
}
