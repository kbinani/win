package win

//ref GUID
//ref ULONG
//ref BOOLEAN

type TRACE_GUID_PROPERTIES struct {
	Guid        GUID
	GuidType    ULONG
	LoggerId    ULONG
	EnableLevel ULONG
	EnableFlags ULONG
	IsEnable    BOOLEAN
}
