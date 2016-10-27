package win

//ref LONG
//ref WORD
//ref LPVOID
type BITMAP struct {
	BmType       LONG
	BmWidth      LONG
	BmHeight     LONG
	BmWidthBytes LONG
	BmPlanes     WORD
	BmBitsPixel  WORD
	BmBits       LPVOID
}
