package win

//ref GENERIC_BINDING_ROUTINE
//ref GENERIC_UNBIND_ROUTINE
type GENERIC_BINDING_INFO struct {
	PObj      uintptr
	Size      uint32
	PfnBind   uintptr // GENERIC_BINDING_ROUTINE
	PfnUnbind uintptr // GENERIC_UNBIND_ROUTINE
}
