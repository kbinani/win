package win

//ref WORD
//ref WCHAR
//ref UINT

type JOYCAPS struct {
	WMid        WORD
	WPid        WORD
	SzPname     [MAXPNAMELEN]WCHAR
	WXmin       UINT
	WXmax       UINT
	WYmin       UINT
	WYmax       UINT
	WZmin       UINT
	WZmax       UINT
	WNumButtons UINT
	WPeriodMin  UINT
	WPeriodMax  UINT
	WRmin       UINT
	WRmax       UINT
	WUmin       UINT
	WUmax       UINT
	WVmin       UINT
	WVmax       UINT
	WCaps       UINT
	WMaxAxes    UINT
	WNumAxes    UINT
	WMaxButtons UINT
	SzRegKey    [MAXPNAMELEN]WCHAR
	SzOEMVxD    [MAX_JOYSTICKOEMVXDNAME]WCHAR
}
