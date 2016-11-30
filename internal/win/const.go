package win

type BOOL int32

const (
	TRUE  BOOL = 1
	FALSE BOOL = 0
)

const (
	CCHDEVICENAME = 32
	CCHFORMNAME   = 32
)

type SECURITY_IMPERSONATION_LEVEL int32

const (
	SecurityAnonymous SECURITY_IMPERSONATION_LEVEL = iota
	SecurityIdentification
	SecurityImpersonation
	SecurityDelegation
)

const (
	CCHILDREN_TITLEBAR  = 5
	CCHILDREN_SCROLLBAR = 5
)

type SE_OBJECT_TYPE int32

const (
	SE_UNKNOWN_OBJECT_TYPE SE_OBJECT_TYPE = iota
	SE_FILE_OBJECT
	SE_SERVICE
	SE_PRINTER
	SE_REGISTRY_KEY
	SE_LMSHARE
	SE_KERNEL_OBJECT
	SE_WINDOW_OBJECT
	SE_DS_OBJECT
	SE_DS_OBJECT_ALL
	SE_PROVIDER_DEFINED_OBJECT
	SE_WMIGUID_OBJECT
	SE_REGISTRY_WOW64_32KEY
)

type PROPERTYORIGIN int32

const (
	PO_STATE PROPERTYORIGIN = iota
	PO_PART
	PO_CLASS
	PO_GLOBAL
	PO_NOTFOUND
)

type THEMESIZE int32

const (
	TS_MIN THEMESIZE = iota
	TS_TRUE
	TS_DRAW
)

const ANYSIZE_ARRAY = 1

type AUDIT_EVENT_TYPE int32

const (
	AuditEventObjectAccess AUDIT_EVENT_TYPE = iota
	AuditEventDirectoryServiceAccess
)

type POLICY_AUDIT_EVENT_TYPE int32

const (
	AuditCategorySystem POLICY_AUDIT_EVENT_TYPE = iota
	AuditCategoryLogon
	AuditCategoryObjectAccess
	AuditCategoryPrivilegeUse
	AuditCategoryDetailedTracking
	AuditCategoryPolicyChange
	AuditCategoryAccountManagement
	AuditCategoryDirectoryServiceAccess
	AuditCategoryAccountLogon
)

type ACCESS_MODE int32

const (
	NOT_USED_ACCESS ACCESS_MODE = iota
	GRANT_ACCESS
	SET_ACCESS
	DENY_ACCESS
	REVOKE_ACCESS
	SET_AUDIT_SUCCESS
	SET_AUDIT_FAILURE
)

type MULTIPLE_TRUSTEE_OPERATION int32

const (
	NO_MULTIPLE_TRUSTEE MULTIPLE_TRUSTEE_OPERATION = iota
	TRUSTEE_IS_IMPERSONATE
)

type TRUSTEE_FORM int32

const (
	TRUSTEE_IS_SID TRUSTEE_FORM = iota
	TRUSTEE_IS_NAME
	TRUSTEE_BAD_FORM
	TRUSTEE_IS_OBJECTS_AND_SID
	TRUSTEE_IS_OBJECTS_AND_NAME
)

type TRUSTEE_TYPE int32

const (
	TRUSTEE_IS_UNKNOWN TRUSTEE_TYPE = iota
	TRUSTEE_IS_USER
	TRUSTEE_IS_GROUP
	TRUSTEE_IS_DOMAIN
	TRUSTEE_IS_ALIAS
	TRUSTEE_IS_WELL_KNOWN_GROUP
	TRUSTEE_IS_DELETED
	TRUSTEE_IS_INVALID
	TRUSTEE_IS_COMPUTER
)

const MAXPNAMELEN = 32

type XLAT_SIDE int32

const (
	XLAT_SERVER XLAT_SIDE = iota
	XLAT_CLIENT
)

type IDL_CS_CONVERT int32

const (
	IDL_CS_NO_CONVERT IDL_CS_CONVERT = iota
	IDL_CS_IN_PLACE_CONVERT
	IDL_CS_NEW_BUFFER_CONVERT
)

type VARENUM int32

