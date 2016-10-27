package win

//ref DWORD
//ref GUID
//ref SID
type OBJECTS_AND_SID struct {
	ObjectsPresent          DWORD
	ObjectTypeGuid          GUID
	InheritedObjectTypeGuid GUID
	PSid                    *SID
}
