package win

//ref LPWSTR
//ref SERVICE_STATUS

type ENUM_SERVICE_STATUS struct {
	LpServiceName LPWSTR
	LpDisplayName LPWSTR
	ServiceStatus SERVICE_STATUS
}
