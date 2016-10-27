package win

//ref ULONG
//ref SID_AND_ATTRIBUTES
type TOKEN_GROUPS struct {
	GroupCount ULONG
	Groups     [ANYSIZE_ARRAY]SID_AND_ATTRIBUTES
}
