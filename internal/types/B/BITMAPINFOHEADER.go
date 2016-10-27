package win

//ref DWORD
//ref LONG
//ref WORD
type BITMAPINFOHEADER struct {
	BiSize          DWORD
	BiWidth         LONG
	BiHeight        LONG
	BiPlanes        WORD
	BiBitCount      WORD
	BiCompression   DWORD
	BiSizeImage     DWORD
	BiXPelsPerMeter LONG
	BiYPelsPerMeter LONG
	BiClrUsed       DWORD
	BiClrImportant  DWORD
}
