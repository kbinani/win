package win

//ref DWORD
//ref PCERT_RDN

type CERT_NAME_INFO struct {
	CRDN  DWORD
	RgRDN PCERT_RDN
}
