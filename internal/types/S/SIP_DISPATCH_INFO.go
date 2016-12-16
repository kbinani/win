package win

//ref DWORD
//ref HANDLE
//ref PCryptSIPGetSignedDataMsg
//ref PCryptSIPPutSignedDataMsg
//ref PCryptSIPCreateIndirectData
//ref PCryptSIPVerifyIndirectData
//ref PCryptSIPRemoveSignedDataMsg

type SIP_DISPATCH_INFO struct {
	CbSize   DWORD
	HSIP     HANDLE
	PfGet    uintptr // PCryptSIPGetSignedDataMsg
	PfPut    uintptr // PCryptSIPPutSignedDataMsg
	PfCreate uintptr // PCryptSIPCreateIndirectData
	PfVerify uintptr // PCryptSIPVerifyIndirectData
	PfRemove uintptr // PCryptSIPRemoveSignedDataMsg
}
