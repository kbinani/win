package win

//ref IID
//ref IUnknown
//ref HRESULT

type MULTI_QI struct {
	PIID/*const*/ *IID
	PItf *IUnknown
	Hr   HRESULT
}
