package win

type ARRAY_INFO struct {
	Dimension             int32
	BufferConformanceMark *uint32
	BufferVarianceMark    *uint32
	MaxCountArray         *uint32
	OffsetArray           *uint32
	ActualCountArray      *uint32
}
