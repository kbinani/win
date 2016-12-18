package win

//ref DWORD
//ref SID
//ref PEFS_HASH_BLOB
//ref LPWSTR

type ENCRYPTION_CERTIFICATE_HASH struct {
	CbTotalLength        DWORD
	PUserSid             *SID
	PHash                PEFS_HASH_BLOB
	LpDisplayInformation LPWSTR
}
