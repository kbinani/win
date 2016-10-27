package win

//ref BOOL
//ref HBITMAP
type ICONINFO struct {
	FIcon    BOOL
	XHotspot uint32
	YHotspot uint32
	HbmMask  HBITMAP
	HbmColor HBITMAP
}
