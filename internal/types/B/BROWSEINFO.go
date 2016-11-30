package win

//ref HWND
//ref PCIDLIST_ABSOLUTE
//ref LPWSTR
//ref LPCWSTR
//ref UINT
//ref BFFCALLBACK
//ref LPARAM

type BROWSEINFO struct {
	HwndOwner      HWND
	PidlRoot       PCIDLIST_ABSOLUTE
	PszDisplayName LPWSTR
	LpszTitle      LPCWSTR
	UlFlags        UINT
	Lpfn           uintptr // BFFCALLBACK
	LParam         LPARAM
	IImage         int32
}
