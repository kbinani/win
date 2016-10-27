package win

//ref WORD
//ref BSTR
//ref DWORD
//ref SCODE
//ref PVOID
type EXCEPINFO struct {
	WCode             WORD
	WReserved         WORD
	BstrSource        BSTR
	BstrDescription   BSTR
	BstrHelpFile      BSTR
	DwHelpContext     DWORD
	PvReserved        PVOID
	PfnDeferredFillIn uintptr
	Scode             SCODE
}

func (this *EXCEPINFO) DeferredFillIn() func(unnamed0 *EXCEPINFO) HRESULT {
	return func(unnamed0 *EXCEPINFO) HRESULT {
		ret := syscall3(this.PfnDeferredFillIn, 1,
			uintptr(unsafe.Pointer(unnamed0)),
			0,
			0)
		return HRESULT(ret)
	}
}
