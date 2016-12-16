package win

//ref DWORD
//ref PVOID

type CERT_CHAIN_POLICY_PARA struct {
	CbSize            DWORD
	DwFlags           DWORD
	PvExtraPolicyPara PVOID
}
