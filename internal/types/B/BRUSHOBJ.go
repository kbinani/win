package win

//ref ULONG
//ref PVOID
//ref FLONG
type BRUSHOBJ struct {
	ISolidColor ULONG
	PvRbrush    PVOID
	FlColorType FLONG
}
