package win

//ref UCHAR
//ref SID_IDENTIFIER_AUTHORITY
//ref ULONG
type SID struct {
	Revision            UCHAR
	SubAuthorityCount   UCHAR
	IdentifierAuthority SID_IDENTIFIER_AUTHORITY
	SubAuthority        [ANYSIZE_ARRAY]ULONG
}
