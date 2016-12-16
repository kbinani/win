package win

//ref DWORD
//ref PCCERT_CONTEXT
//ref CRYPT_ALGORITHM_IDENTIFIER
//ref PVOID
//ref PCCRL_CONTEXT
//ref PCRYPT_ATTRIBUTE

type CRYPT_SIGN_MESSAGE_PARA struct {
	CbSize             DWORD
	DwMsgEncodingType  DWORD
	PSigningCert       PCCERT_CONTEXT
	HashAlgorithm      CRYPT_ALGORITHM_IDENTIFIER
	PvHashAuxInfo      PVOID
	CMsgCert           DWORD
	RgpMsgCert         *PCCERT_CONTEXT
	CMsgCrl            DWORD
	RgpMsgCrl          *PCCRL_CONTEXT
	CAuthAttr          DWORD
	RgAuthAttr         PCRYPT_ATTRIBUTE
	CUnauthAttr        DWORD
	RgUnauthAttr       PCRYPT_ATTRIBUTE
	DwFlags            DWORD
	DwInnerContentType DWORD
}
