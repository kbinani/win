package win

//ref DWORD
//ref HCERTSTORE

type CERT_CHAIN_ENGINE_CONFIG struct {
	CbSize                    DWORD
	HRestrictedRoot           HCERTSTORE
	HRestrictedTrust          HCERTSTORE
	HRestrictedOther          HCERTSTORE
	CAdditionalStore          DWORD
	RghAdditionalStore        *HCERTSTORE
	DwFlags                   DWORD
	DwUrlRetrievalTimeout     DWORD
	MaximumCachedCertificates DWORD
	CycleDetectionModulus     DWORD
	HExclusiveRoot            HCERTSTORE
	HExclusiveTrustedPeople   HCERTSTORE
	DwExclusiveFlags          DWORD
}
