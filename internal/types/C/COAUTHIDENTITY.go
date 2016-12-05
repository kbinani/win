package win

//ref USHORT
//ref ULONG

type COAUTHIDENTITY struct {
	User           *USHORT
	UserLength     ULONG
	Domain         *USHORT
	DomainLength   ULONG
	Password       *USHORT
	PasswordLength ULONG
	Flags          ULONG
}
