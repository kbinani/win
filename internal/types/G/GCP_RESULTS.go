package win

//ref DWORD
//ref LPWSTR
//ref UINT
//ref LPSTR
type GCP_RESULTS struct {
	LStructSize DWORD
	LpOutString LPWSTR
	LpOrder     *UINT
	LpDx        *int32
	LpCaretPos  *int32
	LpClass     LPSTR
	LpGlyphs    LPWSTR
	NGlyphs     UINT
	NMaxFit     int32
}
