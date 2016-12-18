package win

//ref ULONG
//ref ACTRL_ACCESS_ENTRY

type ACTRL_ACCESS_ENTRY_LIST struct {
	CEntries    ULONG
	PAccessList *ACTRL_ACCESS_ENTRY
}
