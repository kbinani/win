package win

type PDH_RAW_COUNTER struct {
	CStatus     DWORD
	TimeStamp   FILETIME
	FirstValue  LONGLONG
	SecondValue LONGLONG
	MultiCount  DWORD
}
