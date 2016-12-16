package win

//ref DWORD
//ref PCERT_RDN_ATTR

type CERT_RDN struct {
	CRDNAttr  DWORD
	RgRDNAttr PCERT_RDN_ATTR
}
