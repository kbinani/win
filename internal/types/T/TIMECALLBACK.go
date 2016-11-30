package win

//ref UINT
//ref DWORD_PTR

type TIMECALLBACK func(uTimerID UINT, uMsg UINT, dwUser DWORD_PTR, dw1 DWORD_PTR, dw2 DWORD_PTR)
