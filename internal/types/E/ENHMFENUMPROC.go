package win

//ref HDC
//ref ENHMETARECORD
//ref LPARAM
//ref HANDLETABLE
type ENHMFENUMPROC func(hdc HDC, lpht *HANDLETABLE, lpmr /*const*/ *ENHMETARECORD, nHandles int32, data LPARAM) int32
