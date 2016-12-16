package win

//ref DWORD
//ref LPWSTR
//ref HCRYPTPROV
//ref HANDLE
//ref HCRYPTMSG

type CRYPTCATSTORE struct {
	CbStruct        DWORD
	DwPublicVersion DWORD
	PwszP7File      LPWSTR
	HProv           HCRYPTPROV
	DwEncodingType  DWORD
	FdwStoreFlags   DWORD
	HReserved       HANDLE
	HAttrs          HANDLE
	HCryptMsg       HCRYPTMSG
	HSorted         HANDLE
}
