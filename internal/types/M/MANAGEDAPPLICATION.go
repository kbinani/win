package win

//ref LPWSTR
//ref DWORD
//ref GUID
//ref LANGID
//ref BOOL

type MANAGEDAPPLICATION struct {
	PszPackageName LPWSTR
	PszPublisher   LPWSTR
	DwVersionHi    DWORD
	DwVersionLo    DWORD
	DwRevision     DWORD
	GpoId          GUID
	PszPolicyName  LPWSTR
	ProductId      GUID
	Language       LANGID
	PszOwner       LPWSTR
	PszCompany     LPWSTR
	PszComments    LPWSTR
	PszContact     LPWSTR
	PszSupportUrl  LPWSTR
	DwPathType     DWORD
	BInstalled     BOOL
}
