package win

//ref MULTIPLE_TRUSTEE_OPERATION
//ref TRUSTEE_FORM
//ref TRUSTEE_TYPE
//ref LPWSTR
type TRUSTEE struct {
	PMultipleTrustee         *TRUSTEE
	MultipleTrusteeOperation MULTIPLE_TRUSTEE_OPERATION
	TrusteeForm              TRUSTEE_FORM
	TrusteeType              TRUSTEE_TYPE
	PtstrName                LPWSTR
}
