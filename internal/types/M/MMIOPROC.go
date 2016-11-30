package win

//ref LPSTR
//ref UINT
//ref LPARAM
//ref LRESULT

type MMIOPROC func(lpmmioinfo LPSTR, uMsg UINT, lParam1 LPARAM, lParam2 LPARAM) LRESULT
