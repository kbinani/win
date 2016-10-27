package win

//ref WORD
//ref PALETTEENTRY
type LOGPALETTE struct {
	PalVersion    WORD
	PalNumEntries WORD
	PalPalEntry   [1]PALETTEENTRY
}
