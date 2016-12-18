package win

//ref DWORD
//ref LPCWSTR
//ref HANDLE
//ref BYTE
//ref LARGE_INTEGER
//ref ALG_ID
//ref LPBYTE
//ref HWND
//ref ULONG64
//ref BOOL

type SAFER_CODE_PROPERTIES struct {
	CbSize             DWORD
	DwCheckFlags       DWORD
	ImagePath          LPCWSTR
	HImageFileHandle   HANDLE
	UrlZoneId          DWORD
	ImageHash          [SAFER_MAX_HASH_SIZE]BYTE
	DwImageHashSize    DWORD
	ImageSize          LARGE_INTEGER
	HashAlgorithm      ALG_ID
	PByteBlock         LPBYTE
	HWndParent         HWND
	DwWVTUIChoice      DWORD
	PackageMoniker     LPCWSTR
	PackagePublisher   LPCWSTR
	PackageName        LPCWSTR
	padding0           [pad0for64_4for32]byte
	PackageVersion     ULONG64
	PackageIsFramework BOOL
	padding1           [pad0for64_4for32]byte
}
