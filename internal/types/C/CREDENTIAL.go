package win

//ref LPWSTR
//ref FILETIME
//ref CREDENTIAL_ATTRIBUTE
type CREDENTIAL struct {
	Flags              uint32
	Type               uint32
	TargetName         LPWSTR
	Comment            LPWSTR
	LastWritten        FILETIME
	CredentialBlobSize uint32
	CredentialBlob     *byte
	Persist            uint32
	AttributeCount     uint32
	Attributes         *CREDENTIAL_ATTRIBUTE
	TargetAlias        LPWSTR
	UserName           LPWSTR
}
