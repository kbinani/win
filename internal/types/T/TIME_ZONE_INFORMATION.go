package win

//ref LONG
//ref WCHAR
//ref SYSTEMTIME

type TIME_ZONE_INFORMATION struct {
	Bias         LONG
	StandardName [32]WCHAR
	StandardDate SYSTEMTIME
	StandardBias LONG
	DaylightName [32]WCHAR
	DaylightDate SYSTEMTIME
	DaylightBias LONG
}
