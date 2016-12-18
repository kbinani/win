package win

//ref LPWSTR
//ref DWORD
//ref PPROG_INVOKE_SETTING
//ref PVOID
//ref BOOL

type FN_PROGRESS func(pObjectName LPWSTR, Status DWORD, pInvokeSetting PPROG_INVOKE_SETTING, Args PVOID, SecuritySet BOOL)
