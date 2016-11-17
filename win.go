package win

import (
	"unicode/utf16"
	"unsafe"
)

type COORD struct {
	X int16
	Y int16
}

type BLENDFUNCTION struct {
	BlendOp             byte
	BlendFlags          byte
	SourceConstantAlpha byte
	AlphaFormat         byte
}

var is64 bool

type CY struct {
	union1 [8]byte
}

func (this CY) Lo() uint32 {
	return *(*uint32)(unsafe.Pointer(&this.union1[0]))
}

func (this CY) Hi() int32 {
	return *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1[0])) + uintptr(4)))
}

func (this CY) Int64() int64 {
	return *(*int64)(unsafe.Pointer(&this.union1[0]))
}

func init() {
	var ptr uintptr
	is64 = unsafe.Sizeof(ptr) == 8
}

func unicode16FromString(s string) []uint16 {
	r := make([]rune, 0)
	for _, c := range s {
		r = append(r, c)
	}
	b := utf16.Encode(r)
	return append(b, uint16(0))
}

func stringFromUnicode16(s *uint16) string {
	if s == nil {
		return ""
	}
	buffer := []uint16{}
	ptr := uintptr(unsafe.Pointer(s))
	for true {
		ch := *(*uint16)(unsafe.Pointer(ptr))
		if ch == 0 {
			break
		}
		buffer = append(buffer, ch)
		ptr += unsafe.Sizeof(ch)
	}
	return string(utf16.Decode(buffer))
}
