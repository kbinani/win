package win

//ref USHORT
//ref ULONG
//ref PVOID
//ref SAFEARRAYBOUND
type SAFEARRAY struct {
	CDims      USHORT
	FFeatures  USHORT
	CbElements ULONG
	CLocks     ULONG
	PvData     PVOID
	Rgsabound  [1]SAFEARRAYBOUND
}
