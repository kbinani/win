package win

//ref INSTALLSPECTYPE
//ref INSTALLSPEC

type INSTALLDATA struct {
	Type     INSTALLSPECTYPE
	padding1 [pad4for64_0for32]byte
	Spec     INSTALLSPEC
}
