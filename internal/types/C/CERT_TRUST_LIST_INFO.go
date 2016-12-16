package win

//ref DWORD
//ref PCTL_ENTRY
//ref PCCTL_CONTEXT

type CERT_TRUST_LIST_INFO struct {
	CbSize      DWORD
	PCtlEntry   PCTL_ENTRY
	PCtlContext PCCTL_CONTEXT
}
