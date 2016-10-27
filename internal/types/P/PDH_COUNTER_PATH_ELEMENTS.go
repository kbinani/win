package win

type PDH_COUNTER_PATH_ELEMENTS struct {
	SzMachineName    LPWSTR
	SzObjectName     LPWSTR
	SzInstanceName   LPWSTR
	SzParentInstance LPWSTR
	DwInstanceIndex  DWORD
	SzCounterName    LPWSTR
}
