package win

//ref DWORD
//ref CERT_ENHKEY_USAGE

type CERT_USAGE_MATCH struct {
	DwType DWORD
	Usage  CERT_ENHKEY_USAGE
}
