package win

//ref PCWSTR
//import unsafe
type TASKDIALOG_BUTTON struct {
	storage [12]byte
}

func (this *TASKDIALOG_BUTTON) NButtonID() *int32 {
	return (*int32)(unsafe.Pointer(&this.storage[0]))
}
func (this *TASKDIALOG_BUTTON) PszButtonText() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[4]))
}
