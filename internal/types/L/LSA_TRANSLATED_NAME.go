package win

//ref SID_NAME_USE
//ref LSA_UNICODE_STRING
//ref LONG

type LSA_TRANSLATED_NAME struct {
	Use         SID_NAME_USE
	Name        LSA_UNICODE_STRING
	DomainIndex LONG
}
