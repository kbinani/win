package win

//ref DWORD
//ref LPWSTR
//ref COAUTHINFO

type COSERVERINFO struct {
	DwReserved1 DWORD
	PwszName    LPWSTR
	PAuthInfo   *COAUTHINFO
	DwReserved2 DWORD
}
