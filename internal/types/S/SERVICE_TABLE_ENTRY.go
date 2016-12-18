package win

//ref LPWSTR
//ref LPSERVICE_MAIN_FUNCTION

type SERVICE_TABLE_ENTRY struct {
	LpServiceName LPWSTR
	LpServiceProc uintptr // LPSERVICE_MAIN_FUNCTION
}
