package win

//ref UINT
//ref HANDLE
//ref DWORD_PTR
//ref POINT
type HELPINFO struct {
	CbSize       UINT
	IContextType int32
	ICtrlId      int32
	HItemHandle  HANDLE
	DwContextId  DWORD_PTR
	MousePos     POINT
}
