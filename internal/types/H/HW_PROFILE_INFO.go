package win

//ref DWORD
//ref WCHAR

type HW_PROFILE_INFO struct {
	DwDockInfo      DWORD
	SzHwProfileGuid [HW_PROFILE_GUIDLEN]WCHAR
	SzHwProfileName [MAX_PROFILE_LEN]WCHAR
}
