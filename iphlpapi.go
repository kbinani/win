// +build windows

package win

import (
	"unsafe"
)

var (
	// Library
	libiphlpapi uintptr

	// Functions
	addIPAddress                          uintptr
	cancelIPChangeNotify                  uintptr
	createIpForwardEntry                  uintptr
	createIpNetEntry                      uintptr
	createPersistentTcpPortReservation    uintptr
	createPersistentUdpPortReservation    uintptr
	createProxyArpEntry                   uintptr
	deleteIPAddress                       uintptr
	deleteIpForwardEntry                  uintptr
	deleteIpNetEntry                      uintptr
	deletePersistentTcpPortReservation    uintptr
	deletePersistentUdpPortReservation    uintptr
	deleteProxyArpEntry                   uintptr
	disableMediaSense                     uintptr
	enableRouter                          uintptr
	flushIpNetTable                       uintptr
	getAdapterIndex                       uintptr
	getAdapterOrderMap                    uintptr
	getAdaptersAddresses                  uintptr
	getAdaptersInfo                       uintptr
	getBestInterface                      uintptr
	getBestInterfaceEx                    uintptr
	getBestRoute                          uintptr
	getExtendedTcpTable                   uintptr
	getExtendedUdpTable                   uintptr
	getFriendlyIfIndex                    uintptr
	getIcmpStatistics                     uintptr
	getIcmpStatisticsEx                   uintptr
	getIfEntry                            uintptr
	getIfTable                            uintptr
	getInterfaceInfo                      uintptr
	getIpAddrTable                        uintptr
	getIpErrorString                      uintptr
	getIpForwardTable                     uintptr
	getIpNetTable                         uintptr
	getIpStatistics                       uintptr
	getIpStatisticsEx                     uintptr
	getNetworkParams                      uintptr
	getNumberOfInterfaces                 uintptr
	getOwnerModuleFromTcp6Entry           uintptr
	getOwnerModuleFromTcpEntry            uintptr
	getOwnerModuleFromUdp6Entry           uintptr
	getOwnerModuleFromUdpEntry            uintptr
	getPerAdapterInfo                     uintptr
	getRTTAndHopCount                     uintptr
	icmp6CreateFile                       uintptr
	icmp6ParseReplies                     uintptr
	icmpCloseHandle                       uintptr
	icmpCreateFile                        uintptr
	icmpParseReplies                      uintptr
	lookupPersistentTcpPortReservation    uintptr
	lookupPersistentUdpPortReservation    uintptr
	notifyAddrChange                      uintptr
	notifyRouteChange                     uintptr
	resolveNeighbor                       uintptr
	restoreMediaSense                     uintptr
	sendARP                               uintptr
	setIfEntry                            uintptr
	setIpForwardEntry                     uintptr
	setIpNetEntry                         uintptr
	setIpStatistics                       uintptr
	setIpStatisticsEx                     uintptr
	setIpTTL                              uintptr
	setTcpEntry                           uintptr
	unenableRouter                        uintptr
	allocateAndGetIfTableFromStack        uintptr
	allocateAndGetIpAddrTableFromStack    uintptr
	allocateAndGetIpForwardTableFromStack uintptr
	allocateAndGetIpNetTableFromStack     uintptr
	cancelMibChangeNotify2                uintptr
	convertInterfaceGuidToLuid            uintptr
	convertInterfaceIndexToLuid           uintptr
	convertInterfaceLuidToGuid            uintptr
	convertInterfaceLuidToIndex           uintptr
	convertInterfaceLuidToName            uintptr
	convertInterfaceNameToLuid            uintptr
	freeMibTable                          uintptr
	notifyIpInterfaceChange               uintptr
	pfDeleteInterface                     uintptr
	pfUnBindInterface                     uintptr
)

