package win

import "unsafe"

//ref UINT
//ref BYTE
//ref DWORD

type MMTIME struct {
	WType UINT
	U     MMTIME_U
}

type MMTIME_U struct {
	storage [8]byte
}

type MMTIME_Smpte struct {
	Hour  BYTE
	Min   BYTE
	Sec   BYTE
	Frame BYTE
	Fps   BYTE
	Dummy BYTE
	Pad   [2]BYTE
}

type MMTIME_Midi struct {
	Songptrpos DWORD
}

func (this *MMTIME_U) Ms() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}

func (this *MMTIME_U) Sample() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}

func (this *MMTIME_U) Cb() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}

func (this *MMTIME_U) Ticks() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}

func (this *MMTIME_U) Smpte() *MMTIME_Smpte {
	return (*MMTIME_Smpte)(unsafe.Pointer(&this.storage[0]))
}

func (this *MMTIME_U) Midi() *MMTIME_Midi {
	return (*MMTIME_Midi)(unsafe.Pointer(&this.storage[0]))
}
