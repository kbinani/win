package win

//ref LPWSABUF
//ref LPQOS
//ref GROUP
//ref DWORD_PTR
type LPCONDITIONPROC func(lpCallerId LPWSABUF, lpCallerData LPWSABUF, lpSQOS LPQOS, lpGQOS LPQOS, lpCalleeId LPWSABUF, lpCalleeData LPWSABUF, g *GROUP, dwCallbackData DWORD_PTR) int32
