package win

//ref CRYPT_DIGEST_BLOB
//ref PCRYPT_RESOLVE_HCRYPTPROV_FUNC
//ref LPVOID
//ref PCRYPT_DECRYPT_PRIVATE_KEY_FUNC

type CRYPT_PKCS8_IMPORT_PARAMS struct {
	PrivateKey             CRYPT_DIGEST_BLOB
	PResolvehCryptProvFunc uintptr // PCRYPT_RESOLVE_HCRYPTPROV_FUNC
	PVoidResolveFunc       LPVOID
	PDecryptPrivateKeyFunc uintptr // PCRYPT_DECRYPT_PRIVATE_KEY_FUNC
	PVoidDecryptFunc       LPVOID
}
