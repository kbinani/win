package win

//ref DWORD
type MIBICMPSTATS struct {
	DwMsgs          DWORD
	DwErrors        DWORD
	DwDestUnreachs  DWORD
	DwTimeExcds     DWORD
	DwParmProbs     DWORD
	DwSrcQuenchs    DWORD
	DwRedirects     DWORD
	DwEchos         DWORD
	DwEchoReps      DWORD
	DwTimestamps    DWORD
	DwTimestampReps DWORD
	DwAddrMasks     DWORD
	DwAddrMaskReps  DWORD
}
