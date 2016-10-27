package win

//ref DHSURF
//ref HSURF
//ref DHPDEV
//ref HDEV
//ref SIZEL
//ref ULONG
//ref PVOID
//ref LONG
//ref USHORT
type SURFOBJ struct {
	Dhsurf        DHSURF
	Hsurf         HSURF
	Dhpdev        DHPDEV
	Hdev          HDEV
	SizlBitmap    SIZEL
	CjBits        ULONG
	PvBits        PVOID
	PvScan0       PVOID
	LDelta        LONG
	IUniq         ULONG
	IBitmapFormat ULONG
	IType         USHORT
	FjBitmap      USHORT
}