const (
	VT_EMPTY            VARENUM = 0
	VT_NULL                     = 1
	VT_I2                       = 2
	VT_I4                       = 3
	VT_R4                       = 4
	VT_R8                       = 5
	VT_CY                       = 6
	VT_DATE                     = 7
	VT_BSTR                     = 8
	VT_DISPATCH                 = 9
	VT_ERROR                    = 10
	VT_BOOL                     = 11
	VT_VARIANT                  = 12
	VT_UNKNOWN                  = 13
	VT_DECIMAL                  = 14
	VT_I1                       = 16
	VT_UI1                      = 17
	VT_UI2                      = 18
	VT_UI4                      = 19
	VT_I8                       = 20
	VT_UI8                      = 21
	VT_INT                      = 22
	VT_UINT                     = 23
	VT_VOID                     = 24
	VT_HRESULT                  = 25
	VT_PTR                      = 26
	VT_SAFEARRAY                = 27
	VT_CARRAY                   = 28
	VT_USERDEFINED              = 29
	VT_LPSTR                    = 30
	VT_LPWSTR                   = 31
	VT_RECORD                   = 36
	VT_INT_PTR                  = 37
	VT_UINT_PTR                 = 38
	VT_FILETIME                 = 64
	VT_BLOB                     = 65
	VT_STREAM                   = 66
	VT_STORAGE                  = 67
	VT_STREAMED_OBJECT          = 68
	VT_STORED_OBJECT            = 69
	VT_BLOB_OBJECT              = 70
	VT_CF                       = 71
	VT_CLSID                    = 72
	VT_VERSIONED_STREAM         = 73
	VT_BSTR_BLOB                = 0xfff
	VT_VECTOR                   = 0x1000
	VT_ARRAY                    = 0x2000
	VT_BYREF                    = 0x4000
	VT_RESERVED                 = 0x8000
	VT_ILLEGAL                  = 0xffff
	VT_ILLEGALMASKED            = 0xfff
	VT_TYPEMASK                 = 0xfff
)

type GpStatus int32

const (
	Ok                        GpStatus = 0
	GenericError                       = 1
	InvalidParameter                   = 2
	OutOfMemory                        = 3
	ObjectBusy                         = 4
	InsufficientBuffer                 = 5
	NotImplemented                     = 6
	Win32Error                         = 7
	WrongState                         = 8
	Aborted                            = 9
	FileNotFound                       = 10
	ValueOverflow                      = 11
	AccessDenied                       = 12
	UnknownImageFormat                 = 13
	FontFamilyNotFound                 = 14
	FontStyleNotFound                  = 15
	NotTrueTypeFont                    = 16
	UnsupportedGdiplusVersion          = 17
	GdiplusNotInitialized              = 18
	PropertyNotFound                   = 19
	PropertyNotSupported               = 20
	ProfileNotFound                    = 21
)

type BP_ANIMATIONSTYLE int32

const (
	BPAS_NONE BP_ANIMATIONSTYLE = iota
	BPAS_LINEAR
	BPAS_CUBIC
	BPAS_SINE
)

type BP_BUFFERFORMAT int32

const (
	BPBF_COMPATIBLEBITMAP BP_BUFFERFORMAT = iota
	BPBF_DIB
	BPBF_TOPDOWNDIB
	BPBF_TOPDOWNMONODIB
)

const MAX_INTLIST_COUNT = 402

type PERCEIVED int32

const (
	PERCEIVED_TYPE_FIRST       PERCEIVED = -3
	PERCEIVED_TYPE_CUSTOM                = -3
	PERCEIVED_TYPE_UNSPECIFIED           = -2
	PERCEIVED_TYPE_FOLDER                = -1
	PERCEIVED_TYPE_UNKNOWN               = 0
	PERCEIVED_TYPE_TEXT                  = 1
	PERCEIVED_TYPE_IMAGE                 = 2
	PERCEIVED_TYPE_AUDIO                 = 3
	PERCEIVED_TYPE_VIDEO                 = 4
	PERCEIVED_TYPE_COMPRESSED            = 5
	PERCEIVED_TYPE_DOCUMENT              = 6
	PERCEIVED_TYPE_SYSTEM                = 7
	PERCEIVED_TYPE_APPLICATION           = 8
	PERCEIVED_TYPE_GAMEMEDIA             = 9
	PERCEIVED_TYPE_CONTACTS              = 10
	PERCEIVED_TYPE_LAST                  = 10
)

type DWORD uint32

type ASSOCF DWORD

