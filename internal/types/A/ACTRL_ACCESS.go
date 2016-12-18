package win

//ref ULONG
//ref PACTRL_PROPERTY_ENTRY

type ACTRL_ACCESS struct {
	CEntries            ULONG
	PPropertyAccessList PACTRL_PROPERTY_ENTRY
}
