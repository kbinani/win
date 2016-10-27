package win

//ref DWORD
//ref LPCWSTR
//ref UINT
type PARSEDURL struct {
	CbSize      DWORD
	PszProtocol LPCWSTR
	CchProtocol UINT
	PszSuffix   LPCWSTR
	CchSuffix   UINT
	NScheme     UINT
}