const (
	ASSOCF_NONE                 ASSOCF = 0x00000000
	ASSOCF_INIT_NOREMAPCLSID           = 0x00000001
	ASSOCF_INIT_BYEXENAME              = 0x00000002
	ASSOCF_OPEN_BYEXENAME              = 0x00000002
	ASSOCF_INIT_DEFAULTTOSTAR          = 0x00000004
	ASSOCF_INIT_DEFAULTTOFOLDER        = 0x00000008
	ASSOCF_NOUSERSETTINGS              = 0x00000010
	ASSOCF_NOTRUNCATE                  = 0x00000020
	ASSOCF_VERIFY                      = 0x00000040
	ASSOCF_REMAPRUNDLL                 = 0x00000080
	ASSOCF_NOFIXUPS                    = 0x00000100
	ASSOCF_IGNOREBASECLASS             = 0x00000200
	ASSOCF_INIT_IGNOREUNKNOWN          = 0x00000400
	ASSOCF_INIT_FIXED_PROGID           = 0x00000800
	ASSOCF_IS_PROTOCOL                 = 0x00001000
	ASSOCF_INIT_FOR_FILE               = 0x00002000
)

type ASSOCKEY int32

const (
	ASSOCKEY_SHELLEXECCLASS ASSOCKEY = iota + 1
	ASSOCKEY_APP
	ASSOCKEY_CLASS
	ASSOCKEY_BASECLASS
	ASSOCKEY_MAX
)

type ASSOCSTR int32

const (
	ASSOCSTR_COMMAND ASSOCSTR = iota + 1
	ASSOCSTR_EXECUTABLE
	ASSOCSTR_FRIENDLYDOCNAME
	ASSOCSTR_FRIENDLYAPPNAME
	ASSOCSTR_NOOPEN
	ASSOCSTR_SHELLNEWVALUE
	ASSOCSTR_DDECOMMAND
	ASSOCSTR_DDEIFEXEC
	ASSOCSTR_DDEAPPLICATION
	ASSOCSTR_DDETOPIC
	ASSOCSTR_INFOTIP
	ASSOCSTR_QUICKTIP
	ASSOCSTR_TILEINFO
	ASSOCSTR_CONTENTTYPE
	ASSOCSTR_DEFAULTICON
	ASSOCSTR_SHELLEXTENSION
	ASSOCSTR_DROPTARGET
	ASSOCSTR_DELEGATEEXECUTE
	ASSOCSTR_SUPPORTED_URI_PROTOCOLS
	ASSOCSTR_MAX
)

type SHREGDEL_FLAGS int32

const (
	SHREGDEL_DEFAULT SHREGDEL_FLAGS = 0x00000000
	SHREGDEL_HKCU                   = 0x00000001
	SHREGDEL_HKLM                   = 0x00000010
	SHREGDEL_BOTH                   = 0x00000011
)

type SHREGENUM_FLAGS int32

const (
	SHREGENUM_DEFAULT SHREGENUM_FLAGS = 0x00000000
	SHREGENUM_HKCU                    = 0x00000001
	SHREGENUM_HKLM                    = 0x00000010
	SHREGENUM_BOTH                    = 0x00000011
)

type URLIS int32

const (
	URLIS_URL URLIS = iota
	URLIS_OPAQUE
	URLIS_NOHISTORY
	URLIS_FILEURL
	URLIS_APPLIABLE
	URLIS_DIRECTORY
	URLIS_HASQUERY
)

type TASKDIALOG_COMMON_BUTTON_FLAGS int32

const (
	TDCBF_OK_BUTTON     TASKDIALOG_COMMON_BUTTON_FLAGS = 0x0001
	TDCBF_YES_BUTTON                                   = 0x0002
	TDCBF_NO_BUTTON                                    = 0x0004
	TDCBF_CANCEL_BUTTON                                = 0x0008
	TDCBF_RETRY_BUTTON                                 = 0x0010
	TDCBF_CLOSE_BUTTON                                 = 0x0020
)

type TASKDIALOG_FLAGS int32

const (
	TDF_ENABLE_HYPERLINKS           TASKDIALOG_FLAGS = 0x0001
	TDF_USE_HICON_MAIN                               = 0x0002
	TDF_USE_HICON_FOOTER                             = 0x0004
	TDF_ALLOW_DIALOG_CANCELLATION                    = 0x0008
	TDF_USE_COMMAND_LINKS                            = 0x0010
	TDF_USE_COMMAND_LINKS_NO_ICON                    = 0x0020
	TDF_EXPAND_FOOTER_AREA                           = 0x0040
	TDF_EXPANDED_BY_DEFAULT                          = 0x0080
	TDF_VERIFICATION_FLAG_CHECKED                    = 0x0100
	TDF_SHOW_PROGRESS_BAR                            = 0x0200
	TDF_SHOW_MARQUEE_PROGRESS_BAR                    = 0x0400
	TDF_CALLBACK_TIMER                               = 0x0800
	TDF_POSITION_RELATIVE_TO_WINDOW                  = 0x1000
	TDF_RTL_LAYOUT                                   = 0x2000
	TDF_NO_DEFAULT_RADIO_BUTTON                      = 0x4000
	TDF_CAN_BE_MINIMIZED                             = 0x8000
	TDF_NO_SET_FOREGROUND                            = 0x00010000
	TDF_SIZE_TO_CONTENT                              = 0x01000000
)

