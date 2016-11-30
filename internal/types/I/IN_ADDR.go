package win

import "unsafe"

type IN_ADDR struct {
	S_un IN_ADDR_S_un
}

type IN_ADDR_S_un struct {
	storage ULONG
}

type IN_ADDR_S_un_b struct {
	s_b1, s_b2, s_b3, s_b4 UCHAR
}

type IN_ADDR_S_un_w struct {
	s_w1, s_w2 USHORT
}

func (this *IN_ADDR_S_un) S_un_b() *IN_ADDR_S_un_b {
	return (*IN_ADDR_S_un_b)(unsafe.Pointer(&this.storage))
}

func (this *IN_ADDR_S_un) S_un_w() *IN_ADDR_S_un_w {
	return (*IN_ADDR_S_un_w)(unsafe.Pointer(&this.storage))
}

func (this *IN_ADDR_S_un) S_addr() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.storage))
}

func (this *IN_ADDR) S_addr() ULONG {
	return *this.S_un.S_addr()
}

func (this *IN_ADDR) S_host() UCHAR {
	return this.S_un.S_un_b().s_b2
}

func (this *IN_ADDR) S_net() UCHAR {
	return this.S_un.S_un_b().s_b1
}

func (this *IN_ADDR) S_imp() USHORT {
	return this.S_un.S_un_w().s_w2
}

func (this *IN_ADDR) S_impno() UCHAR {
	return this.S_un.S_un_b().s_b4
}

func (this *IN_ADDR) S_lh() UCHAR {
	return this.S_un.S_un_b().s_b3
}
