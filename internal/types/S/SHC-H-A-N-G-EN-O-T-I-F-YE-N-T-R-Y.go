package win

import "unsafe"

//ref PCIDLIST_ABSOLUTE
//ref BOOL

type SHChangeNotifyEntry struct {
	storage1 [ptrsize + 4]byte
}

func (this *SHChangeNotifyEntry) Pidl() *PCIDLIST_ABSOLUTE {
	return (*PCIDLIST_ABSOLUTE)(unsafe.Pointer(&this.storage1[0]))
}

func (this *SHChangeNotifyEntry) FRecursive() *BOOL {
	return (*BOOL)(unsafe.Pointer(&this.storage1[ptrsize]))
}
