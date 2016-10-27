package win

//ref WORD
//ref SHORT
type COLORADJUSTMENT struct {
	CaSize            WORD
	CaFlags           WORD
	CaIlluminantIndex WORD
	CaRedGamma        WORD
	CaGreenGamma      WORD
	CaBlueGamma       WORD
	CaReferenceBlack  WORD
	CaReferenceWhite  WORD
	CaContrast        SHORT
	CaBrightness      SHORT
	CaColorfulness    SHORT
	CaRedGreenTint    SHORT
}
