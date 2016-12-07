package win

//ref DWORD
//ref BYTE

type SERIALIZEDPROPERTYVALUE struct {
	DwType DWORD
	Rgb    [1]BYTE
}
