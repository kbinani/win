package win

//ref FOURCC
//ref DWORD

type MMCKINFO struct {
	Ckid         FOURCC
	Cksize       DWORD
	FccType      FOURCC
	DwDataOffset DWORD
	DwFlags      DWORD
}
