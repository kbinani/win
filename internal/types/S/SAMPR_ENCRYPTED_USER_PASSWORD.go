package win

//ref UCHAR

type SAMPR_ENCRYPTED_USER_PASSWORD struct {
	Buffer [(256 * 2) + 4]UCHAR
}
