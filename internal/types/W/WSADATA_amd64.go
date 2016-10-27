package win

type WSADATA struct {
	WVersion       uint16
	WHighVersion   uint16
	IMaxSockets    uint16
	IMaxUdpDg      uint16
	LpVendorInfo   *byte
	SzDescription  [WSADESCRIPTION_LEN + 1]byte
	SzSystemStatus [WSASYS_STATUS_LEN + 1]byte
}