func init() {
	// Library
	libiphlpapi = doLoadLibrary("iphlpapi.dll")

	// Functions
	addIPAddress = doGetProcAddress(libiphlpapi, "AddIPAddress")
	cancelIPChangeNotify = doGetProcAddress(libiphlpapi, "CancelIPChangeNotify")
	createIpForwardEntry = doGetProcAddress(libiphlpapi, "CreateIpForwardEntry")
	createIpNetEntry = doGetProcAddress(libiphlpapi, "CreateIpNetEntry")
	createPersistentTcpPortReservation = doGetProcAddress(libiphlpapi, "CreatePersistentTcpPortReservation")
	createPersistentUdpPortReservation = doGetProcAddress(libiphlpapi, "CreatePersistentUdpPortReservation")
	createProxyArpEntry = doGetProcAddress(libiphlpapi, "CreateProxyArpEntry")
	deleteIPAddress = doGetProcAddress(libiphlpapi, "DeleteIPAddress")
	deleteIpForwardEntry = doGetProcAddress(libiphlpapi, "DeleteIpForwardEntry")
	deleteIpNetEntry = doGetProcAddress(libiphlpapi, "DeleteIpNetEntry")
	deletePersistentTcpPortReservation = doGetProcAddress(libiphlpapi, "DeletePersistentTcpPortReservation")
	deletePersistentUdpPortReservation = doGetProcAddress(libiphlpapi, "DeletePersistentUdpPortReservation")
	deleteProxyArpEntry = doGetProcAddress(libiphlpapi, "DeleteProxyArpEntry")
	disableMediaSense = doGetProcAddress(libiphlpapi, "DisableMediaSense")
	enableRouter = doGetProcAddress(libiphlpapi, "EnableRouter")
	flushIpNetTable = doGetProcAddress(libiphlpapi, "FlushIpNetTable")
	getAdapterIndex = doGetProcAddress(libiphlpapi, "GetAdapterIndex")
	getAdapterOrderMap = doGetProcAddress(libiphlpapi, "GetAdapterOrderMap")
	getAdaptersAddresses = doGetProcAddress(libiphlpapi, "GetAdaptersAddresses")
	getAdaptersInfo = doGetProcAddress(libiphlpapi, "GetAdaptersInfo")
	getBestInterface = doGetProcAddress(libiphlpapi, "GetBestInterface")
	getBestInterfaceEx = doGetProcAddress(libiphlpapi, "GetBestInterfaceEx")
	getBestRoute = doGetProcAddress(libiphlpapi, "GetBestRoute")
	getExtendedTcpTable = doGetProcAddress(libiphlpapi, "GetExtendedTcpTable")
	getExtendedUdpTable = doGetProcAddress(libiphlpapi, "GetExtendedUdpTable")
	getFriendlyIfIndex = doGetProcAddress(libiphlpapi, "GetFriendlyIfIndex")
	getIcmpStatistics = doGetProcAddress(libiphlpapi, "GetIcmpStatistics")
	getIcmpStatisticsEx = doGetProcAddress(libiphlpapi, "GetIcmpStatisticsEx")
	getIfEntry = doGetProcAddress(libiphlpapi, "GetIfEntry")
	getIfTable = doGetProcAddress(libiphlpapi, "GetIfTable")
	getInterfaceInfo = doGetProcAddress(libiphlpapi, "GetInterfaceInfo")
	getIpAddrTable = doGetProcAddress(libiphlpapi, "GetIpAddrTable")
	getIpErrorString = doGetProcAddress(libiphlpapi, "GetIpErrorString")
	getIpForwardTable = doGetProcAddress(libiphlpapi, "GetIpForwardTable")
	getIpNetTable = doGetProcAddress(libiphlpapi, "GetIpNetTable")
	getIpStatistics = doGetProcAddress(libiphlpapi, "GetIpStatistics")
	getIpStatisticsEx = doGetProcAddress(libiphlpapi, "GetIpStatisticsEx")
	getNetworkParams = doGetProcAddress(libiphlpapi, "GetNetworkParams")
	getNumberOfInterfaces = doGetProcAddress(libiphlpapi, "GetNumberOfInterfaces")
	getOwnerModuleFromTcp6Entry = doGetProcAddress(libiphlpapi, "GetOwnerModuleFromTcp6Entry")
	getOwnerModuleFromTcpEntry = doGetProcAddress(libiphlpapi, "GetOwnerModuleFromTcpEntry")
	getOwnerModuleFromUdp6Entry = doGetProcAddress(libiphlpapi, "GetOwnerModuleFromUdp6Entry")
	getOwnerModuleFromUdpEntry = doGetProcAddress(libiphlpapi, "GetOwnerModuleFromUdpEntry")
	getPerAdapterInfo = doGetProcAddress(libiphlpapi, "GetPerAdapterInfo")
	getRTTAndHopCount = doGetProcAddress(libiphlpapi, "GetRTTAndHopCount")
	icmp6CreateFile = doGetProcAddress(libiphlpapi, "Icmp6CreateFile")
	icmp6ParseReplies = doGetProcAddress(libiphlpapi, "Icmp6ParseReplies")
	icmpCloseHandle = doGetProcAddress(libiphlpapi, "IcmpCloseHandle")
	icmpCreateFile = doGetProcAddress(libiphlpapi, "IcmpCreateFile")
	icmpParseReplies = doGetProcAddress(libiphlpapi, "IcmpParseReplies")
	lookupPersistentTcpPortReservation = doGetProcAddress(libiphlpapi, "LookupPersistentTcpPortReservation")
	lookupPersistentUdpPortReservation = doGetProcAddress(libiphlpapi, "LookupPersistentUdpPortReservation")
	notifyAddrChange = doGetProcAddress(libiphlpapi, "NotifyAddrChange")
	notifyRouteChange = doGetProcAddress(libiphlpapi, "NotifyRouteChange")
	resolveNeighbor = doGetProcAddress(libiphlpapi, "ResolveNeighbor")
	restoreMediaSense = doGetProcAddress(libiphlpapi, "RestoreMediaSense")
	sendARP = doGetProcAddress(libiphlpapi, "SendARP")
	setIfEntry = doGetProcAddress(libiphlpapi, "SetIfEntry")
	setIpForwardEntry = doGetProcAddress(libiphlpapi, "SetIpForwardEntry")
	setIpNetEntry = doGetProcAddress(libiphlpapi, "SetIpNetEntry")
	setIpStatistics = doGetProcAddress(libiphlpapi, "SetIpStatistics")
	setIpStatisticsEx = doGetProcAddress(libiphlpapi, "SetIpStatisticsEx")
	setIpTTL = doGetProcAddress(libiphlpapi, "SetIpTTL")
	setTcpEntry = doGetProcAddress(libiphlpapi, "SetTcpEntry")
	unenableRouter = doGetProcAddress(libiphlpapi, "UnenableRouter")
	allocateAndGetIfTableFromStack = doGetProcAddress(libiphlpapi, "AllocateAndGetIfTableFromStack")
	allocateAndGetIpAddrTableFromStack = doGetProcAddress(libiphlpapi, "AllocateAndGetIpAddrTableFromStack")
	allocateAndGetIpForwardTableFromStack = doGetProcAddress(libiphlpapi, "AllocateAndGetIpForwardTableFromStack")
	allocateAndGetIpNetTableFromStack = doGetProcAddress(libiphlpapi, "AllocateAndGetIpNetTableFromStack")
	cancelMibChangeNotify2 = doGetProcAddress(libiphlpapi, "CancelMibChangeNotify2")
	convertInterfaceGuidToLuid = doGetProcAddress(libiphlpapi, "ConvertInterfaceGuidToLuid")
	convertInterfaceIndexToLuid = doGetProcAddress(libiphlpapi, "ConvertInterfaceIndexToLuid")
	convertInterfaceLuidToGuid = doGetProcAddress(libiphlpapi, "ConvertInterfaceLuidToGuid")
	convertInterfaceLuidToIndex = doGetProcAddress(libiphlpapi, "ConvertInterfaceLuidToIndex")
	convertInterfaceLuidToName = doGetProcAddress(libiphlpapi, "ConvertInterfaceLuidToNameW")
	convertInterfaceNameToLuid = doGetProcAddress(libiphlpapi, "ConvertInterfaceNameToLuidW")
	freeMibTable = doGetProcAddress(libiphlpapi, "FreeMibTable")
	notifyIpInterfaceChange = doGetProcAddress(libiphlpapi, "NotifyIpInterfaceChange")
	pfDeleteInterface = doGetProcAddress(libiphlpapi, "PfDeleteInterface")
	pfUnBindInterface = doGetProcAddress(libiphlpapi, "PfUnBindInterface")
}

