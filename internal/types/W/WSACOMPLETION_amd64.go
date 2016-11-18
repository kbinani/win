package win

type WSACOMPLETION struct {
	Type       WSACOMPLETIONTYPE
	padding1   [4]byte
	Parameters WSACOMPLETION_Parameters
}

type WSACOMPLETION_Parameters struct {
	storage [24]byte
}
