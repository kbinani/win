package win

type PDH_RAW_COUNTER struct {
	CStatus     DWORD
	TimeStamp   FILETIME
	padding1    [4]byte
	FirstValue  LONGLONG
	SecondValue LONGLONG
	MultiCount  DWORD
	padding2    [4]byte
}
