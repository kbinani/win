package win

//ref GUID
//ref DWORD
//ref HRESULT

type PCOGETACTIVATIONSTATE func(unnamed0 GUID, unnamed1 DWORD, unnamed2 *DWORD) HRESULT
