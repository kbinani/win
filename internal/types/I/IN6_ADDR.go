package win

import "unsafe"

//ref UCHAR
//ref USHORT

type IN6_ADDR struct {
	U IN6_ADDR_U
}

type IN6_ADDR_U struct {
	storage [16]byte
}

func (this *IN6_ADDR_U) GetByte() [16]UCHAR {
	var ret [16]UCHAR
	for i := 0; i < 16; i++ {
		ret[i] = UCHAR(this.storage[i])
	}
	return ret
}

func (this *IN6_ADDR_U) SetByte(v [16]UCHAR) {
	for i := 0; i < 16; i++ {
		this.storage[i] = byte(v[i])
	}
}

func (this *IN6_ADDR_U) GetWord() [8]USHORT {
	var ret [8]USHORT
	for i := 0; i < 8; i++ {
		ret[i] = *(*USHORT)(unsafe.Pointer(&this.storage[i*2]))
	}
	return ret
}

func (this *IN6_ADDR_U) SetWord(v [8]USHORT) {
	for i := 0; i < 8; i++ {
		ptr := (*USHORT)(unsafe.Pointer(&this.storage[i*2]))
		*ptr = v[i]
	}
}
