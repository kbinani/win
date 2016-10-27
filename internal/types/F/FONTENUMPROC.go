package win

//ref LOGFONT
//ref TEXTMETRIC
//ref DWORD
//ref LPARAM
type FONTENUMPROC func(unnamed0 /*const*/ *LOGFONT, unnamed1 /*const*/ *TEXTMETRIC, unnamed2 DWORD, unnamed3 LPARAM) int32
