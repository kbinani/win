package win

//ref CHAR
//ref IID
//ref LONG_PTR
//ref IIDLookupRtn

type ProxyFileInfo struct {
	pProxyVtblList             uintptr // PCInterfaceProxyVtblList
	pStubVtblList              uintptr // PCInterfaceStubVtblList
	PNamesArray/*const*/ *CHAR // const PCInterfaceName*
	PDelegatedIIDs/*const*/ **IID
	PIIDLookupRtn/*const*/ uintptr // IIDLookupRtn
	TableSize                      uint16
	TableVersion                   uint16
	PAsyncIIDLookup/*const*/ **IID
	Filler2 LONG_PTR
	Filler3 LONG_PTR
	Filler4 LONG_PTR
}
