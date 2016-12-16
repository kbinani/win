package win

//ref DWORD
//ref PFN_CRYPT_ALLOC
//ref PFN_CRYPT_FREE

type CRYPT_DECODE_PARA struct {
	CbSize   DWORD
	PfnAlloc uintptr // PFN_CRYPT_ALLOC
	PfnFree  uintptr // PFN_CRYPT_FREE
}
