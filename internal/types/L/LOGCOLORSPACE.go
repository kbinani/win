package win

//ref DWORD
//ref LCSCSTYPE
//ref LCSGAMUTMATCH
//ref CIEXYZTRIPLE
//ref WCHAR
type LOGCOLORSPACE struct {
	LcsSignature  DWORD
	LcsVersion    DWORD
	LcsSize       DWORD
	LcsCSType     LCSCSTYPE
	LcsIntent     LCSGAMUTMATCH
	LcsEndpoints  CIEXYZTRIPLE
	LcsGammaRed   DWORD
	LcsGammaGreen DWORD
	LcsGammaBlue  DWORD
	LcsFilename   [MAX_PATH]WCHAR
}
