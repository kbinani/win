package win

//ref DWORD
//ref LPCSTR
//ref BOOL
//ref PCERT_REVOCATION_CRL_INFO
//ref LPVOID

type CERT_REVOCATION_INFO struct {
	CbSize             DWORD
	DwRevocationResult DWORD
	PszRevocationOid   LPCSTR
	PvOidSpecificInfo  LPVOID
	FHasFreshnessTime  BOOL
	DwFreshnessTime    DWORD
	PCrlInfo           PCERT_REVOCATION_CRL_INFO
}
