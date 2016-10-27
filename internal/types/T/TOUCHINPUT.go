package win

//ref HANDLE
type TOUCHINPUT struct {
	X           int32 // LONG
	Y           int32 // LONG
	HSource     HANDLE
	DwID        uint32
	DwFlags     uint32
	DwMask      uint32
	DwTime      uint32
	DwExtraInfo uintptr // ULONG_PTR
	CxContact   uint32
	CyContact   uint32
}
