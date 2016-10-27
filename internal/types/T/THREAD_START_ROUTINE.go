package win

//ref DWORD
//ref LPVOID
type THREAD_START_ROUTINE func(lpThreadParameter LPVOID) DWORD
