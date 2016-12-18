package win

//ref PBYTE
//ref PVOID
//ref PULONG
//ref DWORD

type PFE_IMPORT_FUNC func(pbData PBYTE, pvCallbackContext PVOID, ulLength PULONG) DWORD
