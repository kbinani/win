package win

//ref DWORD
//ref ACCESS_MODE
//ref DWORD
//ref TRUSTEE
type EXPLICIT_ACCESS struct {
	GrfAccessPermissions DWORD
	GrfAccessMode        ACCESS_MODE
	GrfInheritance       DWORD
	Trustee              TRUSTEE
}
