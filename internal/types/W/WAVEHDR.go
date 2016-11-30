package win

//ref LPSTR
//ref DWORD
//ref DWORD_PTR

type WAVEHDR struct {
	LpData          LPSTR
	DwBufferLength  DWORD
	DwBytesRecorded DWORD
	DwUser          DWORD_PTR
	DwFlags         DWORD
	DwLoops         DWORD
	LpNext          *WAVEHDR
	Reserved        DWORD_PTR
}
