package win

//import unsafe
type LARGE_INTEGER struct {
	QuadPart int64
}

func (l *LARGE_INTEGER) LowPart() *uint32 {
	return (*uint32)(unsafe.Pointer(&l.QuadPart))
}

func (l *LARGE_INTEGER) HighPart() *int32 {
	return (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&l.QuadPart)) + uintptr(4)))
}
