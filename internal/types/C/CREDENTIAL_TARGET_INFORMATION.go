package win

//ref LPWSTR
//ref DWORD
//ref ULONG
//ref LPDWORD

type CREDENTIAL_TARGET_INFORMATION struct {
	TargetName        LPWSTR
	NetbiosServerName LPWSTR
	DnsServerName     LPWSTR
	NetbiosDomainName LPWSTR
	DnsDomainName     LPWSTR
	DnsTreeName       LPWSTR
	PackageName       LPWSTR
	Flags             ULONG
	CredTypeCount     DWORD
	CredTypes         LPDWORD
}
