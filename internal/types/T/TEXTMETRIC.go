package win

//ref LONG
//ref WCHAR
//ref BYTE
type TEXTMETRIC struct {
	TmHeight           LONG
	TmAscent           LONG
	TmDescent          LONG
	TmInternalLeading  LONG
	TmExternalLeading  LONG
	TmAveCharWidth     LONG
	TmMaxCharWidth     LONG
	TmWeight           LONG
	TmOverhang         LONG
	TmDigitizedAspectX LONG
	TmDigitizedAspectY LONG
	TmFirstChar        WCHAR
	TmLastChar         WCHAR
	TmDefaultChar      WCHAR
	TmBreakChar        WCHAR
	TmItalic           BYTE
	TmUnderlined       BYTE
	TmStruckOut        BYTE
	TmPitchAndFamily   BYTE
	TmCharSet          BYTE
}
