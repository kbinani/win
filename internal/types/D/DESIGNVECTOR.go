package win

//ref DWORD
//ref LONG
type DESIGNVECTOR struct {
	DvReserved DWORD
	DvNumAxes  DWORD
	DvValues   [MM_MAX_NUMAXES]LONG
}