const MAX_PATH = 260

const LF_FACESIZE = 32
const LF_FULLFACESIZE = 64

const MM_MAX_NUMAXES = 16

type CALLCONV int32

const (
	CC_FASTCALL   CALLCONV = 0
	CC_CDECL               = 1
	CC_MSCPASCAL           = CC_CDECL + 1
	CC_PASCAL              = CC_MSCPASCAL
	CC_MACPASCAL           = CC_PASCAL + 1
	CC_STDCALL             = CC_MACPASCAL + 1
	CC_FPFASTCALL          = CC_STDCALL + 1
	CC_SYSCALL             = CC_FPFASTCALL + 1
	CC_MPWCDECL            = CC_SYSCALL + 1
	CC_MPWPASCAL           = CC_MPWCDECL + 1
	CC_MAX                 = CC_MPWPASCAL + 1
)

type SYSKIND int32

const (
	SYS_WIN16 SYSKIND = 0
	SYS_WIN32         = SYS_WIN16 + 1
	SYS_MAC           = SYS_WIN32 + 1
	SYS_WIN64         = SYS_MAC + 1
)

type REGKIND int32

const (
	REGKIND_DEFAULT REGKIND = iota
	REGKIND_REGISTER
	REGKIND_NONE
)

const IMEMENUITEM_STRING_SIZE = 80

const STYLE_DESCRIPTION_SIZE = 32

const MAX_PROTOCOL_CHAIN = 7

const WSAPROTOCOL_LEN = 255

type ADDRESS_FAMILY int16

const (
	AF_UNSPEC     ADDRESS_FAMILY = 0
	AF_UNIX                      = 1
	AF_INET                      = 2
	AF_IMPLINK                   = 3
	AF_PUP                       = 4
	AF_CHAOS                     = 5
	AF_NS                        = 6
	AF_IPX                       = AF_NS
	AF_ISO                       = 7
	AF_OSI                       = AF_ISO
	AF_ECMA                      = 8
	AF_DATAKIT                   = 9
	AF_CCITT                     = 10
	AF_SNA                       = 11
	AF_DECnet                    = 12
	AF_DLI                       = 13
	AF_LAT                       = 14
	AF_HYLINK                    = 15
	AF_APPLETALK                 = 16
	AF_NETBIOS                   = 17
	AF_VOICEVIEW                 = 18
	AF_FIREFOX                   = 19
	AF_UNKNOWN1                  = 20
	AF_BAN                       = 21
	AF_ATM                       = 22
	AF_INET6                     = 23
	AF_CLUSTER                   = 24
	AF_12844                     = 25
	AF_IRDA                      = 26
	AF_NETDES                    = 28
	AF_TCNPROCESS                = 29
	AF_TCNMESSAGE                = 30
	AF_ICLFXBM                   = 31
	AF_BTH                       = 32
	AF_LINK                      = 33
	AF_MAX                       = 34
)

type SERVICETYPE uint32

const (
	SERVICETYPE_NOTRAFFIC           SERVICETYPE = 0x00000000
	SERVICETYPE_BESTEFFORT                      = 0x00000001
	SERVICETYPE_CONTROLLEDLOAD                  = 0x00000002
	SERVICETYPE_GUARANTEED                      = 0x00000003
	SERVICETYPE_NETWORK_UNAVAILABLE             = 0x00000004
	SERVICETYPE_GENERAL_INFORMATION             = 0x00000005
	SERVICETYPE_NOCHANGE                        = 0x00000006
	SERVICETYPE_NONCONFORMING                   = 0x00000009
	SERVICETYPE_NETWORK_CONTROL                 = 0x0000000A
	SERVICETYPE_QUALITATIVE                     = 0x0000000D
)

const FD_MAX_EVENTS = 10

type WSAECOMPARATOR int32

const (
	COMP_EQUAL WSAECOMPARATOR = iota
	COMP_NOTLESS
)

type WSACOMPLETIONTYPE int32

