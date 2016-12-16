package win

//ref DWORD
//ref PCMSG_SIGNER_ENCODE_INFO
//ref PCERT_BLOB
//ref PCRL_BLOB

type CMSG_SIGNED_ENCODE_INFO struct {
	CbSize        DWORD
	CSigners      DWORD
	RgSigners     PCMSG_SIGNER_ENCODE_INFO
	CCertEncoded  DWORD
	RgCertEncoded PCERT_BLOB
	CCrlEncoded   DWORD
	RgCrlEncoded  PCRL_BLOB
}
