package win

//ref USHORT
//ref GUID
type OBJECT_TYPE_LIST struct {
	Level      USHORT
	Sbz        USHORT
	ObjectType *GUID
}
