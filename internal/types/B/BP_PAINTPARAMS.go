package win

//ref DWORD
//ref RECT
//ref BLENDFUNCTION
type BP_PAINTPARAMS struct {
	CbSize  DWORD
	DwFlags DWORD
	PrcExclude/*const*/ *RECT
	PBlendFunction/*const*/ *BLENDFUNCTION
}
