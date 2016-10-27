package win

//ref POINTFIX
//ref LONG
//ref ULONG
//ref RUN
type CLIPLINE struct {
	PtfxA       POINTFIX
	PtfxB       POINTFIX
	LStyleState LONG
	C           ULONG
	Arun        [1]RUN
}
