package win

//ref TRUSTEE
//ref ULONG
//ref ACCESS_RIGHTS
//ref INHERIT_FLAGS
//ref LPWSTR

type ACTRL_ACCESS_ENTRY struct {
	Trustee            TRUSTEE
	FAccessFlags       ULONG
	Access             ACCESS_RIGHTS
	ProvSpecificAccess ACCESS_RIGHTS
	Inheritance        INHERIT_FLAGS
	LpInheritProperty  LPWSTR
}
