package win

//ref DWORD
//ref LPWSAOVERLAPPED
type LPWSAOVERLAPPED_COMPLETION_ROUTINE func(dwError DWORD, cbTransferred DWORD, lpOverlapped LPWSAOVERLAPPED, dwFlags DWORD)
