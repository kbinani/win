package win

//ref LSA_UNICODE_STRING
//ref PSID
//ref ULONG

type TRUSTED_DOMAIN_INFORMATION_EX struct {
	Name            LSA_UNICODE_STRING
	FlatName        LSA_UNICODE_STRING
	Sid             PSID
	TrustDirection  ULONG
	TrustType       ULONG
	TrustAttributes ULONG
}
