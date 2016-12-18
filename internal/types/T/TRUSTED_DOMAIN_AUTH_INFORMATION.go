package win

//ref ULONG
//ref PLSA_AUTH_INFORMATION

type TRUSTED_DOMAIN_AUTH_INFORMATION struct {
	IncomingAuthInfos                         ULONG
	IncomingAuthenticationInformation         PLSA_AUTH_INFORMATION
	IncomingPreviousAuthenticationInformation PLSA_AUTH_INFORMATION
	OutgoingAuthInfos                         ULONG
	OutgoingAuthenticationInformation         PLSA_AUTH_INFORMATION
	OutgoingPreviousAuthenticationInformation PLSA_AUTH_INFORMATION
}
