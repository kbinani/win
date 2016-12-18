package win

//ref LPWSTR
//ref DWORD
//ref DWORD_PTR

type VALENT struct {
	Ve_valuename LPWSTR
	Ve_valuelen  DWORD
	Ve_valueptr  DWORD_PTR
	Ve_type      DWORD
}
