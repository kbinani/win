package win

//ref HDC
//ref POINT
//ref SIZE
//ref COLORREF
type UPDATELAYEREDWINDOWINFO struct {
	CbSize   uint32
	HdcDst   HDC
	PptDst   *POINT // const POINT*
	Psize    *SIZE  // const SIZE*
	HdcSrc   HDC
	PptSrc   *POINT // const POINT*
	CrKey    COLORREF
	Pblend   uintptr // const BLENDFUNCTION*
	DwFlags  uint32
	PrcDirty *RECT // const RECT*
}
