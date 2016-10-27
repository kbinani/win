package win

type RAWKEYBOARD struct {
	MakeCode         uint16
	Flags            uint16
	Reserved         int16
	VKey             uint16
	Message          uint32
	ExtraInformation uint32
}
