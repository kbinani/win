package win

//ref SID_NAME_USE
//ref PSID
//ref LONG
//ref ULONG

type LSA_TRANSLATED_SID2 struct {
	Use         SID_NAME_USE
	Sid         PSID
	DomainIndex LONG
	Flags       ULONG
}
