package win

//ref ULONG
//ref BYTE
type FD_GLYPHATTR struct {
	CjThis     ULONG
	CGlyphs    ULONG
	IMode      ULONG
	AGlyphAttr [1]BYTE
}
