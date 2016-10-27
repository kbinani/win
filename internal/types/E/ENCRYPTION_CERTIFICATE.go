package win

//ref DWORD
//ref SID
//ref EFS_CERTIFICATE_BLOB
type ENCRYPTION_CERTIFICATE struct {
	CbTotalLength DWORD
	PUserSid      *SID
	PCertBlob     *EFS_CERTIFICATE_BLOB
}
