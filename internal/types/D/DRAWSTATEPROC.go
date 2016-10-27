package win

//ref HDC
//ref BOOL
type DRAWSTATEPROC func(hdc HDC, lData uintptr, wData uintptr, cx int32, cy int32) BOOL
