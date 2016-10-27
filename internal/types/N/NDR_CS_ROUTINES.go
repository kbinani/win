package win

//ref NDR_CS_SIZE_CONVERT_ROUTINES
//ref CS_TAG_GETTING_ROUTINE
type NDR_CS_ROUTINES struct {
	PSizeConvertRoutines *NDR_CS_SIZE_CONVERT_ROUTINES // NDR_CS_SIZE_CONVERT_ROUTINES
	PTagGettingRoutines  uintptr                       // CS_TAG_GETTING_ROUTINE
}
