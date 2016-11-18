package win

import "unsafe"

type COORD struct {
	X int16
	Y int16
}

func getCOORDFromUintptr(v uintptr) COORD {
	var ret COORD
	u32 := uint32(v)
	ret.X = *(*int16)(unsafe.Pointer(&u32))
	ret.Y = *(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&u32)) + uintptr(2)))
	return ret
}

func getUintptrFromCOORD(c COORD) uintptr {
	var ret uintptr
	xPtr := (*int16)(unsafe.Pointer(&ret))
	*xPtr = c.X
	yPtr := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&ret)) + uintptr(2)))
	*yPtr = c.Y
	return ret
}
