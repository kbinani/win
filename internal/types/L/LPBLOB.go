package win

type tagBLOB struct {
	CbSize    ULONG
	PBlobData *BYTE
}
type LPBLOB *tagBLOB
