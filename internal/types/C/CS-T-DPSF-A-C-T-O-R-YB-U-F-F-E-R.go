package win

//ref LONG
//ref ProxyFileInfo

type CStdPSFactoryBuffer struct {
	lpVtbl         uintptr
	RefCount       LONG
	PProxyFileList **ProxyFileInfo
	Filler1        LONG
}
