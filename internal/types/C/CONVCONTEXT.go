package win

//ref SECURITY_QUALITY_OF_SERVICE
type CONVCONTEXT struct {
	Cb         uint32
	WFlags     uint32
	WCountryID uint32
	ICodePage  int32
	DwLangID   uint32
	DwSecurity uint32
	Qos        SECURITY_QUALITY_OF_SERVICE
}
