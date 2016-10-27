package win

//ref XMIT_HELPER_ROUTINE
type XMIT_ROUTINE_QUINTUPLE struct {
	PfnTranslateToXmit   uintptr // XMIT_HELPER_ROUTINE
	PfnTranslateFromXmit uintptr // XMIT_HELPER_ROUTINE
	PfnFreeXmit          uintptr // XMIT_HELPER_ROUTINE
	PfnFreeInst          uintptr // XMIT_HELPER_ROUTINE
}
