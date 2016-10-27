package win

//ref USHORT
//ref BYTE
//import unsafe
type SHITEMID struct {
	storage [3]byte
}

func (this *SHITEMID) Cb() *USHORT {
	return (*USHORT)(unsafe.Pointer(&this.storage[0]))
}
func (this *SHITEMID) AbID() *BYTE {
	return (*BYTE)(unsafe.Pointer(&this.storage[2]))
}
