package win

type HANDLER_FUNCTION_EX func(dwControl uint32, dwEventType uint32, lpEventData uintptr, lpContext uintptr) uint32
