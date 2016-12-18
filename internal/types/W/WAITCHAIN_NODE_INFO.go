package win

//ref WCT_OBJECT_TYPE
//ref WCT_OBJECT_STATUS
//ref WCHAR
//ref LARGE_INTEGER
//ref BOOL
//ref DWORD

type WAITCHAIN_NODE_INFO struct {
	ObjectType   WCT_OBJECT_TYPE
	ObjectStatus WCT_OBJECT_STATUS
	union1       [272]byte
}

type WAITCHAIN_NODE_INFO_LockObject struct {
	ObjectName [WCT_OBJNAME_LENGTH]WCHAR
	Timeout    LARGE_INTEGER
	Alertable  BOOL
}

type WAITCHAIN_NODE_INFO_ThreadObject struct {
	ProcessId       DWORD
	ThreadId        DWORD
	WaitTime        DWORD
	ContextSwitches DWORD
}

func (this *WAITCHAIN_NODE_INFO) GetLockObject() WAITCHAIN_NODE_INFO_LockObject {
	return *(*WAITCHAIN_NODE_INFO_LockObject)(unsafe.Pointer(&this.union1[0]))
}

func (this *WAITCHAIN_NODE_INFO) SetLockObject(v WAITCHAIN_NODE_INFO_LockObject) {
	*(*WAITCHAIN_NODE_INFO_LockObject)(unsafe.Pointer(&this.union1[0])) = v
}

func (this *WAITCHAIN_NODE_INFO) GetThreadObject() WAITCHAIN_NODE_INFO_ThreadObject {
	return *(*WAITCHAIN_NODE_INFO_ThreadObject)(unsafe.Pointer(&this.union1[0]))
}

func (this *WAITCHAIN_NODE_INFO) SetThreadObject(v WAITCHAIN_NODE_INFO_ThreadObject) {
	*(*WAITCHAIN_NODE_INFO_ThreadObject)(unsafe.Pointer(&this.union1[0])) = v
}
