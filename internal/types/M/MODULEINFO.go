package win

//ref LPVOID
//ref DWORD
type MODULEINFO struct {
	LpBaseOfDll LPVOID
	SizeOfImage DWORD
	EntryPoint  LPVOID
}
