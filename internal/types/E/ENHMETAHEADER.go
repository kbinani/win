package win

//ref DWORD
//ref RECTL
//ref WORD
//ref SIZEL
type ENHMETAHEADER struct {
	IType          DWORD
	NSize          DWORD
	RclBounds      RECTL
	RclFrame       RECTL
	DSignature     DWORD
	NVersion       DWORD
	NBytes         DWORD
	NRecords       DWORD
	NHandles       WORD
	SReserved      WORD
	NDescription   DWORD
	OffDescription DWORD
	NPalEntries    DWORD
	SzlDevice      SIZEL
	SzlMillimeters SIZEL
	CbPixelFormat  DWORD
	OffPixelFormat DWORD
	BOpenGL        DWORD
	SzlMicrometers SIZEL
}
