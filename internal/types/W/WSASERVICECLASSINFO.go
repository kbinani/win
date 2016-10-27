package win

//ref LPGUID
//ref LPWSTR
//ref DWORD
//ref LPWSANSCLASSINFO
type WSASERVICECLASSINFO struct {
	LpServiceClassId     LPGUID
	LpszServiceClassName LPWSTR
	DwCount              DWORD
	LpClassInfos         LPWSANSCLASSINFO
}
