package win

//ref DWORD
//ref GUID
//ref WCHAR
//ref PWSTR

type SIP_ADD_NEWPROVIDER struct {
	CbStruct               DWORD
	PgSubject              *GUID
	PwszDLLFileName        *WCHAR
	PwszMagicNumber        *WCHAR
	PwszIsFunctionName     *WCHAR
	PwszGetFuncName        *WCHAR
	PwszPutFuncName        *WCHAR
	PwszCreateFuncName     *WCHAR
	PwszVerifyFuncName     *WCHAR
	PwszRemoveFuncName     *WCHAR
	PwszIsFunctionNameFmt2 *WCHAR
	PwszGetCapFuncName     PWSTR
}
