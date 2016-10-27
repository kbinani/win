package win

//ref BYTE
//ref WORD
type ACCEL struct {
	FVirt BYTE
	Key   WORD
	Cmd   WORD
}
