package win

//ref UINT
//ref LPOLESTR
//ref CY
//ref SHORT
//ref BOOL
type FONTDESC struct {
	CbSizeofstruct UINT
	LpstrName      LPOLESTR
	CySize         CY
	SWeight        SHORT
	SCharset       SHORT
	FItalic        BOOL
	FUnderline     BOOL
	FStrikethrough BOOL
}
