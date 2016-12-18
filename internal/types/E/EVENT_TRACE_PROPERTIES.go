package win

//ref WNODE_HEADER
//ref ULONG
//ref LONG
//ref HANDLE

func (this *EVENT_TRACE_PROPERTIES) GetAgeLimit() LONG {
	return this.union1
}

func (this *EVENT_TRACE_PROPERTIES) SetAgeLimit(v LONG) {
	this.union1 = v
}

func (this *EVENT_TRACE_PROPERTIES) GetFlushThreshold() LONG {
	return this.union1
}

func (this *EVENT_TRACE_PROPERTIES) SetFlushThreshold(v LONG) {
	this.union1 = v
}
