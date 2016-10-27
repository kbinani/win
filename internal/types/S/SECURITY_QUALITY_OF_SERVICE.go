package win

//ref SECURITY_IMPERSONATION_LEVEL
//ref SECURITY_CONTEXT_TRACKING_MODE
//ref BOOLEAN
//ref DWORD
//import unsafe
type SECURITY_QUALITY_OF_SERVICE struct {
	storage [12]byte
}

func (this *SECURITY_QUALITY_OF_SERVICE) Length() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}
func (this *SECURITY_QUALITY_OF_SERVICE) ImpersonationLevel() *SECURITY_IMPERSONATION_LEVEL {
	return (*SECURITY_IMPERSONATION_LEVEL)(unsafe.Pointer(&this.storage[4]))
}
func (this *SECURITY_QUALITY_OF_SERVICE) ContextTrackingMode() *SECURITY_CONTEXT_TRACKING_MODE {
	return (*SECURITY_CONTEXT_TRACKING_MODE)(unsafe.Pointer(&this.storage[8]))
}
func (this *SECURITY_QUALITY_OF_SERVICE) EffectiveOnly() *BOOLEAN {
	return (*BOOLEAN)(unsafe.Pointer(&this.storage[9]))
}
