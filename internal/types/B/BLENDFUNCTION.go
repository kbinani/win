package win

type BLENDFUNCTION struct {
	BlendOp             byte
	BlendFlags          byte
	SourceConstantAlpha byte
	AlphaFormat         byte
}

func getUintptrFromBLENDFUNCTION(v BLENDFUNCTION) uintptr {
	ret := (uint32(v.BlendOp) << 12) | (uint32(v.BlendFlags) << 8) | (uint32(v.SourceConstantAlpha) << 4) | uint32(v.AlphaFormat)
	return uintptr(ret)
}
