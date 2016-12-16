package win

//ref DWORD
//ref LPSTR

type CERT_ENHKEY_USAGE struct {
	CUsageIdentifier     DWORD
	RgpszUsageIdentifier *LPSTR
}
