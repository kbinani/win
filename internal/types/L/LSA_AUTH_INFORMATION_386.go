package win

type LSA_AUTH_INFORMATION struct {
	LastUpdateTime LARGE_INTEGER
	AuthType       ULONG
	AuthInfoLength ULONG
	AuthInfo       PUCHAR
	padding0       [4]byte
}