func AddIPAddress(address IPAddr, ipMask IPMask, ifIndex uint32, nTEContext *uint32, nTEInstance *uint32) uint32 {
	ret1 := syscall6(addIPAddress, 5,
		uintptr(address),
		uintptr(ipMask),
		uintptr(ifIndex),
		uintptr(unsafe.Pointer(nTEContext)),
		uintptr(unsafe.Pointer(nTEInstance)),
		0)
	return uint32(ret1)
}

func CancelIPChangeNotify(notifyOverlapped *OVERLAPPED) bool {
	ret1 := syscall3(cancelIPChangeNotify, 1,
		uintptr(unsafe.Pointer(notifyOverlapped)),
		0,
		0)
	return ret1 != 0
}

func CreateIpForwardEntry(pRoute PMIB_IPFORWARDROW) uint32 {
	ret1 := syscall3(createIpForwardEntry, 1,
		uintptr(unsafe.Pointer(pRoute)),
		0,
		0)
	return uint32(ret1)
}

func CreateIpNetEntry(pArpEntry PMIB_IPNETROW) uint32 {
	ret1 := syscall3(createIpNetEntry, 1,
		uintptr(unsafe.Pointer(pArpEntry)),
		0,
		0)
	return uint32(ret1)
}

func CreatePersistentTcpPortReservation(startPort USHORT, numberOfPorts USHORT, token PULONG64) ULONG {
	ret1 := syscall3(createPersistentTcpPortReservation, 3,
		uintptr(startPort),
		uintptr(numberOfPorts),
		uintptr(unsafe.Pointer(token)))
	return ULONG(ret1)
}

func CreatePersistentUdpPortReservation(startPort USHORT, numberOfPorts USHORT, token PULONG64) ULONG {
	ret1 := syscall3(createPersistentUdpPortReservation, 3,
		uintptr(startPort),
		uintptr(numberOfPorts),
		uintptr(unsafe.Pointer(token)))
	return ULONG(ret1)
}

func CreateProxyArpEntry(dwAddress uint32, dwMask uint32, dwIfIndex uint32) uint32 {
	ret1 := syscall3(createProxyArpEntry, 3,
		uintptr(dwAddress),
		uintptr(dwMask),
		uintptr(dwIfIndex))
	return uint32(ret1)
}

func DeleteIPAddress(nTEContext ULONG) uint32 {
	ret1 := syscall3(deleteIPAddress, 1,
		uintptr(nTEContext),
		0,
		0)
	return uint32(ret1)
}

func DeleteIpForwardEntry(pRoute PMIB_IPFORWARDROW) uint32 {
	ret1 := syscall3(deleteIpForwardEntry, 1,
		uintptr(unsafe.Pointer(pRoute)),
		0,
		0)
	return uint32(ret1)
}

func DeleteIpNetEntry(pArpEntry PMIB_IPNETROW) uint32 {
	ret1 := syscall3(deleteIpNetEntry, 1,
		uintptr(unsafe.Pointer(pArpEntry)),
		0,
		0)
	return uint32(ret1)
}

func DeletePersistentTcpPortReservation(startPort USHORT, numberOfPorts USHORT) ULONG {
	ret1 := syscall3(deletePersistentTcpPortReservation, 2,
		uintptr(startPort),
		uintptr(numberOfPorts),
		0)
	return ULONG(ret1)
}

func DeletePersistentUdpPortReservation(startPort USHORT, numberOfPorts USHORT) ULONG {
	ret1 := syscall3(deletePersistentUdpPortReservation, 2,
		uintptr(startPort),
		uintptr(numberOfPorts),
		0)
	return ULONG(ret1)
}

func DeleteProxyArpEntry(dwAddress uint32, dwMask uint32, dwIfIndex uint32) uint32 {
	ret1 := syscall3(deleteProxyArpEntry, 3,
		uintptr(dwAddress),
		uintptr(dwMask),
		uintptr(dwIfIndex))
	return uint32(ret1)
}

func DisableMediaSense(pHandle *HANDLE, pOverLapped *OVERLAPPED) uint32 {
	ret1 := syscall3(disableMediaSense, 2,
		uintptr(unsafe.Pointer(pHandle)),
		uintptr(unsafe.Pointer(pOverLapped)),
		0)
	return uint32(ret1)
}

func EnableRouter(pHandle *HANDLE, pOverlapped *OVERLAPPED) uint32 {
	ret1 := syscall3(enableRouter, 2,
		uintptr(unsafe.Pointer(pHandle)),
		uintptr(unsafe.Pointer(pOverlapped)),
		0)
	return uint32(ret1)
}

func FlushIpNetTable(dwIfIndex uint32) uint32 {
	ret1 := syscall3(flushIpNetTable, 1,
		uintptr(dwIfIndex),
		0,
		0)
	return uint32(ret1)
}

