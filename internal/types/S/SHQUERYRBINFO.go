package win

//ref DWORD

type SHQUERYRBINFO struct {
	CbSize      DWORD
	I64Size     int64
	I64NumItems int64
}
