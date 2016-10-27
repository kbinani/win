package win

//ref DWORD
//ref LONG
//ref DWORD_PTR
//ref LPWSTR
//ref PDH_DATA_ITEM_PATH_ELEMENTS
//ref PDH_COUNTER_PATH_ELEMENTS
type PDH_COUNTER_INFO struct {
	DwLength        DWORD
	DwType          DWORD
	CVersion        DWORD
	CStatus         DWORD
	LScale          LONG
	LDefaultScale   LONG
	DwUserData      DWORD_PTR
	DwQueryUserData DWORD_PTR
	SzFullPath      LPWSTR
	union1          [44]byte
	SzExplainText   LPWSTR
	DataBuffer      [1]DWORD
}

func (this *PDH_COUNTER_INFO) DataItemPath() *PDH_DATA_ITEM_PATH_ELEMENTS {
	return (*PDH_DATA_ITEM_PATH_ELEMENTS)(unsafe.Pointer(&this.union1[0]))
}
func (this *PDH_COUNTER_INFO) CounterPath() *PDH_COUNTER_PATH_ELEMENTS {
	return (*PDH_COUNTER_PATH_ELEMENTS)(unsafe.Pointer(&this.union1[0]))
}
func (this *PDH_COUNTER_INFO) SzMachineName() *LPWSTR {
	return (*LPWSTR)(unsafe.Pointer(&this.union1[0]))
}
func (this *PDH_COUNTER_INFO) SzObjectName() *LPWSTR {
	var ptr LPWSTR
	return (*LPWSTR)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1[0])) + unsafe.Sizeof(ptr)))
}
func (this *PDH_COUNTER_INFO) SzInstanceName() *LPWSTR {
	var ptr LPWSTR
	return (*LPWSTR)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1[0])) + unsafe.Sizeof(ptr)*2))
}
func (this *PDH_COUNTER_INFO) SzParentInstance() *LPWSTR {
	var ptr LPWSTR
	return (*LPWSTR)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1[0])) + unsafe.Sizeof(ptr)*3))
}
func (this *PDH_COUNTER_INFO) DwInstanceIndex() *DWORD {
	var ptr LPWSTR
	return (*DWORD)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1[0])) + unsafe.Sizeof(ptr)*4))
}
func (this *PDH_COUNTER_INFO) SzCounterName() *LPWSTR {
	var ptr LPWSTR
	return (*LPWSTR)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1[0])) + unsafe.Sizeof(ptr)*4 + uintptr(4)))
}
