package win

//ref DWORD
//ref SE_OBJECT_TYPE
//ref LPWSTR
type OBJECTS_AND_NAME struct {
	ObjectsPresent          DWORD
	ObjectType              SE_OBJECT_TYPE
	ObjectTypeName          LPWSTR
	InheritedObjectTypeName LPWSTR
	PtstrName               LPWSTR
}
