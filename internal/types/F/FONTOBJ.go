package win

//ref ULONG
//ref FLONG
//ref ULONG_PTR
//ref SIZE
//ref PVOID
type FONTOBJ struct {
	IUniq        ULONG
	IFace        ULONG
	CxMax        ULONG
	FlFontType   FLONG
	ITTUniq      ULONG_PTR
	IFile        ULONG_PTR
	SizLogResPpi SIZE
	UlStyleSize  ULONG
	PvConsumer   PVOID
	PvProducer   PVOID
}
