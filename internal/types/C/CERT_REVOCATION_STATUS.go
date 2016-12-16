package win

//ref DWORD
//ref BOOL

type CERT_REVOCATION_STATUS struct {
	CbSize            DWORD
	DwIndex           DWORD
	DwError           DWORD
	DwReason          DWORD
	FHasFreshnessTime BOOL
	DwFreshnessTime   DWORD
}
