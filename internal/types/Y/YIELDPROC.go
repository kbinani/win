package win

//ref MCIDEVICEID
//ref DWORD
//ref UINT

type YIELDPROC func(mciId MCIDEVICEID, dwYieldData DWORD) UINT
