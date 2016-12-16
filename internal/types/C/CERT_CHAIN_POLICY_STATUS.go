package win

//ref DWORD
//ref LONG
//ref PVOID

type CERT_CHAIN_POLICY_STATUS struct {
	CbSize              DWORD
	DwError             DWORD
	LChainIndex         LONG
	LElementIndex       LONG
	PvExtraPolicyStatus PVOID
}
