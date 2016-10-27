package win

//ref DWORD
//ref HIMAGELIST
//ref HDC
//ref COLORREF
//ref UINT
type IMAGELISTDRAWPARAMS struct {
	CbSize   DWORD
	Himl     HIMAGELIST
	I        int32
	HdcDst   HDC
	X        int32
	Y        int32
	Cx       int32
	Cy       int32
	XBitmap  int32
	YBitmap  int32
	RgbBk    COLORREF
	RgbFg    COLORREF
	FStyle   UINT
	DwRop    DWORD
	FState   DWORD
	Frame    DWORD
	CrEffect COLORREF
}
