package win

type RAWMOUSE struct {
	UsFlags            uint16
	padding            [2]byte
	UsButtonFlags      uint16
	UsButtonData       uint16
	UlRawButtons       uint32
	LLastX             int32
	LLastY             int32
	UlExtraInformation uint32
}
