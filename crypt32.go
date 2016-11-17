// +build windows

package win

import (
	"unsafe"
)

var (
	// Library
	libcrypt32 uintptr

	// Functions
	certAddCRLContextToStore                          uintptr
	certAddCRLLinkToStore                             uintptr
	certAddEncodedCRLToStore                          uintptr
	certAddEncodedCertificateToSystemStore            uintptr
	certAddSerializedElementToStore                   uintptr
	certAddStoreToCollection                          uintptr
	certAlgIdToOID                                    uintptr
	certCloseStore                                    uintptr
	certCompareCertificateName                        uintptr
	certCompareIntegerBlob                            uintptr
	certControlStore                                  uintptr
	certCreateCRLContext                              uintptr
	certDeleteCRLFromStore                            uintptr
	certDuplicateCRLContext                           uintptr
	certDuplicateStore                                uintptr
	certEnumCRLContextProperties                      uintptr
	certEnumCRLsInStore                               uintptr
	certFindCRLInStore                                uintptr
	certFindExtension                                 uintptr
	certFreeCRLContext                                uintptr
	certGetCRLContextProperty                         uintptr
	certGetStoreProperty                              uintptr
	certNameToStr                                     uintptr
	certOIDToAlgId                                    uintptr
	certRDNValueToStr                                 uintptr
	certRemoveStoreFromCollection                     uintptr
	certSaveStore                                     uintptr
	certSerializeCRLStoreElement                      uintptr
	certSetCRLContextProperty                         uintptr
	certSetStoreProperty                              uintptr
	certStrToName                                     uintptr
	certUnregisterPhysicalStore                       uintptr
	certUnregisterSystemStore                         uintptr
	certVerifyCRLTimeValidity                         uintptr
	cryptBinaryToString                               uintptr
	cryptDecodeObject                                 uintptr
	cryptEncodeObject                                 uintptr
	cryptExportPKCS8                                  uintptr
	cryptFindLocalizedName                            uintptr
	cryptFormatObject                                 uintptr
	cryptGetKeyIdentifierProperty                     uintptr
	cryptGetMessageSignerCount                        uintptr
	cryptGetOIDFunctionValue                          uintptr
	cryptMemAlloc                                     uintptr
	cryptMemFree                                      uintptr
	cryptMemRealloc                                   uintptr
	cryptMsgCalculateEncodedLength                    uintptr
	cryptProtectMemory                                uintptr
	cryptRegisterDefaultOIDFunction                   uintptr
	cryptRegisterOIDFunction                          uintptr
	cryptSIPRemoveProvider                            uintptr
	cryptSIPRetrieveSubjectGuid                       uintptr
	cryptSIPRetrieveSubjectGuidForCatalogFile         uintptr
	cryptSetKeyIdentifierProperty                     uintptr
	cryptSetOIDFunctionValue                          uintptr
	cryptStringToBinary                               uintptr
	cryptUnprotectMemory                              uintptr
	cryptUnregisterDefaultOIDFunction                 uintptr
	cryptUnregisterOIDFunction                        uintptr
	pFXExportCertStore                                uintptr
	pFXExportCertStoreEx                              uintptr
	pFXImportCertStore                                uintptr
	pFXIsPFXBlob                                      uintptr
	pFXVerifyPassword                                 uintptr
	i_CertUpdateStore                                 uintptr
	i_CryptAllocTls                                   uintptr
	i_CryptDetachTls                                  uintptr
	i_CryptFindLruEntry                               uintptr
	i_CryptFindLruEntryData                           uintptr
	i_CryptFreeTls                                    uintptr
	i_CryptGetDefaultCryptProv                        uintptr
	i_CryptGetOssGlobal                               uintptr
	i_CryptGetTls                                     uintptr
	i_CryptInstallOssGlobal                           uintptr
	i_CryptReadTrustedPublisherDWORDValueFromRegistry uintptr
	i_CryptSetTls                                     uintptr
)

func init() {
	// Library
	libcrypt32 = doLoadLibrary("crypt32.dll")

	// Functions
	certAddCRLContextToStore = doGetProcAddress(libcrypt32, "CertAddCRLContextToStore")
	certAddCRLLinkToStore = doGetProcAddress(libcrypt32, "CertAddCRLLinkToStore")
	certAddEncodedCRLToStore = doGetProcAddress(libcrypt32, "CertAddEncodedCRLToStore")
	certAddEncodedCertificateToSystemStore = doGetProcAddress(libcrypt32, "CertAddEncodedCertificateToSystemStoreW")
	certAddSerializedElementToStore = doGetProcAddress(libcrypt32, "CertAddSerializedElementToStore")
	certAddStoreToCollection = doGetProcAddress(libcrypt32, "CertAddStoreToCollection")
	certAlgIdToOID = doGetProcAddress(libcrypt32, "CertAlgIdToOID")
	certCloseStore = doGetProcAddress(libcrypt32, "CertCloseStore")
	certCompareCertificateName = doGetProcAddress(libcrypt32, "CertCompareCertificateName")
	certCompareIntegerBlob = doGetProcAddress(libcrypt32, "CertCompareIntegerBlob")
	certControlStore = doGetProcAddress(libcrypt32, "CertControlStore")
	certCreateCRLContext = doGetProcAddress(libcrypt32, "CertCreateCRLContext")
	certDeleteCRLFromStore = doGetProcAddress(libcrypt32, "CertDeleteCRLFromStore")
	certDuplicateCRLContext = doGetProcAddress(libcrypt32, "CertDuplicateCRLContext")
	certDuplicateStore = doGetProcAddress(libcrypt32, "CertDuplicateStore")
	certEnumCRLContextProperties = doGetProcAddress(libcrypt32, "CertEnumCRLContextProperties")
	certEnumCRLsInStore = doGetProcAddress(libcrypt32, "CertEnumCRLsInStore")
	certFindCRLInStore = doGetProcAddress(libcrypt32, "CertFindCRLInStore")
	certFindExtension = doGetProcAddress(libcrypt32, "CertFindExtension")
	certFreeCRLContext = doGetProcAddress(libcrypt32, "CertFreeCRLContext")
	certGetCRLContextProperty = doGetProcAddress(libcrypt32, "CertGetCRLContextProperty")
	certGetStoreProperty = doGetProcAddress(libcrypt32, "CertGetStoreProperty")
	certNameToStr = doGetProcAddress(libcrypt32, "CertNameToStrW")
	certOIDToAlgId = doGetProcAddress(libcrypt32, "CertOIDToAlgId")
	certRDNValueToStr = doGetProcAddress(libcrypt32, "CertRDNValueToStrW")
	certRemoveStoreFromCollection = doGetProcAddress(libcrypt32, "CertRemoveStoreFromCollection")
	certSaveStore = doGetProcAddress(libcrypt32, "CertSaveStore")
	certSerializeCRLStoreElement = doGetProcAddress(libcrypt32, "CertSerializeCRLStoreElement")
	certSetCRLContextProperty = doGetProcAddress(libcrypt32, "CertSetCRLContextProperty")
	certSetStoreProperty = doGetProcAddress(libcrypt32, "CertSetStoreProperty")
	certStrToName = doGetProcAddress(libcrypt32, "CertStrToNameW")
	certUnregisterPhysicalStore = doGetProcAddress(libcrypt32, "CertUnregisterPhysicalStore")
	certUnregisterSystemStore = doGetProcAddress(libcrypt32, "CertUnregisterSystemStore")
	certVerifyCRLTimeValidity = doGetProcAddress(libcrypt32, "CertVerifyCRLTimeValidity")
	cryptBinaryToString = doGetProcAddress(libcrypt32, "CryptBinaryToStringW")
	cryptDecodeObject = doGetProcAddress(libcrypt32, "CryptDecodeObject")
	cryptEncodeObject = doGetProcAddress(libcrypt32, "CryptEncodeObject")
	cryptExportPKCS8 = doGetProcAddress(libcrypt32, "CryptExportPKCS8")
	cryptFindLocalizedName = doGetProcAddress(libcrypt32, "CryptFindLocalizedName")
	cryptFormatObject = doGetProcAddress(libcrypt32, "CryptFormatObject")
	cryptGetKeyIdentifierProperty = doGetProcAddress(libcrypt32, "CryptGetKeyIdentifierProperty")
	cryptGetMessageSignerCount = doGetProcAddress(libcrypt32, "CryptGetMessageSignerCount")
	cryptGetOIDFunctionValue = doGetProcAddress(libcrypt32, "CryptGetOIDFunctionValue")
	cryptMemAlloc = doGetProcAddress(libcrypt32, "CryptMemAlloc")
	cryptMemFree = doGetProcAddress(libcrypt32, "CryptMemFree")
	cryptMemRealloc = doGetProcAddress(libcrypt32, "CryptMemRealloc")
	cryptMsgCalculateEncodedLength = doGetProcAddress(libcrypt32, "CryptMsgCalculateEncodedLength")
	cryptProtectMemory = doGetProcAddress(libcrypt32, "CryptProtectMemory")
	cryptRegisterDefaultOIDFunction = doGetProcAddress(libcrypt32, "CryptRegisterDefaultOIDFunction")
	cryptRegisterOIDFunction = doGetProcAddress(libcrypt32, "CryptRegisterOIDFunction")
	cryptSIPRemoveProvider = doGetProcAddress(libcrypt32, "CryptSIPRemoveProvider")
	cryptSIPRetrieveSubjectGuid = doGetProcAddress(libcrypt32, "CryptSIPRetrieveSubjectGuid")
	cryptSIPRetrieveSubjectGuidForCatalogFile = doGetProcAddress(libcrypt32, "CryptSIPRetrieveSubjectGuidForCatalogFile")
	cryptSetKeyIdentifierProperty = doGetProcAddress(libcrypt32, "CryptSetKeyIdentifierProperty")
	cryptSetOIDFunctionValue = doGetProcAddress(libcrypt32, "CryptSetOIDFunctionValue")
	cryptStringToBinary = doGetProcAddress(libcrypt32, "CryptStringToBinaryW")
	cryptUnprotectMemory = doGetProcAddress(libcrypt32, "CryptUnprotectMemory")
	cryptUnregisterDefaultOIDFunction = doGetProcAddress(libcrypt32, "CryptUnregisterDefaultOIDFunction")
	cryptUnregisterOIDFunction = doGetProcAddress(libcrypt32, "CryptUnregisterOIDFunction")
	pFXExportCertStore = doGetProcAddress(libcrypt32, "PFXExportCertStore")
	pFXExportCertStoreEx = doGetProcAddress(libcrypt32, "PFXExportCertStoreEx")
	pFXImportCertStore = doGetProcAddress(libcrypt32, "PFXImportCertStore")
	pFXIsPFXBlob = doGetProcAddress(libcrypt32, "PFXIsPFXBlob")
	pFXVerifyPassword = doGetProcAddress(libcrypt32, "PFXVerifyPassword")
	i_CertUpdateStore = doGetProcAddress(libcrypt32, "I_CertUpdateStore")
	i_CryptAllocTls = doGetProcAddress(libcrypt32, "I_CryptAllocTls")
	i_CryptDetachTls = doGetProcAddress(libcrypt32, "I_CryptDetachTls")
	i_CryptFindLruEntry = doGetProcAddress(libcrypt32, "I_CryptFindLruEntry")
	i_CryptFindLruEntryData = doGetProcAddress(libcrypt32, "I_CryptFindLruEntryData")
	i_CryptFreeTls = doGetProcAddress(libcrypt32, "I_CryptFreeTls")
	i_CryptGetDefaultCryptProv = doGetProcAddress(libcrypt32, "I_CryptGetDefaultCryptProv")
	i_CryptGetOssGlobal = doGetProcAddress(libcrypt32, "I_CryptGetOssGlobal")
	i_CryptGetTls = doGetProcAddress(libcrypt32, "I_CryptGetTls")
	i_CryptInstallOssGlobal = doGetProcAddress(libcrypt32, "I_CryptInstallOssGlobal")
	i_CryptReadTrustedPublisherDWORDValueFromRegistry = doGetProcAddress(libcrypt32, "I_CryptReadTrustedPublisherDWORDValueFromRegistry")
	i_CryptSetTls = doGetProcAddress(libcrypt32, "I_CryptSetTls")
}

