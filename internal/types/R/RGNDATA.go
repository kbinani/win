package win

//ref RGNDATAHEADER
type RGNDATA struct {
	Rdh    RGNDATAHEADER
	Buffer [1]byte
}