func GetAdapterIndex(adapterName LPWSTR, ifIndex *uint32) uint32 {
	ret1 := syscall3(getAdapterIndex, 2,
		uintptr(unsafe.Pointer(adapterName)),
		uintptr(unsafe.Pointer(ifIndex)),
		0)
	return uint32(ret1)
}

func GetAdapterOrderMap() PIP_ADAPTER_ORDER_MAP {
	ret1 := syscall3(getAdapterOrderMap, 0,
		0,
		0,
		0)
	return (PIP_ADAPTER_ORDER_MAP)(unsafe.Pointer(ret1))
}

func GetAdaptersAddresses(family ULONG, flags ULONG, reserved uintptr, adapterAddresses PIP_ADAPTER_ADDRESSES, sizePointer *uint32) ULONG {
	ret1 := syscall6(getAdaptersAddresses, 5,
		uintptr(family),
		uintptr(flags),
		reserved,
		uintptr(unsafe.Pointer(adapterAddresses)),
		uintptr(unsafe.Pointer(sizePointer)),
		0)
	return ULONG(ret1)
}

func GetAdaptersInfo(adapterInfo PIP_ADAPTER_INFO, sizePointer *uint32) ULONG {
	ret1 := syscall3(getAdaptersInfo, 2,
		uintptr(unsafe.Pointer(adapterInfo)),
		uintptr(unsafe.Pointer(sizePointer)),
		0)
	return ULONG(ret1)
}

func GetBestInterface(dwDestAddr IPAddr, pdwBestIfIndex *DWORD) uint32 {
	ret1 := syscall3(getBestInterface, 2,
		uintptr(dwDestAddr),
		uintptr(unsafe.Pointer(pdwBestIfIndex)),
		0)
	return uint32(ret1)
}

func GetBestInterfaceEx(pDestAddr *Sockaddr, pdwBestIfIndex *DWORD) uint32 {
	ret1 := syscall3(getBestInterfaceEx, 2,
		uintptr(unsafe.Pointer(pDestAddr)),
		uintptr(unsafe.Pointer(pdwBestIfIndex)),
		0)
	return uint32(ret1)
}

func GetBestRoute(dwDestAddr uint32, dwSourceAddr uint32, pBestRoute PMIB_IPFORWARDROW) uint32 {
	ret1 := syscall3(getBestRoute, 3,
		uintptr(dwDestAddr),
		uintptr(dwSourceAddr),
		uintptr(unsafe.Pointer(pBestRoute)))
	return uint32(ret1)
}

func GetExtendedTcpTable(pTcpTable uintptr, pdwSize *DWORD, bOrder bool, ulAf ULONG, tableClass TCP_TABLE_CLASS, reserved ULONG) uint32 {
	ret1 := syscall6(getExtendedTcpTable, 6,
		pTcpTable,
		uintptr(unsafe.Pointer(pdwSize)),
		getUintptrFromBool(bOrder),
		uintptr(ulAf),
		uintptr(tableClass),
		uintptr(reserved))
	return uint32(ret1)
}

func GetExtendedUdpTable(pUdpTable uintptr, pdwSize *DWORD, bOrder bool, ulAf ULONG, tableClass UDP_TABLE_CLASS, reserved ULONG) uint32 {
	ret1 := syscall6(getExtendedUdpTable, 6,
		pUdpTable,
		uintptr(unsafe.Pointer(pdwSize)),
		getUintptrFromBool(bOrder),
		uintptr(ulAf),
		uintptr(tableClass),
		uintptr(reserved))
	return uint32(ret1)
}

func GetFriendlyIfIndex(ifIndex uint32) uint32 {
	ret1 := syscall3(getFriendlyIfIndex, 1,
		uintptr(ifIndex),
		0,
		0)
	return uint32(ret1)
}

func GetIcmpStatistics(statistics PMIB_ICMP) ULONG {
	ret1 := syscall3(getIcmpStatistics, 1,
		uintptr(unsafe.Pointer(statistics)),
		0,
		0)
	return ULONG(ret1)
}

func GetIcmpStatisticsEx(statistics PMIB_ICMP_EX, family ULONG) ULONG {
	ret1 := syscall3(getIcmpStatisticsEx, 2,
		uintptr(unsafe.Pointer(statistics)),
		uintptr(family),
		0)
	return ULONG(ret1)
}

func GetIfEntry(pIfRow PMIB_IFROW) uint32 {
	ret1 := syscall3(getIfEntry, 1,
		uintptr(unsafe.Pointer(pIfRow)),
		0,
		0)
	return uint32(ret1)
}

func GetIfTable(pIfTable PMIB_IFTABLE, pdwSize *uint32, bOrder bool) uint32 {
	ret1 := syscall3(getIfTable, 3,
		uintptr(unsafe.Pointer(pIfTable)),
		uintptr(unsafe.Pointer(pdwSize)),
		getUintptrFromBool(bOrder))
	return uint32(ret1)
}

func GetInterfaceInfo(pIfTable PIP_INTERFACE_INFO, dwOutBufLen *uint32) uint32 {
	ret1 := syscall3(getInterfaceInfo, 2,
		uintptr(unsafe.Pointer(pIfTable)),
		uintptr(unsafe.Pointer(dwOutBufLen)),
		0)
	return uint32(ret1)
}

func GetIpAddrTable(pIpAddrTable PMIB_IPADDRTABLE, pdwSize *uint32, bOrder bool) uint32 {
	ret1 := syscall3(getIpAddrTable, 3,
		uintptr(unsafe.Pointer(pIpAddrTable)),
		uintptr(unsafe.Pointer(pdwSize)),
		getUintptrFromBool(bOrder))
	return uint32(ret1)
}

func GetIpErrorString(errorCode IP_STATUS, buffer PWSTR, size *DWORD) uint32 {
	ret1 := syscall3(getIpErrorString, 3,
		uintptr(errorCode),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(size)))
	return uint32(ret1)
}

