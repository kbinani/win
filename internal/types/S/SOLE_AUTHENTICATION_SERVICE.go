package win

//ref DWORD
//ref OLECHAR
//ref HRESULT

type SOLE_AUTHENTICATION_SERVICE struct {
	DwAuthnSvc     DWORD
	DwAuthzSvc     DWORD
	PPrincipalName *OLECHAR
	Hr             HRESULT
}
