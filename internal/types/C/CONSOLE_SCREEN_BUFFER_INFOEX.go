package win

//ref ULONG
//ref COORD
//ref WORD
//ref SMALL_RECT
//ref COLORREF
type CONSOLE_SCREEN_BUFFER_INFOEX struct {
	CbSize               ULONG
	DwSize               COORD
	DwCursorPosition     COORD
	WAttributes          WORD
	SrWindow             SMALL_RECT
	DwMaximumWindowSize  COORD
	WPopupAttributes     WORD
	BFullscreenSupported BOOL
	ColorTable           [16]COLORREF
}
