package win

//ref DWORD

type SERVICE_STATUS_PROCESS struct {
	DwServiceType             DWORD
	DwCurrentState            DWORD
	DwControlsAccepted        DWORD
	DwWin32ExitCode           DWORD
	DwServiceSpecificExitCode DWORD
	DwCheckPoint              DWORD
	DwWaitHint                DWORD
	DwProcessId               DWORD
	DwServiceFlags            DWORD
}
