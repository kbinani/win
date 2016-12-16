package win

//ref DWORD
//ref BYTE

type CRYPT_KEY_PROV_PARAM struct {
	DwParam DWORD
	PbData  *BYTE
	CbData  DWORD
	DwFlags DWORD
}
