package win

type WSACOMPLETION struct {
	Type       WSACOMPLETIONTYPE
	Parameters WSACOMPLETION_Parameters
}

type WSACOMPLETION_Parameters struct {
	storage [12]byte
}
