package win

//ref VARIANTARG
//ref DISPID
//ref UINT
type DISPPARAMS struct {
	Rgvarg            *VARIANTARG
	RgdispidNamedArgs *DISPID
	CArgs             UINT
	CNamedArgs        UINT
}
