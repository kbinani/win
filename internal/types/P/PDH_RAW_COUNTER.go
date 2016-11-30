package win

//ref DWORD
//ref FILETIME
//ref LONGLONG

type PDH_RAW_COUNTER struct {
	CStatus     DWORD
	TimeStamp   FILETIME
	padding1    [pad0for64_4for32]byte
	FirstValue  LONGLONG
	SecondValue LONGLONG
	MultiCount  DWORD
	padding2    [pad0for64_4for32]byte
}
