package win

//ref LONG
//ref WCHAR
type LOGFONT struct {
	LfHeight         LONG
	LfWidth          LONG
	LfEscapement     LONG
	LfOrientation    LONG
	LfWeight         LONG
	LfItalic         byte
	LfUnderline      byte
	LfStrikeOut      byte
	LfCharSet        byte
	LfOutPrecision   byte
	LfClipPrecision  byte
	LfQuality        byte
	LfPitchAndFamily byte
	LfFaceName       [LF_FACESIZE]WCHAR
}
