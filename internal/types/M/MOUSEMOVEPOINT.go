package win

type MOUSEMOVEPOINT struct {
	X           int32
	Y           int32
	Time        uint32
	DwExtraInfo uintptr // ULONG_PTR
}
