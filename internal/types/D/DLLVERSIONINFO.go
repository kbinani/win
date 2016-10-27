package win

//ref DWORD
type DLLVERSIONINFO struct {
	CbSize         DWORD
	DwMajorVersion DWORD
	DwMinorVersion DWORD
	DwBuildNumber  DWORD
	DwPlatformID   DWORD
}
