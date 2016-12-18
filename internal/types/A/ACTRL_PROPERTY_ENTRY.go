package win

//ref LPWSTR
//ref ULONG
//ref PACTRL_ACCESS_ENTRY_LIST

type ACTRL_PROPERTY_ENTRY struct {
	LpProperty       LPWSTR
	PAccessEntryList PACTRL_ACCESS_ENTRY_LIST
	FListFlags       ULONG
}
