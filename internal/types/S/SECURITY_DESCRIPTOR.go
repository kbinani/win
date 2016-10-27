package win

//ref SECURITY_DESCRIPTOR_CONTROL
//ref PSID
//ref ACL
type SECURITY_DESCRIPTOR struct {
	Revision byte
	Sbz1     byte
	Control  SECURITY_DESCRIPTOR_CONTROL
	Owner    PSID
	Group    PSID
	Sacl     *ACL
	Dacl     *ACL
}
