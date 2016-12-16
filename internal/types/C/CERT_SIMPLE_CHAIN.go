package win

//ref DWORD
//ref CERT_TRUST_STATUS
//ref PCERT_CHAIN_ELEMENT
//ref PCERT_TRUST_LIST_INFO
//ref BOOL

type CERT_SIMPLE_CHAIN struct {
	CbSize                      DWORD
	TrustStatus                 CERT_TRUST_STATUS
	CElement                    DWORD
	RgpElement                  *PCERT_CHAIN_ELEMENT
	PTrustListInfo              PCERT_TRUST_LIST_INFO
	FHasRevocationFreshnessTime BOOL
	DwRevocationFreshnessTime   DWORD
}
