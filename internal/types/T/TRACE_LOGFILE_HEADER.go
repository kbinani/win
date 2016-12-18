package win

//ref ULONG
//ref LARGE_INTEGER
//ref GUID
//ref LPWSTR
//ref TIME_ZONE_INFORMATION

type TRACE_LOGFILE_HEADER struct {
	BufferSize         ULONG
	union1             ULONG
	ProviderVersion    ULONG
	NumberOfProcessors ULONG
	EndTime            LARGE_INTEGER
	TimerResolution    ULONG
	MaximumFileSize    ULONG
	LogFileMode        ULONG
	BuffersWritten     ULONG
	union2             GUID
	LoggerName         LPWSTR
	LogFileName        LPWSTR
	TimeZone           TIME_ZONE_INFORMATION
	padding0           [pad0for64_4for32]byte
	BootTime           LARGE_INTEGER
	PerfFreq           LARGE_INTEGER
	StartTime          LARGE_INTEGER
	ReservedFlags      ULONG
	BuffersLost        ULONG
}

type TRACE_LOGFILE_HEADER_VersionDetail struct {
	MajorVersion    UCHAR
	MinorVersion    UCHAR
	SubVersion      UCHAR
	SubMinorVersion UCHAR
}

func (this *TRACE_LOGFILE_HEADER) GetVersion() ULONG {
	return this.union1
}

func (this *TRACE_LOGFILE_HEADER) SetVersion(v ULONG) {
	this.union1 = v
}

func (this *TRACE_LOGFILE_HEADER) GetVersionDetail() TRACE_LOGFILE_HEADER_VersionDetail {
	return *(*TRACE_LOGFILE_HEADER_VersionDetail)(unsafe.Pointer(&this.union1))
}

func (this *TRACE_LOGFILE_HEADER) SetVersionDetail(v TRACE_LOGFILE_HEADER_VersionDetail) {
	*(*TRACE_LOGFILE_HEADER_VersionDetail)(unsafe.Pointer(&this.union1)) = v
}

func (this *TRACE_LOGFILE_HEADER) GetLogInstanceGuid() GUID {
	return this.union2
}

func (this *TRACE_LOGFILE_HEADER) SetLogInstanceGuid(v GUID) {
	this.union2 = v
}

func (this *TRACE_LOGFILE_HEADER) GetStartBuffers() ULONG {
	return *(*ULONG)(unsafe.Pointer(&this.union2))
}

func (this *TRACE_LOGFILE_HEADER) SetStartBuffers(v ULONG) {
	*(*ULONG)(unsafe.Pointer(&this.union2)) = v
}

func (this *TRACE_LOGFILE_HEADER) GetPointerSize() ULONG {
	return *(*ULONG)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union2)) + uintptr(4)))
}

func (this *TRACE_LOGFILE_HEADER) SetPointerSize(v ULONG) {
	*(*ULONG)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union2)) + uintptr(4))) = v
}

func (this *TRACE_LOGFILE_HEADER) GetEventsLost() ULONG {
	return *(*ULONG)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union2)) + uintptr(8)))
}

func (this *TRACE_LOGFILE_HEADER) SetEventsLost(v ULONG) {
	*(*ULONG)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union2)) + uintptr(8))) = v
}

func (this *TRACE_LOGFILE_HEADER) GetCpuSpeedInMHz() ULONG {
	return *(*ULONG)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union2)) + uintptr(12)))
}

func (this *TRACE_LOGFILE_HEADER) SetCpuSpeedInMHz(v ULONG) {
	*(*ULONG)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union2)) + uintptr(12))) = v
}
