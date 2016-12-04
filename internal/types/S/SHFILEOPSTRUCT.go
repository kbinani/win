package win

import "unsafe"

//ref HWND
//ref UINT
//ref PCZZWSTR
//ref FILEOP_FLAGS
//ref BOOL
//ref LPVOID
//ref PCWSTR

type SHFILEOPSTRUCT struct {
	storage [ptrsize*5 + 10 + pad6for64_0for32]byte
}

func (this *SHFILEOPSTRUCT) Hwnd() *HWND {
	return (*HWND)(unsafe.Pointer(&this.storage[0]))
}

func (this *SHFILEOPSTRUCT) WFunc() *UINT {
	return (*UINT)(unsafe.Pointer(&this.storage[ptrsize]))
}

func (this *SHFILEOPSTRUCT) PFrom() *PCZZWSTR {
	return (*PCZZWSTR)(unsafe.Pointer(&this.storage[ptrsize+4+pad4for64_0for32]))
}

func (this *SHFILEOPSTRUCT) PTo() *PCZZWSTR {
	return (*PCZZWSTR)(unsafe.Pointer(&this.storage[ptrsize*2+4+pad4for64_0for32]))
}

func (this *SHFILEOPSTRUCT) FFlags() *FILEOP_FLAGS {
	return (*FILEOP_FLAGS)(unsafe.Pointer(&this.storage[ptrsize*3+4+pad4for64_0for32]))
}

func (this *SHFILEOPSTRUCT) FAnyOperationsAborted() *BOOL {
	return (*BOOL)(unsafe.Pointer(&this.storage[ptrsize*3+6+pad6for64_0for32]))
}

func (this *SHFILEOPSTRUCT) HNameMappings() *LPVOID {
	return (*LPVOID)(unsafe.Pointer(&this.storage[ptrsize*3+10+pad6for64_0for32]))
}

func (this *SHFILEOPSTRUCT) LpszProgressTitle() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[ptrsize*4+10+pad6for64_0for32]))
}
