package win

//ref BOOL
type GdiplusStartupInput struct {
	GdiplusVersion           uint32
	DebugEventCallback       uintptr // DebugEventProc
	SuppressBackgroundThread BOOL
	SuppressExternalCodecs   BOOL
}
