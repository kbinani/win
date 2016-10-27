package win

//ref DWORD
//ref COLORREF
//ref POINT
//ref BOOL
//ref DTT_CALLBACK_PROC
//ref LPARAM
type DTTOPTS struct {
	DwSize              DWORD
	DwFlags             DWORD
	CrText              COLORREF
	CrBorder            COLORREF
	CrShadow            COLORREF
	ITextShadowType     int32
	PtShadowOffset      POINT
	IBorderSize         int32
	IFontPropId         int32
	IColorPropId        int32
	IStateId            int32
	FApplyOverlay       BOOL
	IGlowSize           int32
	PfnDrawTextCallback uintptr // DTT_CALLBACK_PROC
	LParam              LPARAM
}
