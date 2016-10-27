package win

//ref FLONG
//ref ULONG
//ref FLOAT_LONG
//ref PFLOAT_LONG
type LINEATTRS struct {
	Fl           FLONG
	IJoin        ULONG
	IEndCap      ULONG
	ElWidth      FLOAT_LONG
	EMiterLimit  FLOATL
	Cstyle       ULONG
	Pstyle       PFLOAT_LONG
	ElStyleState FLOAT_LONG
}
