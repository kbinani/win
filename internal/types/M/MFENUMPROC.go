package win

//ref HDC
//ref HANDLETABLE
//ref METARECORD
//ref LPARAM
type MFENUMPROC func(hdc HDC, lpht *HANDLETABLE, lpMR *METARECORD, nObj int32, param LPARAM) int32
