package win

//ref ULONG
//ref HANDLE
//ref PLSA_UNICODE_STRING
//ref PVOID

type LSA_OBJECT_ATTRIBUTES struct {
	Length                   ULONG
	RootDirectory            HANDLE
	ObjectName               PLSA_UNICODE_STRING
	Attributes               ULONG
	SecurityDescriptor       PVOID
	SecurityQualityOfService PVOID
}
