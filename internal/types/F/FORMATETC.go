package win

//ref CLIPFORMAT
//ref DVTARGETDEVICE
//ref DWORD
//ref LONG

type FORMATETC struct {
	CfFormat CLIPFORMAT
	Ptd      *DVTARGETDEVICE
	DwAspect DWORD
	Lindex   LONG
	Tymed    DWORD
}
