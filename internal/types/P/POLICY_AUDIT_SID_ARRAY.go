package win

//ref ULONG
//ref PSID
type POLICY_AUDIT_SID_ARRAY struct {
	UsersCount   ULONG
	UserSidArray *PSID
}
