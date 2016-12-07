package win

//ref XLAT_SIDE

type FULL_PTR_XLAT_TABLES struct {
	RefIdToPointer uintptr
	PointerToRefId uintptr
	NextRefId      uint32
	XlatSide       XLAT_SIDE
}
