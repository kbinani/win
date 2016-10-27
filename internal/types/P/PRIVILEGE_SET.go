package win

//ref ULONG
//ref LUID_AND_ATTRIBUTES
type PRIVILEGE_SET struct {
	PrivilegeCount ULONG
	Control        ULONG
	Privilege      [ANYSIZE_ARRAY]LUID_AND_ATTRIBUTES
}
