package win

//ref SID_NAME_USE
//ref ULONG
//ref LONG

type LSA_TRANSLATED_SID struct {
	Use         SID_NAME_USE
	RelativeId  ULONG
	DomainIndex LONG
}
