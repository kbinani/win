package win

//ref LSA_UNICODE_STRING
//ref PSID

type LSA_TRUST_INFORMATION struct {
	Name LSA_UNICODE_STRING
	Sid  PSID
}
