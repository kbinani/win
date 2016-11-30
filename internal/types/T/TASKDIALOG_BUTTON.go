package win

import "unsafe"

//ref PCWSTR

type TASKDIALOG_BUTTON struct {
	storage [4 * pad3for64_2for32]byte
}

func (this *TASKDIALOG_BUTTON) NButtonID() *int32 {
	return (*int32)(unsafe.Pointer(&this.storage[0]))
}

func (this *TASKDIALOG_BUTTON) PszButtonText() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[4]))
}
