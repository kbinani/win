package win

//ref DWORD
//ref LPWSTR
//ref COAUTHIDENTITY

type COAUTHINFO struct {
	DwAuthnSvc           DWORD
	DwAuthzSvc           DWORD
	PwszServerPrincName  LPWSTR
	DwAuthnLevel         DWORD
	DwImpersonationLevel DWORD
	PAuthIdentityData    *COAUTHIDENTITY
	DwCapabilities       DWORD
}
