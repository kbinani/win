package win

//ref ULONG
//ref SERVICETYPE
type FLOWSPEC struct {
	TokenRate          ULONG
	TokenBucketSize    ULONG
	PeakBandwidth      ULONG
	Latency            ULONG
	DelayVariation     ULONG
	ServiceType        SERVICETYPE
	MaxSduSize         ULONG
	MinimumPolicedSize ULONG
}
