package win

type RAWHID struct {
	DwSizeHid uint32
	DwCount   uint32
	BRawData  [1]byte
}
