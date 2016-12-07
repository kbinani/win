package win

//ref USHORT
//ref ULONG
//ref WCHAR

type STGOPTIONS struct {
	UsVersion        USHORT
	Reserved         USHORT
	UlSectorSize     ULONG
	PwcsTemplateFile *WCHAR
}
