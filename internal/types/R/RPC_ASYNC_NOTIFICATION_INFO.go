package win

//ref HANDLE
//ref PFN_RPCNOTIFICATION_ROUTINE
//ref DWORD
//ref DWORD_PTR
//ref LPOVERLAPPED
//ref HWND
//ref UINT
//ref PFN_RPCNOTIFICATION_ROUTINE

type RPC_ASYNC_NOTIFICATION_INFO struct {
	storage1 [4 * ptrsize]byte
}

func (this *RPC_ASYNC_NOTIFICATION_INFO) GetAPC() RPC_ASYNC_NOTIFICATION_INFO_APC {
	return *(*RPC_ASYNC_NOTIFICATION_INFO_APC)(unsafe.Pointer(this))
}

func (this *RPC_ASYNC_NOTIFICATION_INFO) SetAPC(v RPC_ASYNC_NOTIFICATION_INFO_APC) {
	*(*RPC_ASYNC_NOTIFICATION_INFO_APC)(unsafe.Pointer(this)) = v
}

func (this *RPC_ASYNC_NOTIFICATION_INFO) GetIOC() RPC_ASYNC_NOTIFICATION_INFO_IOC {
	return *(*RPC_ASYNC_NOTIFICATION_INFO_IOC)(unsafe.Pointer(this))
}

func (this *RPC_ASYNC_NOTIFICATION_INFO) SetIOC(v RPC_ASYNC_NOTIFICATION_INFO_IOC) {
	*(*RPC_ASYNC_NOTIFICATION_INFO_IOC)(unsafe.Pointer(this)) = v
}

func (this *RPC_ASYNC_NOTIFICATION_INFO) GetHWND() RPC_ASYNC_NOTIFICATION_INFO_HWND {
	return *(*RPC_ASYNC_NOTIFICATION_INFO_HWND)(unsafe.Pointer(this))
}

func (this *RPC_ASYNC_NOTIFICATION_INFO) SetHWND(v RPC_ASYNC_NOTIFICATION_INFO_HWND) {
	*(*RPC_ASYNC_NOTIFICATION_INFO_HWND)(unsafe.Pointer(this)) = v
}

func (this *RPC_ASYNC_NOTIFICATION_INFO) GetHEvent() HANDLE {
	return *(*HANDLE)(unsafe.Pointer(this))
}

func (this *RPC_ASYNC_NOTIFICATION_INFO) SetHEvent(v HANDLE) {
	*(*HANDLE)(unsafe.Pointer(this)) = v
}

func (this *RPC_ASYNC_NOTIFICATION_INFO) GetNotificationRoutine() uintptr {
	return *(*uintptr)(unsafe.Pointer(this))
}

func (this *RPC_ASYNC_NOTIFICATION_INFO) SetNotificationRoutine(v uintptr) {
	*(*uintptr)(unsafe.Pointer(this)) = v
}

type RPC_ASYNC_NOTIFICATION_INFO_APC struct {
	NotificationRoutine PFN_RPCNOTIFICATION_ROUTINE
	HThread             HANDLE
}

type RPC_ASYNC_NOTIFICATION_INFO_IOC struct {
	HIOPort                    HANDLE
	DwNumberOfBytesTransferred DWORD
	DwCompletionKey            DWORD_PTR
	LpOverlapped               LPOVERLAPPED
}

type RPC_ASYNC_NOTIFICATION_INFO_HWND struct {
	HWnd HWND
	Msg  UINT
}
