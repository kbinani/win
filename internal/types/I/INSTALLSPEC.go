package win

//ref WCHAR
//ref GUID
//ref DWORD

type INSTALLSPEC struct {
	union1 [ptrsize + 16]byte
}

type INSTALLSPEC_AppName struct {
	Name  *WCHAR
	GPOId GUID
}

type INSTALLSPEC_COMClass struct {
	Clsid  GUID
	ClsCtx DWORD
}

func (this *INSTALLSPEC) GetAppName() INSTALLSPEC_AppName {
	return *(*INSTALLSPEC_AppName)(unsafe.Pointer(&this.union1))
}

func (this *INSTALLSPEC) SetAppName(v INSTALLSPEC_AppName) {
	*(*INSTALLSPEC_AppName)(unsafe.Pointer(&this.union1)) = v
}

func (this *INSTALLSPEC) GetFileExt() *WCHAR {
	return *(**WCHAR)(unsafe.Pointer(&this.union1))
}

func (this *INSTALLSPEC) SetFileExt(v *WCHAR) {
	*(**WCHAR)(unsafe.Pointer(&this.union1)) = v
}

func (this *INSTALLSPEC) GetProgId() *WCHAR {
	return *(**WCHAR)(unsafe.Pointer(&this.union1))
}

func (this *INSTALLSPEC) SetProgId(v *WCHAR) {
	*(**WCHAR)(unsafe.Pointer(&this.union1)) = v
}

func (this *INSTALLSPEC) GetCOMClass() INSTALLSPEC_COMClass {
	return *(*INSTALLSPEC_COMClass)(unsafe.Pointer(&this.union1))
}

func (this *INSTALLSPEC) SetCOMClass(v INSTALLSPEC_COMClass) {
	*(*INSTALLSPEC_COMClass)(unsafe.Pointer(&this.union1)) = v
}
