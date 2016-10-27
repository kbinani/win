package win

//ref LOGFONT
//ref WCHAR
type ENUMLOGFONTEX struct {
	ElfLogFont  LOGFONT
	ElfFullName [LF_FULLFACESIZE]WCHAR
	ElfStyle    [LF_FACESIZE]WCHAR
	ElfScript   [LF_FACESIZE]WCHAR
}
