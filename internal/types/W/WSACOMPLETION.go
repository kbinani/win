package win

//ref WSACOMPLETIONTYPE
//ref HWND
//ref UINT
//ref WPARAM
//ref LPWSAOVERLAPPED
//ref LPWSAOVERLAPPED_COMPLETION_ROUTINE
//ref HANDLE
//ref ULONG_PTR
import "unsafe"

type WSACOMPLETION_Parameters_WindowMessage struct {
	HWnd    HWND
	UMsg    UINT
	Context WPARAM
}

type WSACOMPLETION_Parameters_Event struct {
	LpOverlapped LPWSAOVERLAPPED
}

type WSACOMPLETION_Parameters_Apc struct {
	LpOverlapped       LPWSAOVERLAPPED
	LpfnCompletionProc uintptr // LPWSAOVERLAPPED_COMPLETION_ROUTINE
}

type WSACOMPLETION_Parameters_Port struct {
	LpOverlapped LPWSAOVERLAPPED
	HPort        HANDLE
	Key          ULONG_PTR
}

func (this *WSACOMPLETION_Parameters) WindowMessage() *WSACOMPLETION_Parameters_WindowMessage {
	return (*WSACOMPLETION_Parameters_WindowMessage)(unsafe.Pointer(this))
}

func (this *WSACOMPLETION_Parameters) Event() *WSACOMPLETION_Parameters_Event {
	return (*WSACOMPLETION_Parameters_Event)(unsafe.Pointer(this))
}

func (this *WSACOMPLETION_Parameters) Apc() *WSACOMPLETION_Parameters_Apc {
	return (*WSACOMPLETION_Parameters_Apc)(unsafe.Pointer(this))
}

func (this *WSACOMPLETION_Parameters) Port() *WSACOMPLETION_Parameters_Port {
	return (*WSACOMPLETION_Parameters_Port)(unsafe.Pointer(this))
}
