package win

//ref ULONG
//ref RECTL
//ref BYTE
type CLIPOBJ struct {
	IUniq        ULONG
	RclBounds    RECTL
	IDComplexity BYTE
	IFComplexity BYTE
	IMode        BYTE
	FjOptions    BYTE
}
