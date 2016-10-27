package win

//ref HCONV
//ref HSZ
//ref HDDEDATA
type PFNCALLBACK func(wType uint32, wFmt uint32, hConv HCONV, hsz1 HSZ, hsz2 HSZ, hData HDDEDATA, dwData1 uintptr, dwData2 uintptr) HDDEDATA
