package win

//ref DWORD
//ref BYTE
//ref PCERT_INFO
//ref HCERTSTORE

type CERT_CONTEXT struct {
	DwCertEncodingType DWORD
	PbCertEncoded      *BYTE
	CbCertEncoded      DWORD
	PCertInfo          PCERT_INFO
	HCertStore         HCERTSTORE
}
