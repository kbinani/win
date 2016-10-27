package win

//ref ULONG
//ref LUID_AND_ATTRIBUTES
type TOKEN_PRIVILEGES struct {
	PrivilegeCount ULONG
	Privileges     [ANYSIZE_ARRAY]LUID_AND_ATTRIBUTES
}
