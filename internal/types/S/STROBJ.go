package win

//ref ULONG
//ref FLONG
//ref RECTL
//ref GLYPHPOS
//ref LPWSTR
type STROBJ struct {
	CGlyphs     ULONG
	FlAccel     FLONG
	UlCharInc   ULONG
	RclBkGround RECTL
	Pgp         *GLYPHPOS
	PwszOrg     LPWSTR
}
