package win

//ref DWORD
//ref LPWSTR

type QUERY_SERVICE_CONFIG struct {
	DwServiceType      DWORD
	DwStartType        DWORD
	DwErrorControl     DWORD
	LpBinaryPathName   LPWSTR
	LpLoadOrderGroup   LPWSTR
	DwTagId            DWORD
	LpDependencies     LPWSTR
	LpServiceStartName LPWSTR
	LpDisplayName      LPWSTR
}
