package win

type GESTUREINFO struct {
	CbSize       uint32 // UINT
	DwFlags      uint32
	DwID         uint32
	HwndTarget   HWND
	PtsLocation  POINTS
	DwInstanceID uint32
	DwSequenceID uint32
	UllArguments int64  // ULONGLONG
	CbExtraArgs  uint32 // UINT
}
