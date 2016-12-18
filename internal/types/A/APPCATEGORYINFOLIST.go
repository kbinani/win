package win

//ref DWORD
//ref APPCATEGORYINFO

type APPCATEGORYINFOLIST struct {
	CCategory     DWORD
	PCategoryInfo *APPCATEGORYINFO
}