const (
	NSP_NOTIFY_IMMEDIATELY WSACOMPLETIONTYPE = iota
	NSP_NOTIFY_HWND
	NSP_NOTIFY_EVENT
	NSP_NOTIFY_PORT
	NSP_NOTIFY_APC
)

type WSAESETSERVICEOP int32

const (
	RNRSERVICE_REGISTER WSAESETSERVICEOP = iota
	RNRSERVICE_DEREGISTER
	RNRSERVICE_DELETE
)

const WSADESCRIPTION_LEN = 256
const WSASYS_STATUS_LEN = 128

type MIB_IPFORWARD_TYPE int32

const (
	MIB_IPROUTE_TYPE_OTHER    MIB_IPFORWARD_TYPE = 1
	MIB_IPROUTE_TYPE_INVALID                     = 2
	MIB_IPROUTE_TYPE_DIRECT                      = 3
	MIB_IPROUTE_TYPE_INDIRECT                    = 4
)

type NL_ROUTE_PROTOCOL int32

const (
	RouteProtocolOther   NL_ROUTE_PROTOCOL = 1
	RouteProtocolLocal                     = 2
	RouteProtocolNetMgmt                   = 3
	RouteProtocolIcmp                      = 4
	RouteProtocolEgp                       = 5
	RouteProtocolGgp                       = 6
	RouteProtocolHello                     = 7
	RouteProtocolRip                       = 8
	RouteProtocolIsIs                      = 9
	RouteProtocolEsIs                      = 10
	RouteProtocolCisco                     = 11
	RouteProtocolBbn                       = 12
	RouteProtocolOspf                      = 13
	RouteProtocolBgp                       = 14
	RouteProtocolIdpr                      = 15
	RouteProtocolEigrp                     = 16
	RouteProtocolDvmrp                     = 17
	RouteProtocolRpl                       = 18
	RouteProtocolDhcp                      = 19

	MIB_IPPROTO_OTHER             = 1
	MIB_IPPROTO_LOCAL             = 2
	MIB_IPPROTO_NETMGMT           = 3
	MIB_IPPROTO_ICMP              = 4
	MIB_IPPROTO_EGP               = 5
	MIB_IPPROTO_GGP               = 6
	MIB_IPPROTO_HELLO             = 7
	MIB_IPPROTO_RIP               = 8
	MIB_IPPROTO_IS_IS             = 9
	MIB_IPPROTO_ES_IS             = 10
	MIB_IPPROTO_CISCO             = 11
	MIB_IPPROTO_BBN               = 12
	MIB_IPPROTO_OSPF              = 13
	MIB_IPPROTO_BGP               = 14
	MIB_IPPROTO_IDPR              = 15
	MIB_IPPROTO_EIGRP             = 16
	MIB_IPPROTO_DVMRP             = 17
	MIB_IPPROTO_RPL               = 18
	MIB_IPPROTO_DHCP              = 19
	MIB_IPPROTO_NT_AUTOSTATIC     = 10002
	MIB_IPPROTO_NT_STATIC         = 10006
	MIB_IPPROTO_NT_STATIC_NON_DOD = 10007

	PROTO_IP_OTHER             = 1
	PROTO_IP_LOCAL             = 2
	PROTO_IP_NETMGMT           = 3
	PROTO_IP_ICMP              = 4
	PROTO_IP_EGP               = 5
	PROTO_IP_GGP               = 6
	PROTO_IP_HELLO             = 7
	PROTO_IP_RIP               = 8
	PROTO_IP_IS_IS             = 9
	PROTO_IP_ES_IS             = 10
	PROTO_IP_CISCO             = 11
	PROTO_IP_BBN               = 12
	PROTO_IP_OSPF              = 13
	PROTO_IP_BGP               = 14
	PROTO_IP_IDPR              = 15
	PROTO_IP_EIGRP             = 16
	PROTO_IP_DVMRP             = 17
	PROTO_IP_RPL               = 18
	PROTO_IP_DHCP              = 19
	PROTO_IP_NT_AUTOSTATIC     = 10002
	PROTO_IP_NT_STATIC         = 10006
	PROTO_IP_NT_STATIC_NON_DOD = 10007
)

const MAXLEN_PHYSADDR = 8
const MAXLEN_IFDESCR = 256
const MAX_INTERFACE_NAME_LEN = 256

type MIB_IPNET_TYPE int32

const (
	MIB_IPNET_TYPE_OTHER   MIB_IPNET_TYPE = 1
	MIB_IPNET_TYPE_INVALID                = 2
	MIB_IPNET_TYPE_DYNAMIC                = 3
	MIB_IPNET_TYPE_STATIC                 = 4
)

