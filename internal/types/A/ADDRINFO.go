package win

//ref SIZE_T
//ref PWSTR
type ADDRINFO struct {
	Ai_flags     int32
	Ai_family    int32
	Ai_socktype  int32
	Ai_protocol  int32
	Ai_addrlen   SIZE_T
	Ai_canonname PWSTR
	Ai_addr      uintptr // struct sockaddr *
	Ai_next      *ADDRINFO
}
