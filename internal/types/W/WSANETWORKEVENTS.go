package win

type WSANETWORKEVENTS struct {
	LNetworkEvents int32
	IErrorCode     [FD_MAX_EVENTS]int32
}
