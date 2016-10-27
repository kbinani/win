package win

//ref LPWSTR
//ref GUID
//ref DWORD
type PDH_DATA_ITEM_PATH_ELEMENTS struct {
	SzMachineName  LPWSTR
	ObjectGUID     GUID
	wItemId        DWORD
	SzInstanceName LPWSTR
}
