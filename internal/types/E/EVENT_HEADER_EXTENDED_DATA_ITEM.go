package win

type EVENT_HEADER_EXTENDED_DATA_ITEM struct {
	Reserved1 USHORT
	ExtType   USHORT
	storage1  USHORT
	DataSize  USHORT
	DataPtr   ULONGLONG
}

func (this *EVENT_HEADER_EXTENDED_DATA_ITEM) GetLinkage() USHORT {
	return USHORT(this.storage1 >> 15)
}

func (this *EVENT_HEADER_EXTENDED_DATA_ITEM) GetReserved2() USHORT {
	return USHORT(this.storage1 & 0x7fff)
}
