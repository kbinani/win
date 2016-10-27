package win

type MOUSEINPUT struct {
	Dx          int32 // LONG
	Dy          int32 // LONG
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr // ULONG_PTR
}
