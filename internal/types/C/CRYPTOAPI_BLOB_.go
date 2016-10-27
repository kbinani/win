package win

//ref DWORD
//ref BYTE
type CRYPTOAPI_BLOB_ struct {
	CbData DWORD
	PbData *BYTE
}
