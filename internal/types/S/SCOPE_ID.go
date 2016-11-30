package win

type SCOPE_ID struct {
	Value ULONG
}

func (this *SCOPE_ID) Zone() ULONG {
	return 0xFFFFFFF & this.Value
}

func (this *SCOPE_ID) Level() ULONG {
	return (0xF000000 & this.Value) >> (32 - 4)
}

func (this *SCOPE_ID) SetZone(v ULONG) {
	this.Value = this.Value | (0x0FFFFFFF & v)
}

func (this *SCOPE_ID) SetLevel(v ULONG) {
	this.Value = this.Value | (0xF0000000 & (v << (32 - 4)))
}
