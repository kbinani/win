package win

//ref ACCESS_MASK
type GENERIC_MAPPING struct {
	GenericRead    ACCESS_MASK
	GenericWrite   ACCESS_MASK
	GenericExecute ACCESS_MASK
	GenericAll     ACCESS_MASK
}
