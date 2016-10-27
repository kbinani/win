package win

func OleLoadPictureFile(file VARIANT, picture *LPDISPATCH) HRESULT {
	args := unpackVARIANT(file)
	args = append(args, uintptr(unsafe.Pointer(picture)))
	ret := syscallN(oleLoadPictureFile, args)
	return HRESULT(ret)
}
