package win

//ref CRYPT_ALGORITHM_IDENTIFIER
//ref CRYPT_DATA_BLOB
//ref BYTE
//ref DWORD
//ref LPVOID
//ref BOOL

type PCRYPT_DECRYPT_PRIVATE_KEY_FUNC func(Algorithm CRYPT_ALGORITHM_IDENTIFIER, EncryptedPrivateKey CRYPT_DATA_BLOB, pbClearTextKey *BYTE, pcbClearTextKey *DWORD, pVoidDecryptFunc LPVOID) BOOL
