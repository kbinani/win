package win

//ref DWORD
//ref LPWSTR

type QUERY_SERVICE_LOCK_STATUS struct {
	FIsLocked      DWORD
	LpLockOwner    LPWSTR
	DwLockDuration DWORD
}
