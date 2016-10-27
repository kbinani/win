package win

//ref GLYPHBITS
//ref PATHOBJ
//import unsafe
type GLYPHDEF struct {
	p uintptr
}

func (this *GLYPHDEF) Pgb() *GLYPHBITS {
	return (*GLYPHBITS)(unsafe.Pointer(this.p))
}
func (this *GLYPHDEF) Ppo() *PATHOBJ {
	return (*PATHOBJ)(unsafe.Pointer(this.p))
}
