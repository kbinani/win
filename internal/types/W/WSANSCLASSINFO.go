package win

//ref LPWSTR
//ref DWORD
//ref LPVOID
type WSANSCLASSINFO struct {
	LpszName    LPWSTR
	DwNameSpace DWORD
	DwValueType DWORD
	DwValueSize DWORD
	LpValue     LPVOID
}