func CertAddCRLContextToStore(hCertStore HCERTSTORE, pCrlContext PCCRL_CONTEXT, dwAddDisposition uint32, ppStoreContext *PCCRL_CONTEXT) bool {
	ret1 := syscall6(certAddCRLContextToStore, 4,
		uintptr(hCertStore),
		uintptr(unsafe.Pointer(pCrlContext)),
		uintptr(dwAddDisposition),
		uintptr(unsafe.Pointer(ppStoreContext)),
		0,
		0)
	return ret1 != 0
}

func CertAddCRLLinkToStore(hCertStore HCERTSTORE, pCrlContext PCCRL_CONTEXT, dwAddDisposition uint32, ppStoreContext *PCCRL_CONTEXT) bool {
	ret1 := syscall6(certAddCRLLinkToStore, 4,
		uintptr(hCertStore),
		uintptr(unsafe.Pointer(pCrlContext)),
		uintptr(dwAddDisposition),
		uintptr(unsafe.Pointer(ppStoreContext)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCCTL_CONTEXT, PCCTL_CONTEXT *
// func CertAddCTLContextToStore(hCertStore HCERTSTORE, pCtlContext PCCTL_CONTEXT, dwAddDisposition uint32, ppStoreContext PCCTL_CONTEXT *) bool

// TODO: Unknown type(s): PCCTL_CONTEXT, PCCTL_CONTEXT *
// func CertAddCTLLinkToStore(hCertStore HCERTSTORE, pCtlContext PCCTL_CONTEXT, dwAddDisposition uint32, ppStoreContext PCCTL_CONTEXT *) bool

// TODO: Unknown type(s): PCCERT_CONTEXT, PCCERT_CONTEXT *
// func CertAddCertificateContextToStore(hCertStore HCERTSTORE, pCertContext PCCERT_CONTEXT, dwAddDisposition uint32, ppStoreContext PCCERT_CONTEXT *) bool

// TODO: Unknown type(s): PCCERT_CONTEXT, PCCERT_CONTEXT *
// func CertAddCertificateLinkToStore(hCertStore HCERTSTORE, pCertContext PCCERT_CONTEXT, dwAddDisposition uint32, ppStoreContext PCCERT_CONTEXT *) bool

func CertAddEncodedCRLToStore(hCertStore HCERTSTORE, dwCertEncodingType uint32, pbCrlEncoded /*const*/ *byte, cbCrlEncoded uint32, dwAddDisposition uint32, ppCrlContext *PCCRL_CONTEXT) bool {
	ret1 := syscall6(certAddEncodedCRLToStore, 6,
		uintptr(hCertStore),
		uintptr(dwCertEncodingType),
		uintptr(unsafe.Pointer(pbCrlEncoded)),
		uintptr(cbCrlEncoded),
		uintptr(dwAddDisposition),
		uintptr(unsafe.Pointer(ppCrlContext)))
	return ret1 != 0
}

// TODO: Unknown type(s): PCCTL_CONTEXT *
// func CertAddEncodedCTLToStore(hCertStore HCERTSTORE, dwMsgAndCertEncodingType uint32, pbCtlEncoded /*const*/ *byte, cbCtlEncoded uint32, dwAddDisposition uint32, ppCtlContext PCCTL_CONTEXT *) bool

// TODO: Unknown type(s): PCCERT_CONTEXT *
// func CertAddEncodedCertificateToStore(hCertStore HCERTSTORE, dwCertEncodingType uint32, pbCertEncoded /*const*/ *byte, cbCertEncoded uint32, dwAddDisposition uint32, ppCertContext PCCERT_CONTEXT *) bool

func CertAddEncodedCertificateToSystemStore(szCertStoreName string, pbCertEncoded /*const*/ *byte, cbCertEncoded uint32) bool {
	szCertStoreNameStr := unicode16FromString(szCertStoreName)
	ret1 := syscall3(certAddEncodedCertificateToSystemStore, 3,
		uintptr(unsafe.Pointer(&szCertStoreNameStr[0])),
		uintptr(unsafe.Pointer(pbCertEncoded)),
		uintptr(cbCertEncoded))
	return ret1 != 0
}

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertAddEnhancedKeyUsageIdentifier(pCertContext PCCERT_CONTEXT, pszUsageIdentifier /*const*/ LPCSTR) bool

func CertAddSerializedElementToStore(hCertStore HCERTSTORE, pbElement /*const*/ *byte, cbElement uint32, dwAddDisposition uint32, dwFlags uint32, dwContextTypeFlags uint32, pdwContextType *uint32, ppvContext /*const*/ uintptr) bool {
	ret1 := syscall9(certAddSerializedElementToStore, 8,
		uintptr(hCertStore),
		uintptr(unsafe.Pointer(pbElement)),
		uintptr(cbElement),
		uintptr(dwAddDisposition),
		uintptr(dwFlags),
		uintptr(dwContextTypeFlags),
		uintptr(unsafe.Pointer(pdwContextType)),
		ppvContext,
		0)
	return ret1 != 0
}

func CertAddStoreToCollection(hCollectionStore HCERTSTORE, hSiblingStore HCERTSTORE, dwUpdateFlags uint32, dwPriority uint32) bool {
	ret1 := syscall6(certAddStoreToCollection, 4,
		uintptr(hCollectionStore),
		uintptr(hSiblingStore),
		uintptr(dwUpdateFlags),
		uintptr(dwPriority),
		0,
		0)
	return ret1 != 0
}

func CertAlgIdToOID(dwAlgId uint32) LPCSTR {
	ret1 := syscall3(certAlgIdToOID, 1,
		uintptr(dwAlgId),
		0,
		0)
	return (LPCSTR)(unsafe.Pointer(ret1))
}

func CertCloseStore(hCertStore HCERTSTORE, dwFlags uint32) bool {
	ret1 := syscall3(certCloseStore, 2,
		uintptr(hCertStore),
		uintptr(dwFlags),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCERT_INFO
// func CertCompareCertificate(dwCertEncodingType uint32, pCertId1 PCERT_INFO, pCertId2 PCERT_INFO) bool

func CertCompareCertificateName(dwCertEncodingType uint32, pCertName1 PCERT_NAME_BLOB, pCertName2 PCERT_NAME_BLOB) bool {
	ret1 := syscall3(certCompareCertificateName, 3,
		uintptr(dwCertEncodingType),
		uintptr(unsafe.Pointer(pCertName1)),
		uintptr(unsafe.Pointer(pCertName2)))
	return ret1 != 0
}

func CertCompareIntegerBlob(pInt1 PCRYPT_INTEGER_BLOB, pInt2 PCRYPT_INTEGER_BLOB) bool {
	ret1 := syscall3(certCompareIntegerBlob, 2,
		uintptr(unsafe.Pointer(pInt1)),
		uintptr(unsafe.Pointer(pInt2)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCERT_PUBLIC_KEY_INFO
// func CertComparePublicKeyInfo(dwCertEncodingType uint32, pPublicKey1 PCERT_PUBLIC_KEY_INFO, pPublicKey2 PCERT_PUBLIC_KEY_INFO) bool

func CertControlStore(hCertStore HCERTSTORE, dwFlags uint32, dwCtrlType uint32, pvCtrlPara /*const*/ uintptr) bool {
	ret1 := syscall6(certControlStore, 4,
		uintptr(hCertStore),
		uintptr(dwFlags),
		uintptr(dwCtrlType),
		pvCtrlPara,
		0,
		0)
	return ret1 != 0
}

func CertCreateCRLContext(dwCertEncodingType uint32, pbCrlEncoded /*const*/ *byte, cbCrlEncoded uint32) PCCRL_CONTEXT {
	ret1 := syscall3(certCreateCRLContext, 3,
		uintptr(dwCertEncodingType),
		uintptr(unsafe.Pointer(pbCrlEncoded)),
		uintptr(cbCrlEncoded))
	return (PCCRL_CONTEXT)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): PCCTL_CONTEXT
// func CertCreateCTLContext(dwMsgAndCertEncodingType uint32, pbCtlEncoded /*const*/ *byte, cbCtlEncoded uint32) PCCTL_CONTEXT

// TODO: Unknown type(s): PCCERT_CONTEXT, PCRYPT_ATTRIBUTE, PCTL_ENTRY
// func CertCreateCTLEntryFromCertificateContextProperties(pCertContext PCCERT_CONTEXT, cOptAttr uint32, rgOptAttr PCRYPT_ATTRIBUTE, dwFlags uint32, pvReserved uintptr, pCtlEntry PCTL_ENTRY, pcbCtlEntry *uint32) bool

// TODO: Unknown type(s): HCERTCHAINENGINE *, PCERT_CHAIN_ENGINE_CONFIG
// func CertCreateCertificateChainEngine(pConfig PCERT_CHAIN_ENGINE_CONFIG, phChainEngine HCERTCHAINENGINE *) bool

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertCreateCertificateContext(dwCertEncodingType uint32, pbCertEncoded /*const*/ *byte, cbCertEncoded uint32) PCCERT_CONTEXT

// TODO: Unknown type(s): HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, PCCERT_CONTEXT, PCERT_EXTENSIONS, PCRYPT_ALGORITHM_IDENTIFIER, PCRYPT_KEY_PROV_INFO, PSYSTEMTIME
// func CertCreateSelfSignCertificate(hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, pSubjectIssuerBlob PCERT_NAME_BLOB, dwFlags uint32, pKeyProvInfo PCRYPT_KEY_PROV_INFO, pSignatureAlgorithm PCRYPT_ALGORITHM_IDENTIFIER, pStartTime PSYSTEMTIME, pEndTime PSYSTEMTIME, pExtensions PCERT_EXTENSIONS) PCCERT_CONTEXT

func CertDeleteCRLFromStore(pCrlContext PCCRL_CONTEXT) bool {
	ret1 := syscall3(certDeleteCRLFromStore, 1,
		uintptr(unsafe.Pointer(pCrlContext)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCCTL_CONTEXT
// func CertDeleteCTLFromStore(pCtlContext PCCTL_CONTEXT) bool

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertDeleteCertificateFromStore(pCertContext PCCERT_CONTEXT) bool

func CertDuplicateCRLContext(pCrlContext PCCRL_CONTEXT) PCCRL_CONTEXT {
	ret1 := syscall3(certDuplicateCRLContext, 1,
		uintptr(unsafe.Pointer(pCrlContext)),
		0,
		0)
	return (PCCRL_CONTEXT)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): PCCTL_CONTEXT
// func CertDuplicateCTLContext(pCtlContext PCCTL_CONTEXT) PCCTL_CONTEXT

// TODO: Unknown type(s): PCCERT_CHAIN_CONTEXT
// func CertDuplicateCertificateChain(pChainContext PCCERT_CHAIN_CONTEXT) PCCERT_CHAIN_CONTEXT

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertDuplicateCertificateContext(pCertContext PCCERT_CONTEXT) PCCERT_CONTEXT

func CertDuplicateStore(hCertStore HCERTSTORE) HCERTSTORE {
	ret1 := syscall3(certDuplicateStore, 1,
		uintptr(hCertStore),
		0,
		0)
	return HCERTSTORE(ret1)
}

func CertEnumCRLContextProperties(pCrlContext PCCRL_CONTEXT, dwPropId uint32) uint32 {
	ret1 := syscall3(certEnumCRLContextProperties, 2,
		uintptr(unsafe.Pointer(pCrlContext)),
		uintptr(dwPropId),
		0)
	return uint32(ret1)
}

func CertEnumCRLsInStore(hCertStore HCERTSTORE, pPrevCrlContext PCCRL_CONTEXT) PCCRL_CONTEXT {
	ret1 := syscall3(certEnumCRLsInStore, 2,
		uintptr(hCertStore),
		uintptr(unsafe.Pointer(pPrevCrlContext)),
		0)
	return (PCCRL_CONTEXT)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): PCCTL_CONTEXT
// func CertEnumCTLContextProperties(pCtlContext PCCTL_CONTEXT, dwPropId uint32) uint32

// TODO: Unknown type(s): PCCTL_CONTEXT
// func CertEnumCTLsInStore(hCertStore HCERTSTORE, pPrevCtlContext PCCTL_CONTEXT) PCCTL_CONTEXT

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertEnumCertificateContextProperties(pCertContext PCCERT_CONTEXT, dwPropId uint32) uint32

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertEnumCertificatesInStore(hCertStore HCERTSTORE, pPrevCertContext PCCERT_CONTEXT) PCCERT_CONTEXT

// TODO: Unknown type(s): PFN_CERT_ENUM_PHYSICAL_STORE
// func CertEnumPhysicalStore(pvSystemStore /*const*/ uintptr, dwFlags uint32, pvArg uintptr, pfnEnum PFN_CERT_ENUM_PHYSICAL_STORE) bool

// TODO: Unknown type(s): PCCTL_CONTEXT
// func CertEnumSubjectInSortedCTL(pCtlContext PCCTL_CONTEXT, ppvNextSubject uintptr, pSubjectIdentifier PCRYPT_DER_BLOB, pEncodedAttributes PCRYPT_DER_BLOB) bool

// TODO: Unknown type(s): PFN_CERT_ENUM_SYSTEM_STORE
// func CertEnumSystemStore(dwFlags uint32, pvSystemStoreLocationPara uintptr, pvArg uintptr, pfnEnum PFN_CERT_ENUM_SYSTEM_STORE) bool

// TODO: Unknown type(s): PFN_CERT_ENUM_SYSTEM_STORE_LOCATION
// func CertEnumSystemStoreLocation(dwFlags uint32, pvArg uintptr, pfnEnum PFN_CERT_ENUM_SYSTEM_STORE_LOCATION) bool

// TODO: Unknown type(s): CRYPT_ATTRIBUTE*, PCRYPT_ATTRIBUTE
// func CertFindAttribute(pszObjId /*const*/ LPCSTR, cAttr uint32, rgAttr CRYPT_ATTRIBUTE*) PCRYPT_ATTRIBUTE

func CertFindCRLInStore(hCertStore HCERTSTORE, dwCertEncodingType uint32, dwFindFlags uint32, dwFindType uint32, pvFindPara /*const*/ uintptr, pPrevCrlContext PCCRL_CONTEXT) PCCRL_CONTEXT {
	ret1 := syscall6(certFindCRLInStore, 6,
		uintptr(hCertStore),
		uintptr(dwCertEncodingType),
		uintptr(dwFindFlags),
		uintptr(dwFindType),
		pvFindPara,
		uintptr(unsafe.Pointer(pPrevCrlContext)))
	return (PCCRL_CONTEXT)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): PCCTL_CONTEXT
// func CertFindCTLInStore(hCertStore HCERTSTORE, dwMsgAndCertEncodingType uint32, dwFindFlags uint32, dwFindType uint32, pvFindPara /*const*/ uintptr, pPrevCtlContext PCCTL_CONTEXT) PCCTL_CONTEXT

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertFindCertificateInCRL(pCert PCCERT_CONTEXT, pCrlContext PCCRL_CONTEXT, dwFlags uint32, pvReserved uintptr, ppCrlEntry *PCRL_ENTRY) bool

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertFindCertificateInStore(hCertStore HCERTSTORE, dwCertEncodingType uint32, dwFindFlags uint32, dwFindType uint32, pvFindPara /*const*/ uintptr, pPrevCertContext PCCERT_CONTEXT) PCCERT_CONTEXT

// TODO: Unknown type(s): PCCERT_CHAIN_CONTEXT
// func CertFindChainInStore(hCertStore HCERTSTORE, dwCertEncodingType uint32, dwFindFlags uint32, dwFindType uint32, pvFindPara /*const*/ uintptr, pPrevChainContext PCCERT_CHAIN_CONTEXT) PCCERT_CHAIN_CONTEXT

func CertFindExtension(pszObjId /*const*/ LPCSTR, cExtensions uint32, rgExtensions *CERT_EXTENSION) PCERT_EXTENSION {
	ret1 := syscall3(certFindExtension, 3,
		uintptr(unsafe.Pointer(pszObjId)),
		uintptr(cExtensions),
		uintptr(unsafe.Pointer(rgExtensions)))
	return (PCERT_EXTENSION)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): PCERT_NAME_INFO, PCERT_RDN_ATTR
// func CertFindRDNAttr(pszObjId /*const*/ LPCSTR, pName PCERT_NAME_INFO) PCERT_RDN_ATTR

// TODO: Unknown type(s): PCCTL_CONTEXT, PCTL_ENTRY
// func CertFindSubjectInCTL(dwEncodingType uint32, dwSubjectType uint32, pvSubject uintptr, pCtlContext PCCTL_CONTEXT, dwFlags uint32) PCTL_ENTRY

// TODO: Unknown type(s): PCCTL_CONTEXT
// func CertFindSubjectInSortedCTL(pSubjectIdentifier PCRYPT_DATA_BLOB, pCtlContext PCCTL_CONTEXT, dwFlags uint32, pvReserved uintptr, pEncodedAttributes PCRYPT_DER_BLOB) bool

func CertFreeCRLContext(pCrlContext PCCRL_CONTEXT) bool {
	ret1 := syscall3(certFreeCRLContext, 1,
		uintptr(unsafe.Pointer(pCrlContext)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCCTL_CONTEXT
// func CertFreeCTLContext(pCtlContext PCCTL_CONTEXT) bool

// TODO: Unknown type(s): PCCERT_CHAIN_CONTEXT
// func CertFreeCertificateChain(pChainContext PCCERT_CHAIN_CONTEXT)

// TODO: Unknown type(s): HCERTCHAINENGINE
// func CertFreeCertificateChainEngine(hChainEngine HCERTCHAINENGINE)

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertFreeCertificateContext(pCertContext PCCERT_CONTEXT) bool

func CertGetCRLContextProperty(pCrlContext PCCRL_CONTEXT, dwPropId uint32, pvData uintptr, pcbData *uint32) bool {
	ret1 := syscall6(certGetCRLContextProperty, 4,
		uintptr(unsafe.Pointer(pCrlContext)),
		uintptr(dwPropId),
		pvData,
		uintptr(unsafe.Pointer(pcbData)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertGetCRLFromStore(hCertStore HCERTSTORE, pIssuerContext PCCERT_CONTEXT, pPrevCrlContext PCCRL_CONTEXT, pdwFlags *uint32) PCCRL_CONTEXT

// TODO: Unknown type(s): PCCTL_CONTEXT
// func CertGetCTLContextProperty(pCtlContext PCCTL_CONTEXT, dwPropId uint32, pvData uintptr, pcbData *uint32) bool

// TODO: Unknown type(s): HCERTCHAINENGINE, PCCERT_CHAIN_CONTEXT *, PCCERT_CONTEXT, PCERT_CHAIN_PARA
// func CertGetCertificateChain(hChainEngine HCERTCHAINENGINE, pCertContext PCCERT_CONTEXT, pTime *FILETIME, hAdditionalStore HCERTSTORE, pChainPara PCERT_CHAIN_PARA, dwFlags uint32, pvReserved LPVOID, ppChainContext PCCERT_CHAIN_CONTEXT *) bool

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertGetCertificateContextProperty(pCertContext PCCERT_CONTEXT, dwPropId uint32, pvData uintptr, pcbData *uint32) bool

// TODO: Unknown type(s): PCCERT_CONTEXT, PCERT_ENHKEY_USAGE
// func CertGetEnhancedKeyUsage(pCertContext PCCERT_CONTEXT, dwFlags uint32, pUsage PCERT_ENHKEY_USAGE, pcbUsage *uint32) bool

// TODO: Unknown type(s): PCERT_INFO
// func CertGetIntendedKeyUsage(dwCertEncodingType uint32, pCertInfo PCERT_INFO, pbKeyUsage *byte, cbKeyUsage uint32) bool

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertGetIssuerCertificateFromStore(hCertStore HCERTSTORE, pSubjectContext PCCERT_CONTEXT, pPrevIssuerContext PCCERT_CONTEXT, pdwFlags *uint32) PCCERT_CONTEXT

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertGetNameString(pCertContext PCCERT_CONTEXT, dwType uint32, dwFlags uint32, pvTypePara uintptr, pszNameString LPWSTR, cchNameString uint32) uint32

// TODO: Unknown type(s): PCERT_PUBLIC_KEY_INFO
// func CertGetPublicKeyLength(dwCertEncodingType uint32, pPublicKey PCERT_PUBLIC_KEY_INFO) uint32

func CertGetStoreProperty(hCertStore HCERTSTORE, dwPropId uint32, pvData uintptr, pcbData *uint32) bool {
	ret1 := syscall6(certGetStoreProperty, 4,
		uintptr(hCertStore),
		uintptr(dwPropId),
		pvData,
		uintptr(unsafe.Pointer(pcbData)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCCERT_CONTEXT, PCERT_INFO
// func CertGetSubjectCertificateFromStore(hCertStore HCERTSTORE, dwCertEncodingType uint32, pCertId PCERT_INFO) PCCERT_CONTEXT

// TODO: Unknown type(s): PCCERT_CONTEXT *
// func CertGetValidUsages(cCerts uint32, rghCerts PCCERT_CONTEXT *, cNumOIDs *int, rghOIDs *LPSTR, pcbOIDs *uint32) bool

// TODO: Unknown type(s): PCERT_RDN
// func CertIsRDNAttrsInCertificateName(dwCertEncodingType uint32, dwFlags uint32, pCertName PCERT_NAME_BLOB, pRDN PCERT_RDN) bool

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertIsValidCRLForCertificate(pCert PCCERT_CONTEXT, pCrl PCCRL_CONTEXT, dwFlags uint32, pvReserved uintptr) bool

func CertNameToStr(dwCertEncodingType uint32, pName PCERT_NAME_BLOB, dwStrType uint32, psz LPWSTR, csz uint32) uint32 {
	ret1 := syscall6(certNameToStr, 5,
		uintptr(dwCertEncodingType),
		uintptr(unsafe.Pointer(pName)),
		uintptr(dwStrType),
		uintptr(unsafe.Pointer(psz)),
		uintptr(csz),
		0)
	return uint32(ret1)
}

func CertOIDToAlgId(pszObjId /*const*/ LPCSTR) uint32 {
	ret1 := syscall3(certOIDToAlgId, 1,
		uintptr(unsafe.Pointer(pszObjId)),
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): HCRYPTPROV_LEGACY
// func CertOpenStore(lpszStoreProvider /*const*/ LPCSTR, dwEncodingType uint32, hCryptProv HCRYPTPROV_LEGACY, dwFlags uint32, pvPara /*const*/ uintptr) HCERTSTORE

// TODO: Unknown type(s): HCRYPTPROV_LEGACY
// func CertOpenSystemStore(hProv HCRYPTPROV_LEGACY, szSubsystemProtocol string) HCERTSTORE

func CertRDNValueToStr(dwValueType uint32, pValue PCERT_RDN_VALUE_BLOB, psz LPWSTR, csz uint32) uint32 {
	ret1 := syscall6(certRDNValueToStr, 4,
		uintptr(dwValueType),
		uintptr(unsafe.Pointer(pValue)),
		uintptr(unsafe.Pointer(psz)),
		uintptr(csz),
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PCERT_PHYSICAL_STORE_INFO
// func CertRegisterPhysicalStore(pvSystemStore /*const*/ uintptr, dwFlags uint32, pwszStoreName string, pStoreInfo PCERT_PHYSICAL_STORE_INFO, pvReserved uintptr) bool

// TODO: Unknown type(s): PCERT_SYSTEM_STORE_INFO
// func CertRegisterSystemStore(pvSystemStore /*const*/ uintptr, dwFlags uint32, pStoreInfo PCERT_SYSTEM_STORE_INFO, pvReserved uintptr) bool

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertRemoveEnhancedKeyUsageIdentifier(pCertContext PCCERT_CONTEXT, pszUsageIdentifier /*const*/ LPCSTR) bool

func CertRemoveStoreFromCollection(hCollectionStore HCERTSTORE, hSiblingStore HCERTSTORE) {
	syscall3(certRemoveStoreFromCollection, 2,
		uintptr(hCollectionStore),
		uintptr(hSiblingStore),
		0)
}

// TODO: Unknown type(s): HCERTCHAINENGINE
// func CertResyncCertificateChainEngine(hChainEngine HCERTCHAINENGINE) bool

func CertSaveStore(hCertStore HCERTSTORE, dwEncodingType uint32, dwSaveAs uint32, dwSaveTo uint32, pvSaveToPara uintptr, dwFlags uint32) bool {
	ret1 := syscall6(certSaveStore, 6,
		uintptr(hCertStore),
		uintptr(dwEncodingType),
		uintptr(dwSaveAs),
		uintptr(dwSaveTo),
		pvSaveToPara,
		uintptr(dwFlags))
	return ret1 != 0
}

func CertSerializeCRLStoreElement(pCrlContext PCCRL_CONTEXT, dwFlags uint32, pbElement *byte, pcbElement *uint32) bool {
	ret1 := syscall6(certSerializeCRLStoreElement, 4,
		uintptr(unsafe.Pointer(pCrlContext)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbElement)),
		uintptr(unsafe.Pointer(pcbElement)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCCTL_CONTEXT
// func CertSerializeCTLStoreElement(pCtlContext PCCTL_CONTEXT, dwFlags uint32, pbElement *byte, pcbElement *uint32) bool

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertSerializeCertificateStoreElement(pCertContext PCCERT_CONTEXT, dwFlags uint32, pbElement *byte, pcbElement *uint32) bool

func CertSetCRLContextProperty(pCrlContext PCCRL_CONTEXT, dwPropId uint32, dwFlags uint32, pvData /*const*/ uintptr) bool {
	ret1 := syscall6(certSetCRLContextProperty, 4,
		uintptr(unsafe.Pointer(pCrlContext)),
		uintptr(dwPropId),
		uintptr(dwFlags),
		pvData,
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCCTL_CONTEXT
// func CertSetCTLContextProperty(pCtlContext PCCTL_CONTEXT, dwPropId uint32, dwFlags uint32, pvData /*const*/ uintptr) bool

// TODO: Unknown type(s): PCCERT_CONTEXT, PCTL_ENTRY
// func CertSetCertificateContextPropertiesFromCTLEntry(pCertContext PCCERT_CONTEXT, pCtlEntry PCTL_ENTRY, dwFlags uint32) bool

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertSetCertificateContextProperty(pCertContext PCCERT_CONTEXT, dwPropId uint32, dwFlags uint32, pvData /*const*/ uintptr) bool

// TODO: Unknown type(s): PCCERT_CONTEXT, PCERT_ENHKEY_USAGE
// func CertSetEnhancedKeyUsage(pCertContext PCCERT_CONTEXT, pUsage PCERT_ENHKEY_USAGE) bool

func CertSetStoreProperty(hCertStore HCERTSTORE, dwPropId uint32, dwFlags uint32, pvData /*const*/ uintptr) bool {
	ret1 := syscall6(certSetStoreProperty, 4,
		uintptr(hCertStore),
		uintptr(dwPropId),
		uintptr(dwFlags),
		pvData,
		0,
		0)
	return ret1 != 0
}

func CertStrToName(dwCertEncodingType uint32, pszX500 string, dwStrType uint32, pvReserved uintptr, pbEncoded *byte, pcbEncoded *uint32, ppszError *LPCWSTR) bool {
	pszX500Str := unicode16FromString(pszX500)
	ret1 := syscall9(certStrToName, 7,
		uintptr(dwCertEncodingType),
		uintptr(unsafe.Pointer(&pszX500Str[0])),
		uintptr(dwStrType),
		pvReserved,
		uintptr(unsafe.Pointer(pbEncoded)),
		uintptr(unsafe.Pointer(pcbEncoded)),
		uintptr(unsafe.Pointer(ppszError)),
		0,
		0)
	return ret1 != 0
}

func CertUnregisterPhysicalStore(pvSystemStore /*const*/ uintptr, dwFlags uint32, pwszStoreName string) bool {
	pwszStoreNameStr := unicode16FromString(pwszStoreName)
	ret1 := syscall3(certUnregisterPhysicalStore, 3,
		pvSystemStore,
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(&pwszStoreNameStr[0])))
	return ret1 != 0
}

func CertUnregisterSystemStore(pvSystemStore /*const*/ uintptr, dwFlags uint32) bool {
	ret1 := syscall3(certUnregisterSystemStore, 2,
		pvSystemStore,
		uintptr(dwFlags),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCERT_INFO
// func CertVerifyCRLRevocation(dwCertEncodingType uint32, pCertId PCERT_INFO, cCrlInfo uint32, rgpCrlInfo *PCRL_INFO) bool

func CertVerifyCRLTimeValidity(pTimeToVerify *FILETIME, pCrlInfo PCRL_INFO) LONG {
	ret1 := syscall3(certVerifyCRLTimeValidity, 2,
		uintptr(unsafe.Pointer(pTimeToVerify)),
		uintptr(unsafe.Pointer(pCrlInfo)),
		0)
	return LONG(ret1)
}

// TODO: Unknown type(s): PCTL_USAGE, PCTL_VERIFY_USAGE_PARA, PCTL_VERIFY_USAGE_STATUS
// func CertVerifyCTLUsage(dwEncodingType uint32, dwSubjectType uint32, pvSubject uintptr, pSubjectUsage PCTL_USAGE, dwFlags uint32, pVerifyUsagePara PCTL_VERIFY_USAGE_PARA, pVerifyUsageStatus PCTL_VERIFY_USAGE_STATUS) bool

// TODO: Unknown type(s): PCCERT_CHAIN_CONTEXT, PCERT_CHAIN_POLICY_PARA, PCERT_CHAIN_POLICY_STATUS
// func CertVerifyCertificateChainPolicy(pszPolicyOID /*const*/ LPCSTR, pChainContext PCCERT_CHAIN_CONTEXT, pPolicyPara PCERT_CHAIN_POLICY_PARA, pPolicyStatus PCERT_CHAIN_POLICY_STATUS) bool

// TODO: Unknown type(s): PCERT_REVOCATION_PARA, PCERT_REVOCATION_STATUS
// func CertVerifyRevocation(dwEncodingType uint32, dwRevType uint32, cContext uint32, rgpvContext *PVOID, dwFlags uint32, pRevPara PCERT_REVOCATION_PARA, pRevStatus PCERT_REVOCATION_STATUS) bool

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CertVerifySubjectCertificateContext(pSubject PCCERT_CONTEXT, pIssuer PCCERT_CONTEXT, pdwFlags *uint32) bool

// TODO: Unknown type(s): PCERT_INFO
// func CertVerifyTimeValidity(pTimeToVerify *FILETIME, pCertInfo PCERT_INFO) LONG

// TODO: Unknown type(s): PCERT_INFO
// func CertVerifyValidityNesting(pSubjectInfo PCERT_INFO, pIssuerInfo PCERT_INFO) bool

// TODO: Unknown type(s): HCRYPTPROV_OR_NCRYPT_KEY_HANDLE *, PCCERT_CONTEXT
// func CryptAcquireCertificatePrivateKey(pCert PCCERT_CONTEXT, dwFlags uint32, pvParameters uintptr, phCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE *, pdwKeySpec *uint32, pfCallerFreeProvOrNCryptKey *BOOL) bool

func CryptBinaryToString(pbBinary /*const*/ *byte, cbBinary uint32, dwFlags uint32, pszString LPWSTR, pcchString *uint32) bool {
	ret1 := syscall6(cryptBinaryToString, 5,
		uintptr(unsafe.Pointer(pbBinary)),
		uintptr(cbBinary),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pszString)),
		uintptr(unsafe.Pointer(pcchString)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): HCRYPTASYNC
// func CryptCloseAsyncHandle(hAsync HCRYPTASYNC) bool

// TODO: Unknown type(s): PHCRYPTASYNC
// func CryptCreateAsyncHandle(dwFlags uint32, phAsync PHCRYPTASYNC) bool

// TODO: Unknown type(s): const PUBLICKEYSTRUC *
// func CryptCreateKeyIdentifierFromCSP(dwCertEncodingType uint32, pszPubKeyOID /*const*/ LPCSTR, pPubKeyStruc /*const*/ const PUBLICKEYSTRUC *, cbPubKeyStruc uint32, dwFlags uint32, pvReserved uintptr, pbHash *byte, pcbHash *uint32) bool

// TODO: Unknown type(s): PCCERT_CONTEXT *, PCRYPT_DECRYPT_MESSAGE_PARA, PCRYPT_VERIFY_MESSAGE_PARA
// func CryptDecodeMessage(dwMsgTypeFlags uint32, pDecryptPara PCRYPT_DECRYPT_MESSAGE_PARA, pVerifyPara PCRYPT_VERIFY_MESSAGE_PARA, dwSignerIndex uint32, pbEncodedBlob /*const*/ *byte, cbEncodedBlob uint32, dwPrevInnerContentType uint32, pdwMsgType *uint32, pdwInnerContentType *uint32, pbDecoded *byte, pcbDecoded *uint32, ppXchgCert PCCERT_CONTEXT *, ppSignerCert PCCERT_CONTEXT *) bool

func CryptDecodeObject(dwCertEncodingType uint32, lpszStructType /*const*/ LPCSTR, pbEncoded /*const*/ *byte, cbEncoded uint32, dwFlags uint32, pvStructInfo uintptr, pcbStructInfo *uint32) bool {
	ret1 := syscall9(cryptDecodeObject, 7,
		uintptr(dwCertEncodingType),
		uintptr(unsafe.Pointer(lpszStructType)),
		uintptr(unsafe.Pointer(pbEncoded)),
		uintptr(cbEncoded),
		uintptr(dwFlags),
		pvStructInfo,
		uintptr(unsafe.Pointer(pcbStructInfo)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCRYPT_DECODE_PARA
// func CryptDecodeObjectEx(dwCertEncodingType uint32, lpszStructType /*const*/ LPCSTR, pbEncoded /*const*/ *byte, cbEncoded uint32, dwFlags uint32, pDecodePara PCRYPT_DECODE_PARA, pvStructInfo uintptr, pcbStructInfo *uint32) bool

// TODO: Unknown type(s): PCCERT_CONTEXT *, PCRYPT_DECRYPT_MESSAGE_PARA, PCRYPT_VERIFY_MESSAGE_PARA
// func CryptDecryptAndVerifyMessageSignature(pDecryptPara PCRYPT_DECRYPT_MESSAGE_PARA, pVerifyPara PCRYPT_VERIFY_MESSAGE_PARA, dwSignerIndex uint32, pbEncryptedBlob /*const*/ *byte, cbEncryptedBlob uint32, pbDecrypted *byte, pcbDecrypted *uint32, ppXchgCert PCCERT_CONTEXT *, ppSignerCert PCCERT_CONTEXT *) bool

// TODO: Unknown type(s): PCCERT_CONTEXT *, PCRYPT_DECRYPT_MESSAGE_PARA
// func CryptDecryptMessage(pDecryptPara PCRYPT_DECRYPT_MESSAGE_PARA, pbEncryptedBlob /*const*/ *byte, cbEncryptedBlob uint32, pbDecrypted *byte, pcbDecrypted *uint32, ppXchgCert PCCERT_CONTEXT *) bool

func CryptEncodeObject(dwCertEncodingType uint32, lpszStructType /*const*/ LPCSTR, pvStructInfo /*const*/ uintptr, pbEncoded *byte, pcbEncoded *uint32) bool {
	ret1 := syscall6(cryptEncodeObject, 5,
		uintptr(dwCertEncodingType),
		uintptr(unsafe.Pointer(lpszStructType)),
		pvStructInfo,
		uintptr(unsafe.Pointer(pbEncoded)),
		uintptr(unsafe.Pointer(pcbEncoded)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCRYPT_ENCODE_PARA
// func CryptEncodeObjectEx(dwCertEncodingType uint32, lpszStructType /*const*/ LPCSTR, pvStructInfo /*const*/ uintptr, dwFlags uint32, pEncodePara PCRYPT_ENCODE_PARA, pvEncoded uintptr, pcbEncoded *uint32) bool

// TODO: Unknown type(s): PCCERT_CONTEXT*, PCRYPT_ENCRYPT_MESSAGE_PARA
// func CryptEncryptMessage(pEncryptPara PCRYPT_ENCRYPT_MESSAGE_PARA, cRecipientCert uint32, rgpRecipientCert PCCERT_CONTEXT*, pbToBeEncrypted /*const*/ *byte, cbToBeEncrypted uint32, pbEncryptedBlob *byte, pcbEncryptedBlob *uint32) bool

// TODO: Unknown type(s): PFN_CRYPT_ENUM_KEYID_PROP
// func CryptEnumKeyIdentifierProperties(pKeyIdentifier /*const*/ *CRYPT_HASH_BLOB, dwPropId uint32, dwFlags uint32, pwszComputerName string, pvReserved uintptr, pvArg uintptr, pfnEnum PFN_CRYPT_ENUM_KEYID_PROP) bool

// TODO: Unknown type(s): PFN_CRYPT_ENUM_OID_FUNC
// func CryptEnumOIDFunction(dwEncodingType uint32, pszFuncName /*const*/ LPCSTR, pszOID /*const*/ LPCSTR, dwFlags uint32, pvArg uintptr, pfnEnumOIDFunc PFN_CRYPT_ENUM_OID_FUNC) bool

// TODO: Unknown type(s): PFN_CRYPT_ENUM_OID_INFO
// func CryptEnumOIDInfo(dwGroupId uint32, dwFlags uint32, pvArg uintptr, pfnEnumOIDInfo PFN_CRYPT_ENUM_OID_INFO) bool

func CryptExportPKCS8(hCryptProv HCRYPTPROV, dwKeySpec uint32, pszPrivateKeyObjId LPSTR, dwFlags uint32, pvAuxInfo uintptr, pbPrivateKeyBlob *byte, pcbPrivateKeyBlob *uint32) bool {
	ret1 := syscall9(cryptExportPKCS8, 7,
		uintptr(hCryptProv),
		uintptr(dwKeySpec),
		uintptr(unsafe.Pointer(pszPrivateKeyObjId)),
		uintptr(dwFlags),
		pvAuxInfo,
		uintptr(unsafe.Pointer(pbPrivateKeyBlob)),
		uintptr(unsafe.Pointer(pcbPrivateKeyBlob)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, PCERT_PUBLIC_KEY_INFO
// func CryptExportPublicKeyInfo(hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, dwKeySpec uint32, dwCertEncodingType uint32, pInfo PCERT_PUBLIC_KEY_INFO, pcbInfo *uint32) bool

// TODO: Unknown type(s): HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, PCERT_PUBLIC_KEY_INFO
// func CryptExportPublicKeyInfoEx(hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, dwKeySpec uint32, dwCertEncodingType uint32, pszPublicKeyObjId LPSTR, dwFlags uint32, pvAuxInfo uintptr, pInfo PCERT_PUBLIC_KEY_INFO, pcbInfo *uint32) bool

// TODO: Unknown type(s): PCCERT_CONTEXT
// func CryptFindCertificateKeyProvInfo(pCert PCCERT_CONTEXT, dwFlags uint32, pvReserved uintptr) bool

func CryptFindLocalizedName(pwszCryptName string) string {
	pwszCryptNameStr := unicode16FromString(pwszCryptName)
	ret1 := syscall3(cryptFindLocalizedName, 1,
		uintptr(unsafe.Pointer(&pwszCryptNameStr[0])),
		0,
		0)
	return stringFromUnicode16((*uint16)(unsafe.Pointer(ret1)))
}

// TODO: Unknown type(s): PCCRYPT_OID_INFO
// func CryptFindOIDInfo(dwKeyType uint32, pvKey uintptr, dwGroupId uint32) PCCRYPT_OID_INFO

func CryptFormatObject(dwCertEncodingType uint32, dwFormatType uint32, dwFormatStrType uint32, pFormatStruct uintptr, lpszStructType /*const*/ LPCSTR, pbEncoded /*const*/ *byte, cbEncoded uint32, pbFormat uintptr, pcbFormat *uint32) bool {
	ret1 := syscall9(cryptFormatObject, 9,
		uintptr(dwCertEncodingType),
		uintptr(dwFormatType),
		uintptr(dwFormatStrType),
		pFormatStruct,
		uintptr(unsafe.Pointer(lpszStructType)),
		uintptr(unsafe.Pointer(pbEncoded)),
		uintptr(cbEncoded),
		pbFormat,
		uintptr(unsafe.Pointer(pcbFormat)))
	return ret1 != 0
}

// TODO: Unknown type(s): HCRYPTOIDFUNCADDR
// func CryptFreeOIDFunctionAddress(hFuncAddr HCRYPTOIDFUNCADDR, dwFlags uint32) bool

// TODO: Unknown type(s): HCRYPTASYNC, PFN_CRYPT_ASYNC_PARAM_FREE_FUNC *
// func CryptGetAsyncParam(hAsync HCRYPTASYNC, pszParamOid LPSTR, ppvParam *LPVOID, ppfnFree PFN_CRYPT_ASYNC_PARAM_FREE_FUNC *) bool

// TODO: Unknown type(s): HCRYPTOIDFUNCSET
// func CryptGetDefaultOIDDllList(hFuncSet HCRYPTOIDFUNCSET, dwEncodingType uint32, pwszDllList *WCHAR, pcchDllList *uint32) bool

// TODO: Unknown type(s): HCRYPTOIDFUNCADDR *, HCRYPTOIDFUNCSET
// func CryptGetDefaultOIDFunctionAddress(hFuncSet HCRYPTOIDFUNCSET, dwEncodingType uint32, pwszDll string, dwFlags uint32, ppvFuncAddr uintptr, phFuncAddr HCRYPTOIDFUNCADDR *) bool

func CryptGetKeyIdentifierProperty(pKeyIdentifier /*const*/ *CRYPT_HASH_BLOB, dwPropId uint32, dwFlags uint32, pwszComputerName string, pvReserved uintptr, pvData uintptr, pcbData *uint32) bool {
	pwszComputerNameStr := unicode16FromString(pwszComputerName)
	ret1 := syscall9(cryptGetKeyIdentifierProperty, 7,
		uintptr(unsafe.Pointer(pKeyIdentifier)),
		uintptr(dwPropId),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(&pwszComputerNameStr[0])),
		pvReserved,
		pvData,
		uintptr(unsafe.Pointer(pcbData)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): HCRYPTPROV_LEGACY
// func CryptGetMessageCertificates(dwMsgAndCertEncodingType uint32, hCryptProv HCRYPTPROV_LEGACY, dwFlags uint32, pbSignedBlob /*const*/ *byte, cbSignedBlob uint32) HCERTSTORE

func CryptGetMessageSignerCount(dwMsgEncodingType uint32, pbSignedBlob /*const*/ *byte, cbSignedBlob uint32) LONG {
	ret1 := syscall3(cryptGetMessageSignerCount, 3,
		uintptr(dwMsgEncodingType),
		uintptr(unsafe.Pointer(pbSignedBlob)),
		uintptr(cbSignedBlob))
	return LONG(ret1)
}

// TODO: Unknown type(s): HCRYPTOIDFUNCADDR *, HCRYPTOIDFUNCSET
// func CryptGetOIDFunctionAddress(hFuncSet HCRYPTOIDFUNCSET, dwEncodingType uint32, pszOID /*const*/ LPCSTR, dwFlags uint32, ppvFuncAddr uintptr, phFuncAddr HCRYPTOIDFUNCADDR *) bool

func CryptGetOIDFunctionValue(dwEncodingType uint32, pszFuncName /*const*/ LPCSTR, pszOID /*const*/ LPCSTR, pwszValueName string, pdwValueType *uint32, pbValueData *byte, pcbValueData *uint32) bool {
	pwszValueNameStr := unicode16FromString(pwszValueName)
	ret1 := syscall9(cryptGetOIDFunctionValue, 7,
		uintptr(dwEncodingType),
		uintptr(unsafe.Pointer(pszFuncName)),
		uintptr(unsafe.Pointer(pszOID)),
		uintptr(unsafe.Pointer(&pwszValueNameStr[0])),
		uintptr(unsafe.Pointer(pdwValueType)),
		uintptr(unsafe.Pointer(pbValueData)),
		uintptr(unsafe.Pointer(pcbValueData)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): HCRYPTPROV_LEGACY
// func CryptHashCertificate(hCryptProv HCRYPTPROV_LEGACY, algid ALG_ID, dwFlags uint32, pbEncoded /*const*/ *byte, cbEncoded uint32, pbComputedHash *byte, pcbComputedHash *uint32) bool

// TODO: Unknown type(s): PCRYPT_HASH_MESSAGE_PARA
// func CryptHashMessage(pHashPara PCRYPT_HASH_MESSAGE_PARA, fDetachedHash bool, cToBeHashed uint32, rgpbToBeHashed /*const*/ **BYTE, rgcbToBeHashed *uint32, pbHashedBlob *byte, pcbHashedBlob *uint32, pbComputedHash *byte, pcbComputedHash *uint32) bool

// TODO: Unknown type(s): HCRYPTPROV_LEGACY, PCERT_PUBLIC_KEY_INFO
// func CryptHashPublicKeyInfo(hCryptProv HCRYPTPROV_LEGACY, algid ALG_ID, dwFlags uint32, dwCertEncodingType uint32, pInfo PCERT_PUBLIC_KEY_INFO, pbComputedHash *byte, pcbComputedHash *uint32) bool

// TODO: Unknown type(s): HCRYPTPROV_LEGACY
// func CryptHashToBeSigned(hCryptProv HCRYPTPROV_LEGACY, dwCertEncodingType uint32, pbEncoded /*const*/ *byte, cbEncoded uint32, pbComputedHash *byte, pcbComputedHash *uint32) bool

// TODO: Unknown type(s): CRYPT_PKCS8_IMPORT_PARAMS
// func CryptImportPKCS8(sPrivateKeyAndParams CRYPT_PKCS8_IMPORT_PARAMS, dwFlags uint32, phCryptProv *HCRYPTPROV, pvAuxInfo uintptr) bool

// TODO: Unknown type(s): PCERT_PUBLIC_KEY_INFO
// func CryptImportPublicKeyInfo(hCryptProv HCRYPTPROV, dwCertEncodingType uint32, pInfo PCERT_PUBLIC_KEY_INFO, phKey *HCRYPTKEY) bool

// TODO: Unknown type(s): PCERT_PUBLIC_KEY_INFO
// func CryptImportPublicKeyInfoEx(hCryptProv HCRYPTPROV, dwCertEncodingType uint32, pInfo PCERT_PUBLIC_KEY_INFO, aiKeyAlg ALG_ID, dwFlags uint32, pvAuxInfo uintptr, phKey *HCRYPTKEY) bool

// TODO: Unknown type(s): HCRYPTOIDFUNCSET
// func CryptInitOIDFunctionSet(pszFuncName /*const*/ LPCSTR, dwFlags uint32) HCRYPTOIDFUNCSET

// TODO: Unknown type(s): HCRYPTDEFAULTCONTEXT *
// func CryptInstallDefaultContext(hCryptProv HCRYPTPROV, dwDefaultType uint32, pvDefaultPara /*const*/ uintptr, dwFlags uint32, pvReserved uintptr, phDefaultContext HCRYPTDEFAULTCONTEXT *) bool

// TODO: Unknown type(s): const CRYPT_OID_FUNC_ENTRY*
// func CryptInstallOIDFunctionAddress(hModule HMODULE, dwEncodingType uint32, pszFuncName /*const*/ LPCSTR, cFuncEntry uint32, rgFuncEntry /*const*/ const CRYPT_OID_FUNC_ENTRY*, dwFlags uint32) bool

func CryptMemAlloc(cbSize ULONG) LPVOID {
	ret1 := syscall3(cryptMemAlloc, 1,
		uintptr(cbSize),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func CryptMemFree(pv LPVOID) {
	syscall3(cryptMemFree, 1,
		uintptr(unsafe.Pointer(pv)),
		0,
		0)
}

func CryptMemRealloc(pv LPVOID, cbSize ULONG) LPVOID {
	ret1 := syscall3(cryptMemRealloc, 2,
		uintptr(unsafe.Pointer(pv)),
		uintptr(cbSize),
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func CryptMsgCalculateEncodedLength(dwMsgEncodingType uint32, dwFlags uint32, dwMsgType uint32, pvMsgEncodeInfo /*const*/ uintptr, pszInnerContentObjID LPSTR, cbData uint32) uint32 {
	ret1 := syscall6(cryptMsgCalculateEncodedLength, 6,
		uintptr(dwMsgEncodingType),
		uintptr(dwFlags),
		uintptr(dwMsgType),
		pvMsgEncodeInfo,
		uintptr(unsafe.Pointer(pszInnerContentObjID)),
		uintptr(cbData))
	return uint32(ret1)
}

// TODO: Unknown type(s): HCRYPTMSG
// func CryptMsgClose(hCryptMsg HCRYPTMSG) bool

// TODO: Unknown type(s): HCRYPTMSG
// func CryptMsgControl(hCryptMsg HCRYPTMSG, dwFlags uint32, dwCtrlType uint32, pvCtrlPara /*const*/ uintptr) bool

// TODO: Unknown type(s): HCRYPTMSG, PCMSG_SIGNER_ENCODE_INFO
// func CryptMsgCountersign(hCryptMsg HCRYPTMSG, dwIndex uint32, cCountersigners uint32, rgCountersigners PCMSG_SIGNER_ENCODE_INFO) bool

// TODO: Unknown type(s): PCMSG_SIGNER_ENCODE_INFO
// func CryptMsgCountersignEncoded(dwEncodingType uint32, pbSignerInfo *byte, cbSignerInfo uint32, cCountersigners uint32, rgCountersigners PCMSG_SIGNER_ENCODE_INFO, pbCountersignature *byte, pcbCountersignature *DWORD) bool

// TODO: Unknown type(s): HCRYPTMSG
// func CryptMsgDuplicate(hCryptMsg HCRYPTMSG) HCRYPTMSG

// TODO: Unknown type(s): PCMSG_SIGNED_ENCODE_INFO, PCTL_INFO
// func CryptMsgEncodeAndSignCTL(dwMsgEncodingType uint32, pCtlInfo PCTL_INFO, pSignInfo PCMSG_SIGNED_ENCODE_INFO, dwFlags uint32, pbEncoded *byte, pcbEncoded *uint32) bool

// TODO: Unknown type(s): HCRYPTMSG, PCCERT_CONTEXT *
// func CryptMsgGetAndVerifySigner(hCryptMsg HCRYPTMSG, cSignerStore uint32, rghSignerStore *HCERTSTORE, dwFlags uint32, ppSigner PCCERT_CONTEXT *, pdwSignerIndex *uint32) bool

// TODO: Unknown type(s): HCRYPTMSG
// func CryptMsgGetParam(hCryptMsg HCRYPTMSG, dwParamType uint32, dwIndex uint32, pvData uintptr, pcbData *uint32) bool

// TODO: Unknown type(s): HCRYPTMSG, HCRYPTPROV_LEGACY, PCERT_INFO, PCMSG_STREAM_INFO
// func CryptMsgOpenToDecode(dwMsgEncodingType uint32, dwFlags uint32, dwMsgType uint32, hCryptProv HCRYPTPROV_LEGACY, pRecipientInfo PCERT_INFO, pStreamInfo PCMSG_STREAM_INFO) HCRYPTMSG

// TODO: Unknown type(s): HCRYPTMSG, PCMSG_STREAM_INFO
// func CryptMsgOpenToEncode(dwMsgEncodingType uint32, dwFlags uint32, dwMsgType uint32, pvMsgEncodeInfo /*const*/ uintptr, pszInnerContentObjID LPSTR, pStreamInfo PCMSG_STREAM_INFO) HCRYPTMSG

// TODO: Unknown type(s): PCMSG_SIGNED_ENCODE_INFO
// func CryptMsgSignCTL(dwMsgEncodingType uint32, pbCtlContent *byte, cbCtlContent uint32, pSignInfo PCMSG_SIGNED_ENCODE_INFO, dwFlags uint32, pbEncoded *byte, pcbEncoded *uint32) bool

// TODO: Unknown type(s): HCRYPTMSG
// func CryptMsgUpdate(hCryptMsg HCRYPTMSG, pbData /*const*/ *byte, cbData uint32, fFinal bool) bool

// TODO: Unknown type(s): HCRYPTPROV_LEGACY, PCERT_INFO
// func CryptMsgVerifyCountersignatureEncoded(hCryptProv HCRYPTPROV_LEGACY, dwEncodingType uint32, pbSignerInfo *byte, cbSignerInfo uint32, pbSignerInfoCountersignature *byte, cbSignerInfoCountersignature uint32, pciCountersigner PCERT_INFO) bool

// TODO: Unknown type(s): HCRYPTPROV_LEGACY
// func CryptMsgVerifyCountersignatureEncodedEx(hCryptProv HCRYPTPROV_LEGACY, dwEncodingType uint32, pbSignerInfo *byte, cbSignerInfo uint32, pbSignerInfoCountersignature *byte, cbSignerInfoCountersignature uint32, dwSignerType uint32, pvSigner uintptr, dwFlags uint32, pvExtra uintptr) bool

// TODO: Unknown type(s): CRYPTPROTECT_PROMPTSTRUCT *
// func CryptProtectData(pDataIn *DATA_BLOB, szDataDescr string, pOptionalEntropy *DATA_BLOB, pvReserved uintptr, pPromptStruct CRYPTPROTECT_PROMPTSTRUCT *, dwFlags uint32, pDataOut *DATA_BLOB) bool

func CryptProtectMemory(pDataIn LPVOID, cbDataIn uint32, dwFlags uint32) bool {
	ret1 := syscall3(cryptProtectMemory, 3,
		uintptr(unsafe.Pointer(pDataIn)),
		uintptr(cbDataIn),
		uintptr(dwFlags))
	return ret1 != 0
}

// TODO: Unknown type(s): HCRYPTMSG *
// func CryptQueryObject(dwObjectType uint32, pvObject /*const*/ uintptr, dwExpectedContentTypeFlags uint32, dwExpectedFormatTypeFlags uint32, dwFlags uint32, pdwMsgAndCertEncodingType *uint32, pdwContentType *uint32, pdwFormatType *uint32, phCertStore *HCERTSTORE, phMsg HCRYPTMSG *, ppvContext /*const*/ uintptr) bool

func CryptRegisterDefaultOIDFunction(dwEncodingType uint32, pszFuncName /*const*/ LPCSTR, dwIndex uint32, pwszDll string) bool {
	pwszDllStr := unicode16FromString(pwszDll)
	ret1 := syscall6(cryptRegisterDefaultOIDFunction, 4,
		uintptr(dwEncodingType),
		uintptr(unsafe.Pointer(pszFuncName)),
		uintptr(dwIndex),
		uintptr(unsafe.Pointer(&pwszDllStr[0])),
		0,
		0)
	return ret1 != 0
}

func CryptRegisterOIDFunction(dwEncodingType uint32, pszFuncName /*const*/ LPCSTR, pszOID /*const*/ LPCSTR, pwszDll string, pszOverrideFuncName /*const*/ LPCSTR) bool {
	pwszDllStr := unicode16FromString(pwszDll)
	ret1 := syscall6(cryptRegisterOIDFunction, 5,
		uintptr(dwEncodingType),
		uintptr(unsafe.Pointer(pszFuncName)),
		uintptr(unsafe.Pointer(pszOID)),
		uintptr(unsafe.Pointer(&pwszDllStr[0])),
		uintptr(unsafe.Pointer(pszOverrideFuncName)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCCRYPT_OID_INFO
// func CryptRegisterOIDInfo(pInfo PCCRYPT_OID_INFO, dwFlags uint32) bool

// TODO: Unknown type(s): SIP_ADD_NEWPROVIDER *
// func CryptSIPAddProvider(psNewProv SIP_ADD_NEWPROVIDER *) bool

// TODO: Unknown type(s): SIP_INDIRECT_DATA *, SIP_SUBJECTINFO *
// func CryptSIPCreateIndirectData(pSubjectInfo SIP_SUBJECTINFO *, pcbIndirectData *uint32, pIndirectData SIP_INDIRECT_DATA *) bool

// TODO: Unknown type(s): SIP_SUBJECTINFO *
// func CryptSIPGetSignedDataMsg(pSubjectInfo SIP_SUBJECTINFO *, pdwEncodingType *uint32, dwIndex uint32, pcbSignedDataMsg *uint32, pbSignedDataMsg *byte) bool

// TODO: Unknown type(s): SIP_DISPATCH_INFO *
// func CryptSIPLoad(pgSubject /*const*/ *GUID, dwFlags uint32, pSipDispatch SIP_DISPATCH_INFO *) bool

// TODO: Unknown type(s): SIP_SUBJECTINFO *
// func CryptSIPPutSignedDataMsg(pSubjectInfo SIP_SUBJECTINFO *, dwEncodingType uint32, pdwIndex *uint32, cbSignedDataMsg uint32, pbSignedDataMsg *byte) bool

func CryptSIPRemoveProvider(pgProv *GUID) bool {
	ret1 := syscall3(cryptSIPRemoveProvider, 1,
		uintptr(unsafe.Pointer(pgProv)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): SIP_SUBJECTINFO *
// func CryptSIPRemoveSignedDataMsg(pSubjectInfo SIP_SUBJECTINFO *, dwIndex uint32) bool

func CryptSIPRetrieveSubjectGuid(fileName string, hFileIn HANDLE, pgSubject *GUID) bool {
	fileNameStr := unicode16FromString(fileName)
	ret1 := syscall3(cryptSIPRetrieveSubjectGuid, 3,
		uintptr(unsafe.Pointer(&fileNameStr[0])),
		uintptr(hFileIn),
		uintptr(unsafe.Pointer(pgSubject)))
	return ret1 != 0
}

func CryptSIPRetrieveSubjectGuidForCatalogFile(fileName string, hFileIn HANDLE, pgSubject *GUID) bool {
	fileNameStr := unicode16FromString(fileName)
	ret1 := syscall3(cryptSIPRetrieveSubjectGuidForCatalogFile, 3,
		uintptr(unsafe.Pointer(&fileNameStr[0])),
		uintptr(hFileIn),
		uintptr(unsafe.Pointer(pgSubject)))
	return ret1 != 0
}

// TODO: Unknown type(s): SIP_INDIRECT_DATA *, SIP_SUBJECTINFO *
// func CryptSIPVerifyIndirectData(pSubjectInfo SIP_SUBJECTINFO *, pIndirectData SIP_INDIRECT_DATA *) bool

// TODO: Unknown type(s): HCRYPTASYNC, PFN_CRYPT_ASYNC_PARAM_FREE_FUNC
// func CryptSetAsyncParam(hAsync HCRYPTASYNC, pszParamOid LPSTR, pvParam LPVOID, pfnFree PFN_CRYPT_ASYNC_PARAM_FREE_FUNC) bool

func CryptSetKeyIdentifierProperty(pKeyIdentifier /*const*/ *CRYPT_HASH_BLOB, dwPropId uint32, dwFlags uint32, pwszComputerName string, pvReserved uintptr, pvData /*const*/ uintptr) bool {
	pwszComputerNameStr := unicode16FromString(pwszComputerName)
	ret1 := syscall6(cryptSetKeyIdentifierProperty, 6,
		uintptr(unsafe.Pointer(pKeyIdentifier)),
		uintptr(dwPropId),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(&pwszComputerNameStr[0])),
		pvReserved,
		pvData)
	return ret1 != 0
}

func CryptSetOIDFunctionValue(dwEncodingType uint32, pszFuncName /*const*/ LPCSTR, pszOID /*const*/ LPCSTR, pwszValueName string, dwValueType uint32, pbValueData /*const*/ *byte, cbValueData uint32) bool {
	pwszValueNameStr := unicode16FromString(pwszValueName)
	ret1 := syscall9(cryptSetOIDFunctionValue, 7,
		uintptr(dwEncodingType),
		uintptr(unsafe.Pointer(pszFuncName)),
		uintptr(unsafe.Pointer(pszOID)),
		uintptr(unsafe.Pointer(&pwszValueNameStr[0])),
		uintptr(dwValueType),
		uintptr(unsafe.Pointer(pbValueData)),
		uintptr(cbValueData),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, PCRYPT_ALGORITHM_IDENTIFIER
// func CryptSignAndEncodeCertificate(hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, dwKeySpec uint32, dwCertEncodingType uint32, lpszStructType /*const*/ LPCSTR, pvStructInfo /*const*/ uintptr, pSignatureAlgorithm PCRYPT_ALGORITHM_IDENTIFIER, pvHashAuxInfo /*const*/ uintptr, pbEncoded *byte, pcbEncoded *uint32) bool

// TODO: Unknown type(s): PCCERT_CONTEXT*, PCRYPT_ENCRYPT_MESSAGE_PARA, PCRYPT_SIGN_MESSAGE_PARA
// func CryptSignAndEncryptMessage(pSignPara PCRYPT_SIGN_MESSAGE_PARA, pEncryptPara PCRYPT_ENCRYPT_MESSAGE_PARA, cRecipientCert uint32, rgpRecipientCert PCCERT_CONTEXT*, pbToBeSignedAndEncrypted /*const*/ *byte, cbToBeSignedAndEncrypted uint32, pbSignedAndEncryptedBlob *byte, pcbSignedAndEncryptedBlob *uint32) bool

// TODO: Unknown type(s): HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, PCRYPT_ALGORITHM_IDENTIFIER
// func CryptSignCertificate(hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, dwKeySpec uint32, dwCertEncodingType uint32, pbEncodedToBeSigned /*const*/ *byte, cbEncodedToBeSigned uint32, pSignatureAlgorithm PCRYPT_ALGORITHM_IDENTIFIER, pvHashAuxInfo /*const*/ uintptr, pbSignature *byte, pcbSignature *uint32) bool

// TODO: Unknown type(s): PCRYPT_SIGN_MESSAGE_PARA
// func CryptSignMessage(pSignPara PCRYPT_SIGN_MESSAGE_PARA, fDetachedSignature bool, cToBeSigned uint32, rgpbToBeSigned /*const*/ **BYTE, rgcbToBeSigned *uint32, pbSignedBlob *byte, pcbSignedBlob *uint32) bool

// TODO: Unknown type(s): PCRYPT_KEY_SIGN_MESSAGE_PARA
// func CryptSignMessageWithKey(pSignPara PCRYPT_KEY_SIGN_MESSAGE_PARA, pbToBeSigned /*const*/ *byte, cbToBeSigned uint32, pbSignedBlob *byte, pcbSignedBlob *uint32) bool

func CryptStringToBinary(pszString string, cchString uint32, dwFlags uint32, pbBinary *byte, pcbBinary *uint32, pdwSkip *uint32, pdwFlags *uint32) bool {
	pszStringStr := unicode16FromString(pszString)
	ret1 := syscall9(cryptStringToBinary, 7,
		uintptr(unsafe.Pointer(&pszStringStr[0])),
		uintptr(cchString),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbBinary)),
		uintptr(unsafe.Pointer(pcbBinary)),
		uintptr(unsafe.Pointer(pdwSkip)),
		uintptr(unsafe.Pointer(pdwFlags)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): HCRYPTDEFAULTCONTEXT
// func CryptUninstallDefaultContext(hDefaultContext HCRYPTDEFAULTCONTEXT, dwFlags uint32, pvReserved uintptr) bool

// TODO: Unknown type(s): CRYPTPROTECT_PROMPTSTRUCT *
// func CryptUnprotectData(pDataIn *DATA_BLOB, ppszDataDescr *LPWSTR, pOptionalEntropy *DATA_BLOB, pvReserved uintptr, pPromptStruct CRYPTPROTECT_PROMPTSTRUCT *, dwFlags uint32, pDataOut *DATA_BLOB) bool

func CryptUnprotectMemory(pDataIn LPVOID, cbDataIn uint32, dwFlags uint32) bool {
	ret1 := syscall3(cryptUnprotectMemory, 3,
		uintptr(unsafe.Pointer(pDataIn)),
		uintptr(cbDataIn),
		uintptr(dwFlags))
	return ret1 != 0
}

func CryptUnregisterDefaultOIDFunction(dwEncodingType uint32, pszFuncName /*const*/ LPCSTR, pwszDll string) bool {
	pwszDllStr := unicode16FromString(pwszDll)
	ret1 := syscall3(cryptUnregisterDefaultOIDFunction, 3,
		uintptr(dwEncodingType),
		uintptr(unsafe.Pointer(pszFuncName)),
		uintptr(unsafe.Pointer(&pwszDllStr[0])))
	return ret1 != 0
}

func CryptUnregisterOIDFunction(dwEncodingType uint32, pszFuncName /*const*/ LPCSTR, pszOID /*const*/ LPCSTR) bool {
	ret1 := syscall3(cryptUnregisterOIDFunction, 3,
		uintptr(dwEncodingType),
		uintptr(unsafe.Pointer(pszFuncName)),
		uintptr(unsafe.Pointer(pszOID)))
	return ret1 != 0
}

// TODO: Unknown type(s): PCCRYPT_OID_INFO
// func CryptUnregisterOIDInfo(pInfo PCCRYPT_OID_INFO) bool

// TODO: Unknown type(s): HCRYPTPROV_LEGACY, PCERT_PUBLIC_KEY_INFO
// func CryptVerifyCertificateSignature(hCryptProv HCRYPTPROV_LEGACY, dwCertEncodingType uint32, pbEncoded /*const*/ *byte, cbEncoded uint32, pPublicKey PCERT_PUBLIC_KEY_INFO) bool

// TODO: Unknown type(s): HCRYPTPROV_LEGACY
// func CryptVerifyCertificateSignatureEx(hCryptProv HCRYPTPROV_LEGACY, dwCertEncodingType uint32, dwSubjectType uint32, pvSubject uintptr, dwIssuerType uint32, pvIssuer uintptr, dwFlags uint32, pvExtra uintptr) bool

// TODO: Unknown type(s): PCRYPT_HASH_MESSAGE_PARA
// func CryptVerifyDetachedMessageHash(pHashPara PCRYPT_HASH_MESSAGE_PARA, pbDetachedHashBlob *byte, cbDetachedHashBlob uint32, cToBeHashed uint32, rgpbToBeHashed /*const*/ **BYTE, rgcbToBeHashed *uint32, pbComputedHash *byte, pcbComputedHash *uint32) bool

// TODO: Unknown type(s): PCCERT_CONTEXT *, PCRYPT_VERIFY_MESSAGE_PARA
// func CryptVerifyDetachedMessageSignature(pVerifyPara PCRYPT_VERIFY_MESSAGE_PARA, dwSignerIndex uint32, pbDetachedSignBlob /*const*/ *byte, cbDetachedSignBlob uint32, cToBeSigned uint32, rgpbToBeSigned /*const*/ **BYTE, rgcbToBeSigned *uint32, ppSignerCert PCCERT_CONTEXT *) bool

// TODO: Unknown type(s): PCRYPT_HASH_MESSAGE_PARA
// func CryptVerifyMessageHash(pHashPara PCRYPT_HASH_MESSAGE_PARA, pbHashedBlob *byte, cbHashedBlob uint32, pbToBeHashed *byte, pcbToBeHashed *uint32, pbComputedHash *byte, pcbComputedHash *uint32) bool

// TODO: Unknown type(s): PCCERT_CONTEXT *, PCRYPT_VERIFY_MESSAGE_PARA
// func CryptVerifyMessageSignature(pVerifyPara PCRYPT_VERIFY_MESSAGE_PARA, dwSignerIndex uint32, pbSignedBlob /*const*/ *byte, cbSignedBlob uint32, pbDecoded *byte, pcbDecoded *uint32, ppSignerCert PCCERT_CONTEXT *) bool

// TODO: Unknown type(s): PCERT_PUBLIC_KEY_INFO, PCRYPT_KEY_VERIFY_MESSAGE_PARA
// func CryptVerifyMessageSignatureWithKey(pVerifyPara PCRYPT_KEY_VERIFY_MESSAGE_PARA, pPublicKeyInfo PCERT_PUBLIC_KEY_INFO, pbSignedBlob /*const*/ *byte, cbSignedBlob uint32, pbDecoded *byte, pcbDecoded *uint32) bool

// TODO: Unknown type(s): ASN1decoding_t, HCRYPTASN1MODULE
// func I_CryptGetAsn1Decoder(hAsn1Module HCRYPTASN1MODULE) ASN1decoding_t

// TODO: Unknown type(s): ASN1encoding_t, HCRYPTASN1MODULE
// func I_CryptGetAsn1Encoder(hAsn1Module HCRYPTASN1MODULE) ASN1encoding_t

// TODO: Unknown type(s): ASN1module_t, HCRYPTASN1MODULE
// func I_CryptInstallAsn1Module(pMod ASN1module_t, dwFlags uint32, pvReserved uintptr) HCRYPTASN1MODULE

// TODO: Unknown type(s): HCRYPTASN1MODULE
// func I_CryptUninstallAsn1Module(hAsn1Module HCRYPTASN1MODULE) bool

func PFXExportCertStore(hStore HCERTSTORE, pPFX *CRYPT_DATA_BLOB, szPassword string, dwFlags uint32) bool {
	szPasswordStr := unicode16FromString(szPassword)
	ret1 := syscall6(pFXExportCertStore, 4,
		uintptr(hStore),
		uintptr(unsafe.Pointer(pPFX)),
		uintptr(unsafe.Pointer(&szPasswordStr[0])),
		uintptr(dwFlags),
		0,
		0)
	return ret1 != 0
}

func PFXExportCertStoreEx(hStore HCERTSTORE, pPFX *CRYPT_DATA_BLOB, szPassword string, pvPara uintptr, dwFlags uint32) bool {
	szPasswordStr := unicode16FromString(szPassword)
	ret1 := syscall6(pFXExportCertStoreEx, 5,
		uintptr(hStore),
		uintptr(unsafe.Pointer(pPFX)),
		uintptr(unsafe.Pointer(&szPasswordStr[0])),
		pvPara,
		uintptr(dwFlags),
		0)
	return ret1 != 0
}

func PFXImportCertStore(pPFX *CRYPT_DATA_BLOB, szPassword string, dwFlags uint32) HCERTSTORE {
	szPasswordStr := unicode16FromString(szPassword)
	ret1 := syscall3(pFXImportCertStore, 3,
		uintptr(unsafe.Pointer(pPFX)),
		uintptr(unsafe.Pointer(&szPasswordStr[0])),
		uintptr(dwFlags))
	return HCERTSTORE(ret1)
}

func PFXIsPFXBlob(pPFX *CRYPT_DATA_BLOB) bool {
	ret1 := syscall3(pFXIsPFXBlob, 1,
		uintptr(unsafe.Pointer(pPFX)),
		0,
		0)
	return ret1 != 0
}

func PFXVerifyPassword(pPFX *CRYPT_DATA_BLOB, szPassword string, dwFlags uint32) bool {
	szPasswordStr := unicode16FromString(szPassword)
	ret1 := syscall3(pFXVerifyPassword, 3,
		uintptr(unsafe.Pointer(pPFX)),
		uintptr(unsafe.Pointer(&szPasswordStr[0])),
		uintptr(dwFlags))
	return ret1 != 0
}

func I_CertUpdateStore(store1 HCERTSTORE, store2 HCERTSTORE, unk0 uint32, unk1 uint32) bool {
	ret1 := syscall6(i_CertUpdateStore, 4,
		uintptr(store1),
		uintptr(store2),
		uintptr(unk0),
		uintptr(unk1),
		0,
		0)
	return ret1 != 0
}

func I_CryptAllocTls() uint32 {
	ret1 := syscall3(i_CryptAllocTls, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): HLRUCACHE *
// func I_CryptCreateLruCache(unknown uintptr, out HLRUCACHE *) bool

// TODO: Unknown type(s): HLRUCACHE
// func I_CryptCreateLruEntry(h HLRUCACHE, unk0 uint32, unk1 uint32) bool

func I_CryptDetachTls(dwTlsIndex uint32) LPVOID {
	ret1 := syscall3(i_CryptDetachTls, 1,
		uintptr(dwTlsIndex),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func I_CryptFindLruEntry(unk0 uint32, unk1 uint32) bool {
	ret1 := syscall3(i_CryptFindLruEntry, 2,
		uintptr(unk0),
		uintptr(unk1),
		0)
	return ret1 != 0
}

func I_CryptFindLruEntryData(unk0 uint32, unk1 uint32, unk2 uint32) bool {
	ret1 := syscall3(i_CryptFindLruEntryData, 3,
		uintptr(unk0),
		uintptr(unk1),
		uintptr(unk2))
	return ret1 != 0
}

// TODO: Unknown type(s): HLRUCACHE
// func I_CryptFlushLruCache(h HLRUCACHE, unk0 uint32, unk1 uint32) uint32

// TODO: Unknown type(s): HLRUCACHE
// func I_CryptFreeLruCache(h HLRUCACHE, unk0 uint32, unk1 uint32) HLRUCACHE

func I_CryptFreeTls(dwTlsIndex uint32, unknown uint32) bool {
	ret1 := syscall3(i_CryptFreeTls, 2,
		uintptr(dwTlsIndex),
		uintptr(unknown),
		0)
	return ret1 != 0
}

func I_CryptGetDefaultCryptProv(reserved uint32) HCRYPTPROV {
	ret1 := syscall3(i_CryptGetDefaultCryptProv, 1,
		uintptr(reserved),
		0,
		0)
	return HCRYPTPROV(ret1)
}

func I_CryptGetOssGlobal(x uint32) bool {
	ret1 := syscall3(i_CryptGetOssGlobal, 1,
		uintptr(x),
		0,
		0)
	return ret1 != 0
}

func I_CryptGetTls(dwTlsIndex uint32) LPVOID {
	ret1 := syscall3(i_CryptGetTls, 1,
		uintptr(dwTlsIndex),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func I_CryptInstallOssGlobal(x uint32, y uint32, z uint32) uint32 {
	ret1 := syscall3(i_CryptInstallOssGlobal, 3,
		uintptr(x),
		uintptr(y),
		uintptr(z))
	return uint32(ret1)
}

func I_CryptReadTrustedPublisherDWORDValueFromRegistry(name string, value *uint32) bool {
	nameStr := unicode16FromString(name)
	ret1 := syscall3(i_CryptReadTrustedPublisherDWORDValueFromRegistry, 2,
		uintptr(unsafe.Pointer(&nameStr[0])),
		uintptr(unsafe.Pointer(value)),
		0)
	return ret1 != 0
}

func I_CryptSetTls(dwTlsIndex uint32, lpTlsValue LPVOID) bool {
	ret1 := syscall3(i_CryptSetTls, 2,
		uintptr(dwTlsIndex),
		uintptr(unsafe.Pointer(lpTlsValue)),
		0)
	return ret1 != 0
}