type MIB_TCP_STATE int32

const (
	MIB_TCP_STATE_CLOSED     MIB_TCP_STATE = 1
	MIB_TCP_STATE_LISTEN                   = 2
	MIB_TCP_STATE_SYN_SENT                 = 3
	MIB_TCP_STATE_SYN_RCVD                 = 4
	MIB_TCP_STATE_ESTAB                    = 5
	MIB_TCP_STATE_FIN_WAIT1                = 6
	MIB_TCP_STATE_FIN_WAIT2                = 7
	MIB_TCP_STATE_CLOSE_WAIT               = 8
	MIB_TCP_STATE_CLOSING                  = 9
	MIB_TCP_STATE_LAST_ACK                 = 10
	MIB_TCP_STATE_TIME_WAIT                = 11
	MIB_TCP_STATE_DELETE_TCB               = 12
)

type MIB_IPSTATS_FORWARDING int32

const (
	MIB_IP_FORWARDING     MIB_IPSTATS_FORWARDING = 1
	MIB_IP_NOT_FORWARDING                        = 2
)

type NL_PREFIX_ORIGIN int32

const (
	IpPrefixOriginOther NL_PREFIX_ORIGIN = iota
	IpPrefixOriginManual
	IpPrefixOriginWellKnown
	IpPrefixOriginDhcp
	IpPrefixOriginRouterAdvertisement
	IpPrefixOriginUnchanged = 1 << 4
)

type NL_SUFFIX_ORIGIN int32

const (
	NlsoOther NL_SUFFIX_ORIGIN = iota
	NlsoManual
	NlsoWellKnown
	NlsoDhcp
	NlsoLinkLayerAddress
	NlsoRandom
)
const (
	IpSuffixOriginOther NL_SUFFIX_ORIGIN = iota
	IpSuffixOriginManual
	IpSuffixOriginWellKnown
	IpSuffixOriginDhcp
	IpSuffixOriginLinkLayerAddress
	IpSuffixOriginRandom
	IpSuffixOriginUnchanged = 1 << 4
)

type NL_DAD_STATE int32

const (
	NldsInvalid NL_DAD_STATE = iota
	NldsTentative
	NldsDuplicate
	NldsDeprecated
	NldsPreferred
)
const (
	IpDadStateInvalid NL_DAD_STATE = iota
	IpDadStateTentative
	IpDadStateDuplicate
	IpDadStateDeprecated
	IpDadStatePreferred
)

const MAX_ADAPTER_NAME_LENGTH = 256
const MAX_ADAPTER_DESCRIPTION_LENGTH = 128
const MAX_ADAPTER_ADDRESS_LENGTH = 8

type TCP_TABLE_CLASS int32

const (
	TCP_TABLE_BASIC_LISTENER TCP_TABLE_CLASS = iota
	TCP_TABLE_BASIC_CONNECTIONS
	TCP_TABLE_BASIC_ALL
	TCP_TABLE_OWNER_PID_LISTENER
	TCP_TABLE_OWNER_PID_CONNECTIONS
	TCP_TABLE_OWNER_PID_ALL
	TCP_TABLE_OWNER_MODULE_LISTENER
	TCP_TABLE_OWNER_MODULE_CONNECTIONS
	TCP_TABLE_OWNER_MODULE_ALL
)

type UDP_TABLE_CLASS int32

const (
	UDP_TABLE_BASIC UDP_TABLE_CLASS = iota
	UDP_TABLE_OWNER_PID
	UDP_TABLE_OWNER_MODULE
)

type INTERNAL_IF_OPER_STATUS int32

const (
	IF_OPER_STATUS_NON_OPERATIONAL INTERNAL_IF_OPER_STATUS = 0
	IF_OPER_STATUS_UNREACHABLE                             = 1
	IF_OPER_STATUS_DISCONNECTED                            = 2
	IF_OPER_STATUS_CONNECTING                              = 3
	IF_OPER_STATUS_CONNECTED                               = 4
	IF_OPER_STATUS_OPERATIONAL                             = 5
)

type IF_OPER_STATUS int32

const (
	IfOperStatusUp IF_OPER_STATUS = iota + 1
	IfOperStatusDown
	IfOperStatusTesting
	IfOperStatusUnknown
	IfOperStatusDormant
	IfOperStatusNotPresent
	IfOperStatusLowerLayerDown
)

const MAX_DHCPV6_DUID_LENGTH = 130

