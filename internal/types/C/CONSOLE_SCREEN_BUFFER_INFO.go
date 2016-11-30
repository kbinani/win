package win

//ref COORD
//ref WORD
//ref SMALL_RECT
type CONSOLE_SCREEN_BUFFER_INFO struct {
	DwSize              COORD
	DwCursorPosition    COORD
	WAttributes         WORD
	SrWindow            SMALL_RECT
	DwMaximumWindowSize COORD
}
