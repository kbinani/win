package win

//ref HCONV
//ref HSZ
//ref HCONVLIST
//ref CONVCONTEXT
//ref HWND
type CONVINFO struct {
	Cb            uint32
	HUser         *uint32 // DWORD_PTR
	HConvPartner  HCONV
	HszSvcPartner HSZ
	HszServiceReq HSZ
	HszTopic      HSZ
	HszItem       HSZ
	WFmt          uint32 // UINT
	WType         uint32 // UINT
	WStatus       uint32 // UINT
	WConvst       uint32 // UINT
	WLastError    uint32 // UINT
	HConvList     HCONVLIST
	ConvCtxt      CONVCONTEXT
	Hwnd          HWND
	HwndPartner   HWND
}