func GetIpForwardTable(pIpForwardTable PMIB_IPFORWARDTABLE, pdwSize *uint32, bOrder bool) uint32 {
	ret1 := syscall3(getIpForwardTable, 3,
		uintptr(unsafe.Pointer(pIpForwardTable)),
		uintptr(unsafe.Pointer(pdwSize)),
		getUintptrFromBool(bOrder))
	return uint32(ret1)
}

func GetIpNetTable(ipNetTable PMIB_IPNETTABLE, sizePointer *uint32, order bool) ULONG {
	ret1 := syscall3(getIpNetTable, 3,
		uintptr(unsafe.Pointer(ipNetTable)),
		uintptr(unsafe.Pointer(sizePointer)),
		getUintptrFromBool(order))
	return ULONG(ret1)
}

func GetIpStatistics(statistics PMIB_IPSTATS) ULONG {
	ret1 := syscall3(getIpStatistics, 1,
		uintptr(unsafe.Pointer(statistics)),
		0,
		0)
	return ULONG(ret1)
}

func GetIpStatisticsEx(statistics PMIB_IPSTATS, family ULONG) ULONG {
	ret1 := syscall3(getIpStatisticsEx, 2,
		uintptr(unsafe.Pointer(statistics)),
		uintptr(family),
		0)
	return ULONG(ret1)
}

func GetNetworkParams(pFixedInfo PFIXED_INFO, pOutBufLen *uint32) uint32 {
	ret1 := syscall3(getNetworkParams, 2,
		uintptr(unsafe.Pointer(pFixedInfo)),
		uintptr(unsafe.Pointer(pOutBufLen)),
		0)
	return uint32(ret1)
}

func GetNumberOfInterfaces(pdwNumIf *DWORD) uint32 {
	ret1 := syscall3(getNumberOfInterfaces, 1,
		uintptr(unsafe.Pointer(pdwNumIf)),
		0,
		0)
	return uint32(ret1)
}

func GetOwnerModuleFromTcp6Entry(pTcpEntry PMIB_TCP6ROW_OWNER_MODULE, class TCPIP_OWNER_MODULE_INFO_CLASS, pBuffer uintptr, pdwSize *DWORD) uint32 {
	ret1 := syscall6(getOwnerModuleFromTcp6Entry, 4,
		uintptr(unsafe.Pointer(pTcpEntry)),
		uintptr(class),
		pBuffer,
		uintptr(unsafe.Pointer(pdwSize)),
		0,
		0)
	return uint32(ret1)
}

func GetOwnerModuleFromTcpEntry(pTcpEntry PMIB_TCPROW_OWNER_MODULE, class TCPIP_OWNER_MODULE_INFO_CLASS, pBuffer uintptr, pdwSize *DWORD) uint32 {
	ret1 := syscall6(getOwnerModuleFromTcpEntry, 4,
		uintptr(unsafe.Pointer(pTcpEntry)),
		uintptr(class),
		pBuffer,
		uintptr(unsafe.Pointer(pdwSize)),
		0,
		0)
	return uint32(ret1)
}

func GetOwnerModuleFromUdp6Entry(pUdpEntry PMIB_UDP6ROW_OWNER_MODULE, class TCPIP_OWNER_MODULE_INFO_CLASS, pBuffer uintptr, pdwSize *DWORD) uint32 {
	ret1 := syscall6(getOwnerModuleFromUdp6Entry, 4,
		uintptr(unsafe.Pointer(pUdpEntry)),
		uintptr(class),
		pBuffer,
		uintptr(unsafe.Pointer(pdwSize)),
		0,
		0)
	return uint32(ret1)
}

func GetOwnerModuleFromUdpEntry(pUdpEntry PMIB_UDPROW_OWNER_MODULE, class TCPIP_OWNER_MODULE_INFO_CLASS, pBuffer uintptr, pdwSize *DWORD) uint32 {
	ret1 := syscall6(getOwnerModuleFromUdpEntry, 4,
		uintptr(unsafe.Pointer(pUdpEntry)),
		uintptr(class),
		pBuffer,
		uintptr(unsafe.Pointer(pdwSize)),
		0,
		0)
	return uint32(ret1)
}

func GetPerAdapterInfo(ifIndex ULONG, pPerAdapterInfo PIP_PER_ADAPTER_INFO, pOutBufLen *uint32) uint32 {
	ret1 := syscall3(getPerAdapterInfo, 3,
		uintptr(ifIndex),
		uintptr(unsafe.Pointer(pPerAdapterInfo)),
		uintptr(unsafe.Pointer(pOutBufLen)))
	return uint32(ret1)
}

// TODO: Unknown type(s): PMIB_TCP6ROW, PUCHAR, TCP_ESTATS_TYPE
// func GetPerTcp6ConnectionEStats(row PMIB_TCP6ROW, estatsType TCP_ESTATS_TYPE, rw PUCHAR, rwVersion ULONG, rwSize ULONG, ros PUCHAR, rosVersion ULONG, rosSize ULONG, rod PUCHAR, rodVersion ULONG, rodSize ULONG) ULONG

// TODO: Unknown type(s): PUCHAR, TCP_ESTATS_TYPE
// func GetPerTcpConnectionEStats(row PMIB_TCPROW, estatsType TCP_ESTATS_TYPE, rw PUCHAR, rwVersion ULONG, rwSize ULONG, ros PUCHAR, rosVersion ULONG, rosSize ULONG, rod PUCHAR, rodVersion ULONG, rodSize ULONG) ULONG

