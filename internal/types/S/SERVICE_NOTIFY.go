package win

//ref DWORD
//ref PFN_SC_NOTIFY_CALLBACK
//ref PVOID
//ref SERVICE_STATUS_PROCESS
//ref LPWSTR

type SERVICE_NOTIFY struct {
	DwVersion               DWORD
	PfnNotifyCallback       uintptr // PFN_SC_NOTIFY_CALLBACK
	PContext                PVOID
	DwNotificationStatus    DWORD
	ServiceStatus           SERVICE_STATUS_PROCESS
	DwNotificationTriggered DWORD
	PszServiceNames         LPWSTR
}
