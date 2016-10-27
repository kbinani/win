package win

type NET_LUID_LH struct {
	Value ULONG64
}

func (this *NET_LUID_LH) Reserved() ULONG64 {
	v := this.Value
	return (v & 0xFFFFFF0000000000) >> (64 - 24)
}
func (this *NET_LUID_LH) NetLuidIndex() ULONG64 {
	v := this.Value
	return (v & 0xFFFFFF0000) >> 16
}
func (this *NET_LUID_LH) IfType() ULONG64 {
	v := this.Value
	return v & 0xFFFF
}
