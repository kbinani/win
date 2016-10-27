package win

//ref DWORD
//ref WCRANGE
type GLYPHSET struct {
	CbThis           DWORD
	FlAccel          DWORD
	CGlyphsSupported DWORD
	CRanges          DWORD
	Ranges           [1]WCRANGE
}
