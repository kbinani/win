package win

//ref ULONG
//ref HWND
//ref LPCOLESTR
//ref CLSID
//ref LCID
//ref DISPID
type OCPFIPARAMS struct {
	CbStructSize          ULONG
	HWndOwner             HWND
	X                     int32
	Y                     int32
	LpszCaption           LPCOLESTR
	CObjects              ULONG
	LplpUnk               *LPUNKNOWN
	CPages                ULONG
	LpPages               *CLSID
	Lcid                  LCID
	DispidInitialProperty DISPID
}
