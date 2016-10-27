package win

type SERVICE_STATUS struct {
	DwServiceType             DWORD
	DwCurrentState            DWORD
	DwControlsAccepted        DWORD
	DwWin32ExitCode           DWORD
	DwServiceSpecificExitCode DWORD
	DwCheckPoint              DWORD
	DwWaitHint                DWORD
}
