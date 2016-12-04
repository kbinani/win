package win

//ref DWORD
//ref LPSTR

type CTL_USAGE struct {
	CUsageIdentifier     DWORD
	RgpszUsageIdentifier *LPSTR
}
