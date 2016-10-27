package win

//ref DWORD
//ref PBYTE
type EFS_CERTIFICATE_BLOB struct {
	DwCertEncodingType DWORD
	CbData             DWORD
	PbData             PBYTE
}
