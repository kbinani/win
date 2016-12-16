package win

//ref DWORD
//ref PCCERT_CONTEXT
//ref CERT_TRUST_STATUS
//ref PCERT_REVOCATION_INFO
//ref PCERT_ENHKEY_USAGE
//ref LPCWSTR

type CERT_CHAIN_ELEMENT struct {
	CbSize                DWORD
	PCertContext          PCCERT_CONTEXT
	TrustStatus           CERT_TRUST_STATUS
	PRevocationInfo       PCERT_REVOCATION_INFO
	PIssuanceUsage        PCERT_ENHKEY_USAGE
	PApplicationUsage     PCERT_ENHKEY_USAGE
	PwszExtendedErrorInfo LPCWSTR
}
