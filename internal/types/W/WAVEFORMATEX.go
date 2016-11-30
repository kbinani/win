package win

import "unsafe"

//ref WORD
//reF DWORD

type WAVEFORMATEX struct {
	storage [18]byte
}

func (this *WAVEFORMATEX) WFormatTag() *WORD {
	return (*WORD)(unsafe.Pointer(&this.storage[0]))
}

func (this *WAVEFORMATEX) NChannels() *WORD {
	return (*WORD)(unsafe.Pointer(&this.storage[2]))
}

func (this *WAVEFORMATEX) NSamplesPerSec() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[4]))
}

func (this *WAVEFORMATEX) NAvgBytesPerSec() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[8]))
}

func (this *WAVEFORMATEX) NBlockAlign() *WORD {
	return (*WORD)(unsafe.Pointer(&this.storage[12]))
}

func (this *WAVEFORMATEX) WBitsPerSample() *WORD {
	return (*WORD)(unsafe.Pointer(&this.storage[14]))
}

func (this *WAVEFORMATEX) CbSize() *WORD {
	return (*WORD)(unsafe.Pointer(&this.storage[16]))
}
