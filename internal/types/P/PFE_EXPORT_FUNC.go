package win

//ref PBYTE
//ref PVOID
//ref ULONG
//ref DWORD

type PFE_EXPORT_FUNC func(pbData PBYTE, pvCallbackContext PVOID, ulLength ULONG) DWORD
