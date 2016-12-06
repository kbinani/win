package win

import "unsafe"

//ref DWORD
//ref IUnknown
//ref HBITMAP
//ref HMETAFILEPICT
//ref HENHMETAFILE
//ref HGLOBAL
//ref LPOLESTR
//ref IStream
//ref IStorage

type STGMEDIUM struct {
	Tymed          DWORD
	union1         uintptr
	PUnkForRelease *IUnknown
}

func (this *STGMEDIUM) GetHBitmap() HBITMAP {
	return *(*HBITMAP)(unsafe.Pointer(&this.union1))
}

func (this *STGMEDIUM) SetHBitmap(v HBITMAP) {
	*(*HBITMAP)(unsafe.Pointer(&this.union1)) = v
}

func (this *STGMEDIUM) GetHMetaFilePict() HMETAFILEPICT {
	return *(*HMETAFILEPICT)(unsafe.Pointer(&this.union1))
}

func (this *STGMEDIUM) SetHMetaFilePict(v HMETAFILEPICT) {
	*(*HMETAFILEPICT)(unsafe.Pointer(&this.union1)) = v
}

func (this *STGMEDIUM) GetHEnhMetaFile() HENHMETAFILE {
	return *(*HENHMETAFILE)(unsafe.Pointer(&this.union1))
}

func (this *STGMEDIUM) SetHEnhMetaFile(v HENHMETAFILE) {
	*(*HENHMETAFILE)(unsafe.Pointer(&this.union1)) = v
}

func (this *STGMEDIUM) GetHGlobal() HGLOBAL {
	return *(*HGLOBAL)(unsafe.Pointer(&this.union1))
}

func (this *STGMEDIUM) SetHGlobal(v HGLOBAL) {
	*(*HGLOBAL)(unsafe.Pointer(&this.union1)) = v
}

func (this *STGMEDIUM) GetLpszFileName() LPOLESTR {
	return *(*LPOLESTR)(unsafe.Pointer(&this.union1))
}

func (this *STGMEDIUM) SetLpszFileName(v LPOLESTR) {
	*(*LPOLESTR)(unsafe.Pointer(&this.union1)) = v
}

func (this *STGMEDIUM) GetPstm() *IStream {
	return *(**IStream)(unsafe.Pointer(&this.union1))
}

func (this *STGMEDIUM) SetPstm(v *IStream) {
	*(**IStream)(unsafe.Pointer(&this.union1)) = v
}

func (this *STGMEDIUM) GetPstg() *IStorage {
	return *(**IStorage)(unsafe.Pointer(&this.union1))
}

func (this *STGMEDIUM) SetPstg(v *IStorage) {
	*(**IStorage)(unsafe.Pointer(&this.union1)) = v
}
