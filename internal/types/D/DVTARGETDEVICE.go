package win

//ref DWORD
//ref WORD
//ref BYTE

type DVTARGETDEVICE struct {
	TdSize             DWORD
	TdDriverNameOffset WORD
	TdDeviceNameOffset WORD
	TdPortNameOffset   WORD
	TdExtDevmodeOffset WORD
	TdData             [1]BYTE
}
