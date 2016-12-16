package win

//ref DWORD

type CERT_TRUST_STATUS struct {
	DwErrorStatus DWORD
	DwInfoStatus  DWORD
}