type NET_IF_COMPARTMENT_ID uint32

const (
	NET_IF_COMPARTMENT_ID_UNSPECIFIED NET_IF_COMPARTMENT_ID = 0
	NET_IF_COMPARTMENT_ID_PRIMARY                           = 1
	NET_IF_COMPARTMENT_ID_ALL                               = 0xffffffff
)

type NET_IF_CONNECTION_TYPE int32

const (
	NET_IF_CONNECTION_DEDICATED NET_IF_CONNECTION_TYPE = iota + 1
	NET_IF_CONNECTION_PASSIVE
	NET_IF_CONNECTION_DEMAND
	NET_IF_CONNECTION_MAXIMUM
)

type TUNNEL_TYPE int32

const (
	TUNNEL_TYPE_NONE    TUNNEL_TYPE = 0
	TUNNEL_TYPE_OTHER               = 1
	TUNNEL_TYPE_DIRECT              = 2
	TUNNEL_TYPE_6TO4                = 11
	TUNNEL_TYPE_ISATAP              = 13
	TUNNEL_TYPE_TEREDO              = 14
	TUNNEL_TYPE_IPHTTPS             = 15
)

const MAX_DNS_SUFFIX_STRING_LENGTH = 256

const ANY_SIZE = 1

const MAX_ADAPTER_NAME = 128

const MAX_HOSTNAME_LEN = 128
const MAX_DOMAIN_NAME_LEN = 128
const MAX_SCOPE_ID_LEN = 256

type TCPIP_OWNER_MODULE_INFO_CLASS int32

const (
	TCPIP_OWNER_MODULE_INFO_BASIC TCPIP_OWNER_MODULE_INFO_CLASS = iota
)

const TCPIP_OWNING_MODULE_SIZE = 16
const MAX_DEFAULTCHAR = 2
const MAX_LEADBYTES = 12

const (
	FOREGROUND_BLUE            = 0x0001
	FOREGROUND_GREEN           = 0x0002
	FOREGROUND_RED             = 0x0004
	FOREGROUND_INTENSITY       = 0x0008
	BACKGROUND_BLUE            = 0x0010
	BACKGROUND_GREEN           = 0x0020
	BACKGROUND_RED             = 0x0040
	BACKGROUND_INTENSITY       = 0x0080
	COMMON_LVB_LEADING_BYTE    = 0x0100
	COMMON_LVB_TRAILING_BYTE   = 0x0200
	COMMON_LVB_GRID_HORIZONTAL = 0x0400
	COMMON_LVB_GRID_LVERTICAL  = 0x0800
	COMMON_LVB_GRID_RVERTICAL  = 0x1000
	COMMON_LVB_REVERSE_VIDEO   = 0x4000
	COMMON_LVB_UNDERSCORE      = 0x8000
)

type TCP_ESTATS_TYPE int32

const (
	TcpConnectionEstatsSynOpts TCP_ESTATS_TYPE = iota
	TcpConnectionEstatsData
	TcpConnectionEstatsSndCong
	TcpConnectionEstatsPath
	TcpConnectionEstatsSendBuff
	TcpConnectionEstatsRec
	TcpConnectionEstatsObsRec
	TcpConnectionEstatsBandwidth
	TcpConnectionEstatsFineRtt
	TcpConnectionEstatsMaximum
)

type TCP_CONNECTION_OFFLOAD_STATE int32

const (
	TcpConnectionOffloadStateInHost TCP_CONNECTION_OFFLOAD_STATE = iota
	TcpConnectionOffloadStateOffloading
	TcpConnectionOffloadStateOffloaded
	TcpConnectionOffloadStateUploading
	TcpConnectionOffloadStateMax
)

type TCP_RTO_ALGORITHM int32

const (
	TcpRtoAlgorithmOther TCP_RTO_ALGORITHM = iota
	TcpRtoAlgorithmConstant
	TcpRtoAlgorithmRsre
	TcpRtoAlgorithmVanj

	MIB_TCP_RTO_OTHER    = 1
	MIB_TCP_RTO_CONSTANT = 2
	MIB_TCP_RTO_RSRE     = 3
	MIB_TCP_RTO_VANJ     = 4
)

type NET_ADDRESS_FORMAT int32

const (
	NET_ADDRESS_FORMAT_UNSPECIFIED NET_ADDRESS_FORMAT = iota
	NET_ADDRESS_DNS_NAME
	NET_ADDRESS_IPV4
	NET_ADDRESS_IPV6
)

