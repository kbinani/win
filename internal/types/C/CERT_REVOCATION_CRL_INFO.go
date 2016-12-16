package win

//ref DWORD
//ref PCCRL_CONTEXT
//ref PCRL_ENTRY
//ref BOOL

type CERT_REVOCATION_CRL_INFO struct {
	CbSize           DWORD
	PBaseCrlContext  PCCRL_CONTEXT
	PDeltaCrlContext PCCRL_CONTEXT
	PCrlEntry        PCRL_ENTRY
	FDeltaCrlEntry   BOOL
}
