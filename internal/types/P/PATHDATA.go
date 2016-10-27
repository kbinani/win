package win

//ref FLONG
//ref ULONG
//ref POINTFIX
type PATHDATA struct {
	Flags FLONG
	Count ULONG
	Pptfx *POINTFIX
}
