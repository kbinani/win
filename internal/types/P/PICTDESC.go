package win

//ref UINT
//ref HBITMAP
//ref HPALETTE
//ref HMETAFILE
//ref HICON
//ref HENHMETAFILE
//import unsafe
type PICTDESC struct {
	CbSizeofstruct UINT
	PicType        UINT
	union1         uintptr
	union2         int32
	union3         int32
}

func (this *PICTDESC) Hbitmap() HBITMAP {
	return HBITMAP(this.union1)
}

func (this *PICTDESC) Hpal() HPALETTE {
	var ptr uintptr
	if is64 {
		*(*int32)(unsafe.Pointer(&ptr)) = this.union2
		*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + 4)) = this.union3
	} else {
		*(*int32)(unsafe.Pointer(&ptr)) = this.union2
	}
	return HPALETTE(ptr)
}

func (this *PICTDESC) Hmeta() HMETAFILE {
	return HMETAFILE(this.union1)
}

func (this *PICTDESC) XExt() int32 {
	return this.union2
}

func (this *PICTDESC) YExt() int32 {
	return this.union3
}

func (this *PICTDESC) Hicon() HICON {
	return HICON(this.union1)
}

func (this *PICTDESC) Hemf() HENHMETAFILE {
	return HENHMETAFILE(this.union1)
}
