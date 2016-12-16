package win

//ref BYTE
//ref WORD
//ref ALG_ID

type PUBLICKEYSTRUC struct {
	BType    BYTE
	BVersion BYTE
	Reserved WORD
	AiKeyAlg ALG_ID
}
