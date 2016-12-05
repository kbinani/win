package win

//ref DWORD

type BIND_OPTS struct {
	CbStruct            DWORD
	GrfFlags            DWORD
	GrfMode             DWORD
	DwTickCountDeadline DWORD
}
