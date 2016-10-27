package win

//ref BOOL
type SECURITY_ATTRIBUTES struct {
	NLength              uint32
	LpSecurityDescriptor uintptr
	BInheritHandle       BOOL
}
