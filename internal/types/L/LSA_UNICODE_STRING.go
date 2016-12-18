package win

//ref USHORT
//ref PWSTR

type LSA_UNICODE_STRING struct {
	Length        USHORT
	MaximumLength USHORT
	Buffer        PWSTR
}
