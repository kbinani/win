package win

//ref NET_LUID
//ref NET_IFINDEX
//ref GUID
//ref WCHAR
//ref ULONG
//ref UCHAR
//ref IFTYPE
//ref TUNNEL_TYPE
//ref NDIS_MEDIUM
//ref NDIS_PHYSICAL_MEDIUM
//ref NET_IF_ACCESS_TYPE
//ref NET_IF_DIRECTION_TYPE
//ref BOOLEAN
//ref IF_OPER_STATUS
//ref NET_IF_ADMIN_STATUS
//ref NET_IF_MEDIA_CONNECT_STATE
//ref NET_IF_NETWORK_GUID
//ref NET_IF_CONNECTION_TYPE
//ref ULONG64

type MIB_IF_ROW2 struct {
	InterfaceLuid               NET_LUID
	InterfaceIndex              NET_IFINDEX
	InterfaceGuid               GUID
	Alias                       [IF_MAX_STRING_SIZE + 1]WCHAR
	Description                 [IF_MAX_STRING_SIZE + 1]WCHAR
	PhysicalAddressLength       ULONG
	PhysicalAddress             [IF_MAX_PHYS_ADDRESS_LENGTH]UCHAR
	PermanentPhysicalAddress    [IF_MAX_PHYS_ADDRESS_LENGTH]UCHAR
	Mtu                         ULONG
	Type                        IFTYPE
	TunnelType                  TUNNEL_TYPE
	MediaType                   NDIS_MEDIUM
	PhysicalMediumType          NDIS_PHYSICAL_MEDIUM
	AccessType                  NET_IF_ACCESS_TYPE
	DirectionType               NET_IF_DIRECTION_TYPE
	InterfaceAndOperStatusFlags BOOLEAN
	OperStatus                  IF_OPER_STATUS
	AdminStatus                 NET_IF_ADMIN_STATUS
	MediaConnectState           NET_IF_MEDIA_CONNECT_STATE
	NetworkGuid                 NET_IF_NETWORK_GUID
	ConnectionType              NET_IF_CONNECTION_TYPE
	padding1                    [pad0for64_4for32]byte
	TransmitLinkSpeed           ULONG64
	ReceiveLinkSpeed            ULONG64
	InOctets                    ULONG64
	InUcastPkts                 ULONG64
	InNUcastPkts                ULONG64
	InDiscards                  ULONG64
	InErrors                    ULONG64
	InUnknownProtos             ULONG64
	InUcastOctets               ULONG64
	InMulticastOctets           ULONG64
	InBroadcastOctets           ULONG64
	OutOctets                   ULONG64
	OutUcastPkts                ULONG64
	OutNUcastPkts               ULONG64
	OutDiscards                 ULONG64
	OutErrors                   ULONG64
	OutUcastOctets              ULONG64
	OutMulticastOctets          ULONG64
	OutBroadcastOctets          ULONG64
	OutQLen                     ULONG64
}

func (this *MIB_IF_ROW2) HardwareInterface() BOOLEAN {
	return BOOLEAN(this.InterfaceAndOperStatusFlags & 0x1)
}

func (this *MIB_IF_ROW2) FilterInterface() BOOLEAN {
	return BOOLEAN((this.InterfaceAndOperStatusFlags >> 1) & 0x1)
}

func (this *MIB_IF_ROW2) ConnectorPresent() BOOLEAN {
	return BOOLEAN((this.InterfaceAndOperStatusFlags >> 2) & 0x1)
}

func (this *MIB_IF_ROW2) NotAuthenticated() BOOLEAN {
	return BOOLEAN((this.InterfaceAndOperStatusFlags >> 3) & 0x1)
}

func (this *MIB_IF_ROW2) NotMediaConnected() BOOLEAN {
	return BOOLEAN((this.InterfaceAndOperStatusFlags >> 4) & 0x1)
}

func (this *MIB_IF_ROW2) Paused() BOOLEAN {
	return BOOLEAN((this.InterfaceAndOperStatusFlags >> 5) & 0x1)
}

func (this *MIB_IF_ROW2) LowPower() BOOLEAN {
	return BOOLEAN((this.InterfaceAndOperStatusFlags >> 6) & 0x1)
}

func (this *MIB_IF_ROW2) EndPointInterface() BOOLEAN {
	return BOOLEAN((this.InterfaceAndOperStatusFlags >> 7) & 0x1)
}
