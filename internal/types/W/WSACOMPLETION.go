package win

//ref WSACOMPLETIONTYPE
//ref HWND
//ref UINT
//ref WPARAM
//ref LPWSAOVERLAPPED
//ref LPWSAOVERLAPPED_COMPLETION_ROUTINE
//ref HANDLE
//ref ULONG_PTR
//import unsafe
type WSACOMPLETION struct {
	storage [4]uintptr
}

type WSACOMPLETION_WindowMessage struct {
	HWnd    HWND
	UMsg    UINT
	Context WPARAM
}

type WSACOMPLETION_Event struct {
	LpOverlapped LPWSAOVERLAPPED
}

type WSACOMPLETION_Apc struct {
	LpOverlapped       LPWSAOVERLAPPED
	LpfnCompletionProc uintptr // LPWSAOVERLAPPED_COMPLETION_ROUTINE
}

type WSACOMPLETION_Port struct {
	LpOverlapped LPWSAOVERLAPPED
	HPort        HANDLE
	Key          ULONG_PTR
}

func (this *WSACOMPLETION) Type() *WSACOMPLETIONTYPE {
	return (*WSACOMPLETIONTYPE)(unsafe.Pointer(&this))
}

func (this *WSACOMPLETION) WindowMessage() *WSACOMPLETION_WindowMessage {
	return (*WSACOMPLETION_WindowMessage)(unsafe.Pointer(uintptr(unsafe.Pointer(&this)) + uintptr(4)))
}

func (this *WSACOMPLETION) Event() *WSACOMPLETION_Event {
	return (*WSACOMPLETION_Event)(unsafe.Pointer(uintptr(unsafe.Pointer(&this)) + uintptr(4)))
}

func (this *WSACOMPLETION) Apc() *WSACOMPLETION_Apc {
	return (*WSACOMPLETION_Apc)(unsafe.Pointer(uintptr(unsafe.Pointer(&this)) + uintptr(4)))
}

func (this *WSACOMPLETION) Port() *WSACOMPLETION_Port {
	return (*WSACOMPLETION_Port)(unsafe.Pointer(uintptr(unsafe.Pointer(&this)) + uintptr(4)))
}
