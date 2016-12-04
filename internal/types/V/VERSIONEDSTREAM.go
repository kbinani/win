package win

//ref GUID
//ref IStream

type VERSIONEDSTREAM struct {
	GuidVersion GUID
	PStream     *IStream
}
