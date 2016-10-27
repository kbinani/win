package win

//ref UINT
// STRRET struct is padded with 8 bytes
type STRRET struct {
	UType    UINT
	padding1 [4]byte
	cStr     [260]byte
	padding2 [4]byte
}
