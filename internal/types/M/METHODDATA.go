package win

//ref OLECHAR
//ref PARAMDATA
//ref DISPID
//ref UINT
//ref CALLCONV
//ref WORD
//ref VARTYPE
type METHODDATA struct {
	SzName   *OLECHAR
	Ppdata   *PARAMDATA
	Dispid   DISPID
	IMeth    UINT
	Cc       CALLCONV
	CArgs    UINT
	WFlags   WORD
	VtReturn VARTYPE
}
