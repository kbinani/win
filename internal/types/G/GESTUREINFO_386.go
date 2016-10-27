package win

type GESTUREINFO struct {
	CbSize       uint32 // UINT
	DwFlags      uint32
	DwID         uint32
	HwndTarget   HWND
	PtsLocation  POINTS
	DwInstanceID uint32
	DwSequenceID uint32
	padding1     [4]byte
	UllArguments int64  // ULONGLONG
	CbExtraArgs  uint32 // UINT
	padding2     [4]byte
}
