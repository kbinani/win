package win

//ref HWCT
//ref DWORD_PTR
//ref DWORD
//ref LPDWORD
//ref PWAITCHAIN_NODE_INFO
//ref LPBOOL

type PWAITCHAINCALLBACK func(WctHandle HWCT, Context DWORD_PTR, CallbackStatus DWORD, NodeCount LPDWORD, NodeInfoArray PWAITCHAIN_NODE_INFO, IsCycle LPBOOL)
