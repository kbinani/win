package win

//ref LPCWSTR
//ref UINT
//ref RECT
type POLYTEXT struct {
	X       int32
	Y       int32
	N       UINT
	Lpstr   LPCWSTR
	UiFlags UINT
	Rcl     RECT
	Pdx     *int32
}
