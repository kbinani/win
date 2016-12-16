package win

//ref DWORD
//ref CERT_TRUST_STATUS
//ref PCERT_SIMPLE_CHAIN
//ref BOOL
//ref GUID

type CERT_CHAIN_CONTEXT struct {
	CbSize                      DWORD
	TrustStatus                 CERT_TRUST_STATUS
	CChain                      DWORD
	RgpChain                    *PCERT_SIMPLE_CHAIN
	CLowerQualityChainContext   DWORD
	RgpLowerQualityChainContext *PCCERT_CHAIN_CONTEXT
	FHasRevocationFreshnessTime BOOL
	DwRevocationFreshnessTime   DWORD
	DwCreateFlags               DWORD
	ChainId                     GUID
}
