package win

//ref BITMAPINFOHEADER
//ref RGBQUAD
type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors [1]RGBQUAD
}
