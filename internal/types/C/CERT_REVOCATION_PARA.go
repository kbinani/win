package win

//ref DWORD
//ref PCCERT_CONTEXT
//ref HCERTSTORE
//ref LPFILETIME

type CERT_REVOCATION_PARA struct {
	CbSize       DWORD
	PIssuerCert  PCCERT_CONTEXT
	CCertStore   DWORD
	RgCertStore  *HCERTSTORE
	HCrlStore    HCERTSTORE
	PftTimeToUse LPFILETIME
}
