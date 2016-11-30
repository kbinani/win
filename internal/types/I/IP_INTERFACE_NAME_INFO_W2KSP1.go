package win

//ref ULONG
//ref UCHAR
//ref GUID

type IP_INTERFACE_NAME_INFO_W2KSP1 struct {
	Index          ULONG
	MediaType      ULONG
	ConnectionType UCHAR
	AccessType     UCHAR
	DeviceGuid     GUID
	InterfaceGuid  GUID
}
