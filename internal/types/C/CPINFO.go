package win

//ref BYTE
//ref UINT
type CPINFO struct {
	MaxCharSize UINT
	DefaultChar [MAX_DEFAULTCHAR]BYTE
	LeadByte    [MAX_LEADBYTES]BYTE
}
