package win

//ref DWORD
//ref CUSTDATAITEM
type CUSTDATA struct {
	CCustData   DWORD
	PrgCustData *CUSTDATAITEM
}
