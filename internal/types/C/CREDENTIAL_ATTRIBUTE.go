package win

//ref LPWSTR
type CREDENTIAL_ATTRIBUTE struct {
	Keyword   LPWSTR
	Flags     uint32
	ValueSize uint32
	Value     *byte
}
