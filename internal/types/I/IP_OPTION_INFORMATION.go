package win

//ref UCHAR
//ref PUCHAR

type IP_OPTION_INFORMATION struct {
	Ttl         UCHAR
	Tos         UCHAR
	Flags       UCHAR
	OptionsSize UCHAR
	OptionsData PUCHAR
}