const DNS_MAX_NAME_BUFFER_LENGTH = 256

type NDIS_MEDIUM int32

const (
	NdisMedium802_3 NDIS_MEDIUM = iota
	NdisMedium802_5
	NdisMediumFddi
	NdisMediumWan
	NdisMediumLocalTalk
	NdisMediumDix
	NdisMediumArcnetRaw
	NdisMediumArcnet878_2
	NdisMediumAtm
	NdisMediumWirelessWan
	NdisMediumIrda
	NdisMediumBpc
	NdisMediumCoWan
	NdisMedium1394
	NdisMediumInfiniBand
	NdisMediumTunnel
	NdisMediumNative802_11
	NdisMediumLoopback
	NdisMediumWiMAX
	NdisMediumIP
	NdisMediumMax
)

type NDIS_PHYSICAL_MEDIUM int32

const (
	NdisPhysicalMediumUnspecified NDIS_PHYSICAL_MEDIUM = iota
	NdisPhysicalMediumWirelessLan
	NdisPhysicalMediumCableModem
	NdisPhysicalMediumPhoneLine
	NdisPhysicalMediumPowerLine
	NdisPhysicalMediumDSL
	NdisPhysicalMediumFibreChannel
	NdisPhysicalMedium1394
	NdisPhysicalMediumWirelessWan
	NdisPhysicalMediumNative802_11
	NdisPhysicalMediumBluetooth
	NdisPhysicalMediumInfiniband
	NdisPhysicalMediumWiMax
	NdisPhysicalMediumUWB
	NdisPhysicalMedium802_3
	NdisPhysicalMedium802_5
	NdisPhysicalMediumIrda
	NdisPhysicalMediumWiredWAN
	NdisPhysicalMediumWiredCoWan
	NdisPhysicalMediumOther
	NdisPhysicalMediumMax
)

type NET_IF_ACCESS_TYPE int32

const (
	NET_IF_ACCESS_LOOPBACK             NET_IF_ACCESS_TYPE = 1
	NET_IF_ACCESS_BROADCAST                               = 2
	NET_IF_ACCESS_POINT_TO_POINT                          = 3
	NET_IF_ACCESS_POINT_TO_MULTI_POINT                    = 4
	NET_IF_ACCESS_MAXIMUM                                 = 5
)

type NET_IF_ADMIN_STATUS int32

const (
	NET_IF_ADMIN_STATUS_UP      NET_IF_ADMIN_STATUS = 1
	NET_IF_ADMIN_STATUS_DOWN                        = 2
	NET_IF_ADMIN_STATUS_TESTING                     = 3
)

type NET_IF_DIRECTION_TYPE int32

const (
	NET_IF_DIRECTION_SENDRECEIVE NET_IF_DIRECTION_TYPE = iota
	NET_IF_DIRECTION_SENDONLY
	NET_IF_DIRECTION_RECEIVEONLY
	NET_IF_DIRECTION_MAXIMUM
)

type NET_IF_MEDIA_CONNECT_STATE int32

const (
	MediaConnectStateUnknown NET_IF_MEDIA_CONNECT_STATE = iota
	MediaConnectStateConnected
	MediaConnectStateDisconnected
)

const IF_MAX_STRING_SIZE = 256
const IF_MAX_PHYS_ADDRESS_LENGTH = 32

type PFADDRESSTYPE int32

const (
	PF_IPV4 PFADDRESSTYPE = iota
	PF_IPV6
)

type PFFORWARD_ACTION int32

const (
	PF_ACTION_FORWARD PFFORWARD_ACTION = iota
	PF_ACTION_DROP
)

const MAX_JOYSTICKOEMVXDNAME = 260

const MIXER_SHORT_NAME_CHARS = 16
const MIXER_LONG_NAME_CHARS = 64

const NUM_POINTS = 3

type ASSOC_FILTER int32

const (
	ASSOC_FILTER_NONE        ASSOC_FILTER = 0
	ASSOC_FILTER_RECOMMENDED ASSOC_FILTER = 0x1
)

type FOLDERVIEWMODE int32

const (
	FVM_AUTO       FOLDERVIEWMODE = -1
	FVM_FIRST                     = 1
	FVM_ICON                      = 1
	FVM_SMALLICON                 = 2
	FVM_LIST                      = 3
	FVM_DETAILS                   = 4
	FVM_THUMBNAIL                 = 5
	FVM_TILE                      = 6
	FVM_THUMBSTRIP                = 7
	FVM_CONTENT                   = 8
	FVM_LAST                      = 8
)
