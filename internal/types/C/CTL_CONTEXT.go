package win

//ref DWORD
//ref BYTE
//ref PCTL_INFO
//ref HCERTSTORE
//ref HCRYPTMSG

type CTL_CONTEXT struct {
	DwMsgAndCertEncodingType DWORD
	PbCtlEncoded             *BYTE
	CbCtlEncoded             DWORD
	PCtlInfo                 PCTL_INFO
	HCertStore               HCERTSTORE
	HCryptMsg                HCRYPTMSG
	PbCtlContent             *BYTE
	CbCtlContent             DWORD
}
