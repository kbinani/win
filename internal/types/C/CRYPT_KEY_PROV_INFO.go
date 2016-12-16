package win

//ref LPWSTR
//ref DWORD
//ref PCRYPT_KEY_PROV_PARAM

type CRYPT_KEY_PROV_INFO struct {
	PwszContainerName LPWSTR
	PwszProvName      LPWSTR
	DwProvType        DWORD
	DwFlags           DWORD
	CProvParam        DWORD
	RgProvParam       PCRYPT_KEY_PROV_PARAM
	DwKeySpec         DWORD
}
