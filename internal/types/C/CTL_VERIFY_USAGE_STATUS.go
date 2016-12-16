package win

//ref DWORD
//ref PCCTL_CONTEXT

type CTL_VERIFY_USAGE_STATUS struct {
	CbSize          DWORD
	DwError         DWORD
	DwFlags         DWORD
	PpCtl           *PCCTL_CONTEXT
	DwCtlEntryIndex DWORD
	PpSigner        *PCCERT_CONTEXT
	DwSignerIndex   DWORD
}