func GetRTTAndHopCount(destIpAddress IPAddr, hopCount *uint32, maxHops ULONG, rTT *uint32) bool {
	ret1 := syscall6(getRTTAndHopCount, 4,
		uintptr(destIpAddress),
		uintptr(unsafe.Pointer(hopCount)),
		uintptr(maxHops),
		uintptr(unsafe.Pointer(rTT)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PMIB_TCP6TABLE
// func GetTcp6Table(tcpTable PMIB_TCP6TABLE, sizePointer *uint32, order bool) ULONG

// TODO: Unknown type(s): PMIB_TCP6TABLE2
// func GetTcp6Table2(tcpTable PMIB_TCP6TABLE2, sizePointer *uint32, order bool) ULONG

// TODO: Unknown type(s): PMIB_TCPSTATS
// func GetTcpStatistics(statistics PMIB_TCPSTATS) ULONG

// TODO: Unknown type(s): PMIB_TCPSTATS
// func GetTcpStatisticsEx(statistics PMIB_TCPSTATS, family ULONG) ULONG

// TODO: Unknown type(s): PMIB_TCPTABLE
// func GetTcpTable(tcpTable PMIB_TCPTABLE, sizePointer *uint32, order bool) ULONG

// TODO: Unknown type(s): PMIB_TCPTABLE2
// func GetTcpTable2(tcpTable PMIB_TCPTABLE2, sizePointer *uint32, order bool) ULONG

// TODO: Unknown type(s): PMIB_UDP6TABLE
// func GetUdp6Table(udp6Table PMIB_UDP6TABLE, sizePointer *uint32, order bool) ULONG

// TODO: Unknown type(s): PMIB_UDPSTATS
// func GetUdpStatistics(stats PMIB_UDPSTATS) ULONG

// TODO: Unknown type(s): PMIB_UDPSTATS
// func GetUdpStatisticsEx(statistics PMIB_UDPSTATS, family ULONG) ULONG

// TODO: Unknown type(s): PMIB_UDPTABLE
// func GetUdpTable(udpTable PMIB_UDPTABLE, sizePointer *uint32, order bool) ULONG

// TODO: Unknown type(s): PIP_UNIDIRECTIONAL_ADAPTER_ADDRESS
// func GetUniDirectionalAdapterInfo(pIPIfInfo PIP_UNIDIRECTIONAL_ADAPTER_ADDRESS, dwOutBufLen *uint32) uint32

func Icmp6CreateFile() HANDLE {
	ret1 := syscall3(icmp6CreateFile, 0,
		0,
		0,
		0)
	return HANDLE(ret1)
}

func Icmp6ParseReplies(replyBuffer LPVOID, replySize uint32) uint32 {
	ret1 := syscall3(icmp6ParseReplies, 2,
		uintptr(unsafe.Pointer(replyBuffer)),
		uintptr(replySize),
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PIO_APC_ROUTINE, PIP_OPTION_INFORMATION, struct sockaddr_in6 *
// func Icmp6SendEcho2(icmpHandle HANDLE, event HANDLE, apcRoutine PIO_APC_ROUTINE, apcContext uintptr, sourceAddress struct sockaddr_in6 *, destinationAddress struct sockaddr_in6 *, requestData LPVOID, requestSize uint16, requestOptions PIP_OPTION_INFORMATION, replyBuffer LPVOID, replySize uint32, timeout uint32) uint32

func IcmpCloseHandle(icmpHandle HANDLE) bool {
	ret1 := syscall3(icmpCloseHandle, 1,
		uintptr(icmpHandle),
		0,
		0)
	return ret1 != 0
}

func IcmpCreateFile() HANDLE {
	ret1 := syscall3(icmpCreateFile, 0,
		0,
		0,
		0)
	return HANDLE(ret1)
}

func IcmpParseReplies(replyBuffer LPVOID, replySize uint32) uint32 {
	ret1 := syscall3(icmpParseReplies, 2,
		uintptr(unsafe.Pointer(replyBuffer)),
		uintptr(replySize),
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PIP_OPTION_INFORMATION
// func IcmpSendEcho(icmpHandle HANDLE, destinationAddress IPAddr, requestData LPVOID, requestSize uint16, requestOptions PIP_OPTION_INFORMATION, replyBuffer LPVOID, replySize uint32, timeout uint32) uint32

// TODO: Unknown type(s): PIO_APC_ROUTINE, PIP_OPTION_INFORMATION
// func IcmpSendEcho2(icmpHandle HANDLE, event HANDLE, apcRoutine PIO_APC_ROUTINE, apcContext uintptr, destinationAddress IPAddr, requestData LPVOID, requestSize uint16, requestOptions PIP_OPTION_INFORMATION, replyBuffer LPVOID, replySize uint32, timeout uint32) uint32

// TODO: Unknown type(s): PIP_ADAPTER_INDEX_MAP
// func IpReleaseAddress(adapterInfo PIP_ADAPTER_INDEX_MAP) uint32

// TODO: Unknown type(s): PIP_ADAPTER_INDEX_MAP
// func IpRenewAddress(adapterInfo PIP_ADAPTER_INDEX_MAP) uint32

func LookupPersistentTcpPortReservation(startPort USHORT, numberOfPorts USHORT, token PULONG64) ULONG {
	ret1 := syscall3(lookupPersistentTcpPortReservation, 3,
		uintptr(startPort),
		uintptr(numberOfPorts),
		uintptr(unsafe.Pointer(token)))
	return ULONG(ret1)
}

func LookupPersistentUdpPortReservation(startPort USHORT, numberOfPorts USHORT, token PULONG64) ULONG {
	ret1 := syscall3(lookupPersistentUdpPortReservation, 3,
		uintptr(startPort),
		uintptr(numberOfPorts),
		uintptr(unsafe.Pointer(token)))
	return ULONG(ret1)
}

// TODO: Unknown type(s): IP_INTERFACE_NAME_INFO * *
// func NhpAllocateAndGetInterfaceInfoFromStack(ppTable IP_INTERFACE_NAME_INFO * *, pdwCount *DWORD, bOrder bool, hHeap HANDLE, dwFlags uint32) uint32

func NotifyAddrChange(handle *HANDLE, overlapped *OVERLAPPED) uint32 {
	ret1 := syscall3(notifyAddrChange, 2,
		uintptr(unsafe.Pointer(handle)),
		uintptr(unsafe.Pointer(overlapped)),
		0)
	return uint32(ret1)
}

func NotifyRouteChange(handle *HANDLE, overlapped *OVERLAPPED) uint32 {
	ret1 := syscall3(notifyRouteChange, 2,
		uintptr(unsafe.Pointer(handle)),
		uintptr(unsafe.Pointer(overlapped)),
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PNET_ADDRESS_INFO
// func ParseNetworkString(networkString /*const*/ *WCHAR, types uint32, addressInfo PNET_ADDRESS_INFO, portNumber *USHORT, prefixLength *byte) uint32

func ResolveNeighbor(networkAddress *SOCKADDR, physicalAddress uintptr, physicalAddressLength *uint32) ULONG {
	ret1 := syscall3(resolveNeighbor, 3,
		uintptr(unsafe.Pointer(networkAddress)),
		physicalAddress,
		uintptr(unsafe.Pointer(physicalAddressLength)))
	return ULONG(ret1)
}

func RestoreMediaSense(pOverlapped *OVERLAPPED, lpdwEnableCount *uint32) uint32 {
	ret1 := syscall3(restoreMediaSense, 2,
		uintptr(unsafe.Pointer(pOverlapped)),
		uintptr(unsafe.Pointer(lpdwEnableCount)),
		0)
	return uint32(ret1)
}

func SendARP(destIP IPAddr, srcIP IPAddr, pMacAddr uintptr, phyAddrLen *uint32) uint32 {
	ret1 := syscall6(sendARP, 4,
		uintptr(destIP),
		uintptr(srcIP),
		pMacAddr,
		uintptr(unsafe.Pointer(phyAddrLen)),
		0,
		0)
	return uint32(ret1)
}

func SetIfEntry(pIfRow PMIB_IFROW) uint32 {
	ret1 := syscall3(setIfEntry, 1,
		uintptr(unsafe.Pointer(pIfRow)),
		0,
		0)
	return uint32(ret1)
}

func SetIpForwardEntry(pRoute PMIB_IPFORWARDROW) uint32 {
	ret1 := syscall3(setIpForwardEntry, 1,
		uintptr(unsafe.Pointer(pRoute)),
		0,
		0)
	return uint32(ret1)
}

func SetIpNetEntry(pArpEntry PMIB_IPNETROW) uint32 {
	ret1 := syscall3(setIpNetEntry, 1,
		uintptr(unsafe.Pointer(pArpEntry)),
		0,
		0)
	return uint32(ret1)
}

func SetIpStatistics(pIpStats PMIB_IPSTATS) uint32 {
	ret1 := syscall3(setIpStatistics, 1,
		uintptr(unsafe.Pointer(pIpStats)),
		0,
		0)
	return uint32(ret1)
}

func SetIpStatisticsEx(statistics PMIB_IPSTATS, family ULONG) ULONG {
	ret1 := syscall3(setIpStatisticsEx, 2,
		uintptr(unsafe.Pointer(statistics)),
		uintptr(family),
		0)
	return ULONG(ret1)
}

func SetIpTTL(nTTL uint32) uint32 {
	ret1 := syscall3(setIpTTL, 1,
		uintptr(nTTL),
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PMIB_TCP6ROW, PUCHAR, TCP_ESTATS_TYPE
// func SetPerTcp6ConnectionEStats(row PMIB_TCP6ROW, estatsType TCP_ESTATS_TYPE, rw PUCHAR, rwVersion ULONG, rwSize ULONG, offset ULONG) ULONG

// TODO: Unknown type(s): PUCHAR, TCP_ESTATS_TYPE
// func SetPerTcpConnectionEStats(row PMIB_TCPROW, estatsType TCP_ESTATS_TYPE, rw PUCHAR, rwVersion ULONG, rwSize ULONG, offset ULONG) ULONG

func SetTcpEntry(pTcpRow PMIB_TCPROW) uint32 {
	ret1 := syscall3(setTcpEntry, 1,
		uintptr(unsafe.Pointer(pTcpRow)),
		0,
		0)
	return uint32(ret1)
}

func UnenableRouter(pOverlapped *OVERLAPPED, lpdwEnableCount *uint32) uint32 {
	ret1 := syscall3(unenableRouter, 2,
		uintptr(unsafe.Pointer(pOverlapped)),
		uintptr(unsafe.Pointer(lpdwEnableCount)),
		0)
	return uint32(ret1)
}

func AllocateAndGetIfTableFromStack(ppIfTable *PMIB_IFTABLE, bOrder bool, heap HANDLE, flags uint32) uint32 {
	ret1 := syscall6(allocateAndGetIfTableFromStack, 4,
		uintptr(unsafe.Pointer(ppIfTable)),
		getUintptrFromBool(bOrder),
		uintptr(heap),
		uintptr(flags),
		0,
		0)
	return uint32(ret1)
}

func AllocateAndGetIpAddrTableFromStack(ppIpAddrTable *PMIB_IPADDRTABLE, bOrder bool, heap HANDLE, flags uint32) uint32 {
	ret1 := syscall6(allocateAndGetIpAddrTableFromStack, 4,
		uintptr(unsafe.Pointer(ppIpAddrTable)),
		getUintptrFromBool(bOrder),
		uintptr(heap),
		uintptr(flags),
		0,
		0)
	return uint32(ret1)
}

func AllocateAndGetIpForwardTableFromStack(ppIpForwardTable *PMIB_IPFORWARDTABLE, bOrder bool, heap HANDLE, flags uint32) uint32 {
	ret1 := syscall6(allocateAndGetIpForwardTableFromStack, 4,
		uintptr(unsafe.Pointer(ppIpForwardTable)),
		getUintptrFromBool(bOrder),
		uintptr(heap),
		uintptr(flags),
		0,
		0)
	return uint32(ret1)
}

func AllocateAndGetIpNetTableFromStack(ppIpNetTable *PMIB_IPNETTABLE, bOrder bool, heap HANDLE, flags uint32) uint32 {
	ret1 := syscall6(allocateAndGetIpNetTableFromStack, 4,
		uintptr(unsafe.Pointer(ppIpNetTable)),
		getUintptrFromBool(bOrder),
		uintptr(heap),
		uintptr(flags),
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PMIB_TCPTABLE *
// func AllocateAndGetTcpTableFromStack(ppTcpTable PMIB_TCPTABLE *, bOrder bool, heap HANDLE, flags uint32) uint32

// TODO: Unknown type(s): PMIB_UDPTABLE *
// func AllocateAndGetUdpTableFromStack(ppUdpTable PMIB_UDPTABLE *, bOrder bool, heap HANDLE, flags uint32) uint32

func CancelMibChangeNotify2(handle HANDLE) uint32 {
	ret1 := syscall3(cancelMibChangeNotify2, 1,
		uintptr(handle),
		0,
		0)
	return uint32(ret1)
}

func ConvertInterfaceGuidToLuid(guid /*const*/ *GUID, luid *NET_LUID) uint32 {
	ret1 := syscall3(convertInterfaceGuidToLuid, 2,
		uintptr(unsafe.Pointer(guid)),
		uintptr(unsafe.Pointer(luid)),
		0)
	return uint32(ret1)
}

func ConvertInterfaceIndexToLuid(index NET_IFINDEX, luid *NET_LUID) uint32 {
	ret1 := syscall3(convertInterfaceIndexToLuid, 2,
		uintptr(index),
		uintptr(unsafe.Pointer(luid)),
		0)
	return uint32(ret1)
}

func ConvertInterfaceLuidToGuid(luid /*const*/ *NET_LUID, guid *GUID) uint32 {
	ret1 := syscall3(convertInterfaceLuidToGuid, 2,
		uintptr(unsafe.Pointer(luid)),
		uintptr(unsafe.Pointer(guid)),
		0)
	return uint32(ret1)
}

func ConvertInterfaceLuidToIndex(luid /*const*/ *NET_LUID, index *NET_IFINDEX) uint32 {
	ret1 := syscall3(convertInterfaceLuidToIndex, 2,
		uintptr(unsafe.Pointer(luid)),
		uintptr(unsafe.Pointer(index)),
		0)
	return uint32(ret1)
}

func ConvertInterfaceLuidToName(luid /*const*/ *NET_LUID, name *WCHAR, aLen SIZE_T) uint32 {
	ret1 := syscall3(convertInterfaceLuidToName, 3,
		uintptr(unsafe.Pointer(luid)),
		uintptr(unsafe.Pointer(name)),
		uintptr(aLen))
	return uint32(ret1)
}

func ConvertInterfaceNameToLuid(name /*const*/ *WCHAR, luid *NET_LUID) uint32 {
	ret1 := syscall3(convertInterfaceNameToLuid, 2,
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(luid)),
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PSOCKADDR_IN6_PAIR *, const PSOCKADDR_IN6
// func CreateSortedAddressPairs(src_list /*const*/ const PSOCKADDR_IN6, src_count uint32, dst_list /*const*/ const PSOCKADDR_IN6, dst_count uint32, options uint32, pair_list PSOCKADDR_IN6_PAIR *, pair_count *uint32) uint32

func FreeMibTable(ptr uintptr) {
	syscall3(freeMibTable, 1,
		ptr,
		0,
		0)
}

// TODO: Unknown type(s): MIB_IF_ROW2 *
// func GetIfEntry2(row2 MIB_IF_ROW2 *) uint32

// TODO: Unknown type(s): MIB_IF_TABLE2 * *
// func GetIfTable2(table MIB_IF_TABLE2 * *) uint32

// TODO: Unknown type(s): PIO_APC_ROUTINE, PIP_OPTION_INFORMATION
// func IcmpSendEcho2Ex(icmpHandle HANDLE, event HANDLE, apcRoutine PIO_APC_ROUTINE, apcContext uintptr, sourceAddress IPAddr, destinationAddress IPAddr, requestData LPVOID, requestSize uint16, requestOptions PIP_OPTION_INFORMATION, replyBuffer LPVOID, replySize uint32, timeout uint32) uint32

func NotifyIpInterfaceChange(family ULONG, callback uintptr, context uintptr, init_notify BOOLEAN, handle *HANDLE) uint32 {
	ret1 := syscall6(notifyIpInterfaceChange, 5,
		uintptr(family),
		callback,
		context,
		uintptr(init_notify),
		uintptr(unsafe.Pointer(handle)),
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PFADDRESSTYPE
// func PfBindInterfaceToIPAddress(aInterface INTERFACE_HANDLE, aType PFADDRESSTYPE, ip *byte) uint32

// TODO: Unknown type(s): PFFORWARD_ACTION
// func PfCreateInterface(dwName uint32, inAction PFFORWARD_ACTION, outAction PFFORWARD_ACTION, bUseLog bool, bMustBeUnique bool, ppInterface *INTERFACE_HANDLE) uint32

func PfDeleteInterface(aInterface INTERFACE_HANDLE) uint32 {
	ret1 := syscall3(pfDeleteInterface, 1,
		uintptr(aInterface),
		0,
		0)
	return uint32(ret1)
}

func PfUnBindInterface(aInterface INTERFACE_HANDLE) uint32 {
	ret1 := syscall3(pfUnBindInterface, 1,
		uintptr(aInterface),
		0,
		0)
	return uint32(ret1)
}
