package win

//ref HCERTSTORE
//ref PCRL_INFO
//ref BYTE
//ref DWORD
type CRL_CONTEXT struct {
	DwCertEncodingType DWORD
	PbCrlEncoded       *BYTE
	CbCrlEncoded       DWORD
	PCrlInfo           PCRL_INFO
	HCertStore         HCERTSTORE
}
