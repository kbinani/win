package win

//ref DWORD
//ref PCERT_EXTENSION

type CERT_EXTENSIONS struct {
	CExtension  DWORD
	RgExtension PCERT_EXTENSION
}
