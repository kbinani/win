package win

//ref DWORD
//ref GUID
//ref HANDLE
//ref LPCWSTR
//ref HCRYPTPROV
//ref CRYPT_ALGORITHM_IDENTIFIER
//ref LPVOID
//ref MS_ADDINFO_FLAT
//ref MS_ADDINFO_CATALOGMEMBER
//ref MS_ADDINFO_BLOB

type SIP_SUBJECTINFO struct {
	CbSize              DWORD
	PgSubjectType       *GUID
	HFile               HANDLE
	PwsFileName         LPCWSTR
	PwsDisplayName      LPCWSTR
	DwReserved1         DWORD
	DwIntVersion        DWORD
	HProv               HCRYPTPROV
	DigestAlgorithm     CRYPT_ALGORITHM_IDENTIFIER
	DwFlags             DWORD
	DwEncodingType      DWORD
	DwReserved2         DWORD
	FdwCAPISettings     DWORD
	FdwSecuritySettings DWORD
	DwIndex             DWORD
	DwUnionChoice       DWORD
	union1              uintptr
	PClientData         LPVOID
}

func (this *SIP_SUBJECTINFO) GetPsFlat() *MS_ADDINFO_FLAT {
	return (*MS_ADDINFO_FLAT)(unsafe.Pointer(this.union1))
}

func (this *SIP_SUBJECTINFO) SetPsFlat(v *MS_ADDINFO_FLAT) {
	this.union1 = uintptr(unsafe.Pointer(v))
}

func (this *SIP_SUBJECTINFO) GetPsCatMember() *MS_ADDINFO_CATALOGMEMBER {
	return (*MS_ADDINFO_CATALOGMEMBER)(unsafe.Pointer(this.union1))
}

func (this *SIP_SUBJECTINFO) SetPsCatMember(v *MS_ADDINFO_CATALOGMEMBER) {
	this.union1 = uintptr(unsafe.Pointer(v))
}

func (this *SIP_SUBJECTINFO) GetPsBlob() *MS_ADDINFO_BLOB {
	return (*MS_ADDINFO_BLOB)(unsafe.Pointer(this.union1))
}

func (this *SIP_SUBJECTINFO) SetPsBlob(v *MS_ADDINFO_BLOB) {
	this.union1 = uintptr(unsafe.Pointer(v))
}
