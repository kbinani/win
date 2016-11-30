package win

import "unsafe"

//ref DWORD
//ref FOURCC
//ref MMIOPROC
//ref UINT
//ref HTASK
//ref LONG
//ref HPSTR
//ref HMMIO

type MMIOINFO struct {
	storage [44 + ptrsize*7]byte
}

func (this *MMIOINFO) DwFlags() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}
func (this *MMIOINFO) FccIOProc() *FOURCC {
	return (*FOURCC)(unsafe.Pointer(&this.storage[4]))
}

// LPMMIOPROC
func (this *MMIOINFO) PIOProc() *uintptr {
	return (*uintptr)(unsafe.Pointer(&this.storage[8]))
}

func (this *MMIOINFO) WErrorRet() *UINT {
	return (*UINT)(unsafe.Pointer(&this.storage[8+ptrsize]))
}

func (this *MMIOINFO) Htask() *HTASK {
	return (*HTASK)(unsafe.Pointer(&this.storage[12+ptrsize]))
}

func (this *MMIOINFO) CchBuffer() *LONG {
	return (*LONG)(unsafe.Pointer(&this.storage[12+ptrsize*2]))
}

func (this *MMIOINFO) PchBuffer() *HPSTR {
	return (*HPSTR)(unsafe.Pointer(&this.storage[16+ptrsize*2]))
}

func (this *MMIOINFO) PchNext() *HPSTR {
	return (*HPSTR)(unsafe.Pointer(&this.storage[16+ptrsize*3]))
}

func (this *MMIOINFO) PchEndRead() *HPSTR {
	return (*HPSTR)(unsafe.Pointer(&this.storage[16+ptrsize*4]))
}

func (this *MMIOINFO) PchEndWrite() *HPSTR {
	return (*HPSTR)(unsafe.Pointer(&this.storage[16+ptrsize*5]))
}

func (this *MMIOINFO) LBufOffset() *LONG {
	return (*LONG)(unsafe.Pointer(&this.storage[16+ptrsize*6]))
}

func (this *MMIOINFO) LDiskOffset() *LONG {
	return (*LONG)(unsafe.Pointer(&this.storage[20+ptrsize*6]))
}

func (this *MMIOINFO) AdwInfo() *[3]DWORD {
	return (*[3]DWORD)(unsafe.Pointer(&this.storage[24+ptrsize*6]))
}

func (this *MMIOINFO) DwReserved1() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[36+ptrsize*6]))
}

func (this *MMIOINFO) DwReserved2() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[40+ptrsize*6]))
}

func (this *MMIOINFO) Hmmio() *HMMIO {
	return (*HMMIO)(unsafe.Pointer(&this.storage[44+ptrsize*6]))
}
