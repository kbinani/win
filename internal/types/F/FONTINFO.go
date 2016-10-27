package win

//ref ULONG
//ref FLONG
type FONTINFO struct {
	CjThis           ULONG
	FlCaps           FLONG
	CGlyphsSupported ULONG
	CjMaxGlyph1      ULONG
	CjMaxGlyph4      ULONG
	CjMaxGlyph8      ULONG
	CjMaxGlyph32     ULONG
}
