package win

//ref UINT
//ref POINT
type GLYPHMETRICS struct {
	GmBlackBoxX     UINT
	GmBlackBoxY     UINT
	GmptGlyphOrigin POINT
	GmCellIncX      int16
	GmCellIncY      int16
}
