package win

//ref DWORD

type JOYINFOEX struct {
	DwSize         DWORD
	DwFlags        DWORD
	DwXpos         DWORD
	DwYpos         DWORD
	DwZpos         DWORD
	DwRpos         DWORD
	DwUpos         DWORD
	DwVpos         DWORD
	DwButtons      DWORD
	DwButtonNumber DWORD
	DwPOV          DWORD
	DwReserved1    DWORD
	DwReserved2    DWORD
}
