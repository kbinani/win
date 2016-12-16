package internal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Type struct {
	Name               string
	goName             string
	isConst            bool
	dereferencedGoName string
	isPtr              bool
	funcDecl           string
}

var (
	reg *regexp.Regexp
)

func init() {
	reg = regexp.MustCompile(`{"Name":"(.*)"}`)
}

func NewType(s string) *Type {
	ret := new(Type)
	ret.init(s)
	return ret
}

func (t *Type) UnmarshalJSON(data []byte) error {
	m := reg.FindSubmatch(data)
	if len(m) >= 2 {
		t.init(string(m[1]))
	}
	return nil
}

func (ret *Type) init(s string) {
	ret.Name = s

	goName, isConst, funcDecl, isPtr, err := ret.getGoName()
	if err != nil {
		ret.goName = ""
	} else {
		ret.goName = goName
	}

	ret.isConst = isConst
	ret.funcDecl = funcDecl
	ret.isPtr = isPtr
	ret.dereferencedGoName = getDereferencedGoName(goName)
}

func (t Type) IsPtr() bool {
	return t.isPtr
}

func (t Type) IsConst() bool {
	return t.isConst
}

func (t Type) FuncDecl() string {
	return t.funcDecl
}

func (t Type) GoName() string {
	return t.goName
}

func (t Type) DereferencedGoName() string {
	return t.dereferencedGoName
}

func getDereferencedGoName(name string) string {
	for strings.HasPrefix(name, "*") {
		name = strings.TrimPrefix(name, "*")
	}
	return name
}

func (tt Type) getGoName() (goName string, isConstType bool, funcDeclStr string, isPtrType bool, e error) {
	tokens := strings.Split(tt.Name, " ")
	t := ""
	isConst := false
	for _, token := range tokens {
		if token == "CONST" || token == "const" {
			isConst = true
			continue
		} else if token == "UNALIGNED" {
			continue
		}
		t = t + token + " "
	}
	t = strings.Trim(t, " ")
	t = ReplaceString(t, "  ", " ")
	t = ReplaceString(t, " *", "*")
	t = ReplaceString(t, "* ", "*")

	isPtr := false
	funcDecl := ""

	table := map[string]string{
		// Basic types
		"WINBOOL":                     "bool",
		"BOOL":                        "bool",
		"short":                       "int16",
		"char":                        "byte",
		"__LONG32":                    "int32",
		"WNDCLASSEXW":                 "WNDCLASSEX",
		"DEVMODEW":                    "DEVMODE",
		"WNDCLASSW":                   "WNDCLASS",
		"MENUITEMINFOW":               "MENUITEMINFO",
		"IMEPROW":                     "IMEPRO",
		"DLGTEMPLATEW":                "DLGTEMPLATE",
		"MSGBOXPARAMSW":               "MSGBOXPARAMS",
		"DISPLAY_DEVICEW":             "DISPLAY_DEVICE",
		"float":                       "float32",
		"double":                      "float64",
		"LOGFONTW":                    "LOGFONT",
		"EXPLICIT_ACCESS_W":           "EXPLICIT_ACCESS",
		"TRUSTEE_W":                   "TRUSTEE",
		"OBJECTS_AND_NAME_W":          "OBJECTS_AND_NAME",
		"STARTUPINFOW":                "STARTUPINFO",
		"PDH_COUNTER_PATH_ELEMENTS_W": "PDH_COUNTER_PATH_ELEMENTS",
		"CHOOSECOLORW":                "CHOOSECOLOR",
		"TEXTMETRICW":                 "TEXTMETRIC",
		"PARSEDURLW":                  "PARSEDURL",
		"int":                         "int32",
		"CHOOSEFONTW":                 "CHOOSEFONT",
		"PROPSHEETHEADERW_V2":         "PROPSHEETHEADER_V2",
		"PROPSHEETHEADERW":            "PROPSHEETHEADER",
		"MRUINFOW":                    "MRUINFO",
		"OUTLINETEXTMETRICW":          "OUTLINETEXTMETRIC",
		"POLYTEXTW":                   "POLYTEXT",
		"DOCINFOW":                    "DOCINFO",
		"unsigned int":                "uint32",
		"signed char":                 "int8",
		"IMEMENUITEMINFOW":            "IMEMENUITEMINFO",
		"STYLEBUFW":                   "STYLEBUF",
		"ADDRINFOEXW":                 "ADDRINFOEX",
		"WSAPROTOCOL_INFOW":           "WSAPROTOCOL_INFO",
		"WS_socklen_t":                "Socklen_t",
		"WS_u_long":                   "ULONG",
		"WS_u_short":                  "uint16",
		"WSAQUERYSETW":                "WSAQUERYSET",
		"time_t":                      "Time_t",
		"CPINFOEXW":                   "CPINFOEX",
		"sockaddr_in6":                "SOCKADDR_IN6_LH",
		"JOYCAPSW":                    "JOYCAPS",
		"MIDIOUTCAPSW":                "MIDIOUTCAPS",
		"MIXERCAPSW":                  "MIXERCAPS",
		"MIXERLINECONTROLSW":          "MIXERLINECONTROLS",
		"MIXERCONTROLW":               "MIXERCONTROL",
		"MIXERLINEW":                  "MIXERLINE",
		"WAVEINCAPSW":                 "WAVEINCAPS",
		"WAVEOUTCAPSW":                "WAVEOUTCAPS",
		"BROWSEINFOW":                 "BROWSEINFO",
		"SHFILEOPSTRUCTW":             "SHFILEOPSTRUCT",
		"SHFILEINFOW":                 "SHFILEINFO",
		"NOTIFYICONDATAW":             "NOTIFYICONDATA",
		"handle_t":                    "Handle_t",
		"unsigned char":               "byte",
		"RPC_PROTSEQ_VECTORW":         "RPC_PROTSEQ_VECTOR",
		"DECLSPEC_NORETURN":           "void", // __declspec(noreturn)
		"VOID":                        "void",
		"void":                        "void",
		"ASN1decoding_t":              "uintptr", // void*
		"ASN1encoding_t":              "uintptr", // void*
		"ASN1module_t":                "uintptr", // void*
		"HCRYPTASN1MODULE":            "DWORD",
		"HLRUCACHE":                   "uintptr", // void*
	}

	goType, ok := table[t]
	if ok {
		return goType, isConst, funcDecl, isPtr, nil
	}

	tablePtr := map[string]string{
		// Pointer types
		"PBSMINFO":              "*BSMINFO",
		"LPDWORD":               "*uint32",
		"LONG_PTR":              "uintptr",
		"VOID*":                 "uintptr",
		"void*":                 "uintptr",
		"BYTE*":                 "*byte",
		"LPCWSTR":               "LPCWSTR",
		"PCWSTR":                "LPCWSTR",
		"LPWNDCLASSEXW":         "*WNDCLASSEX",
		"WNDCLASSEXW*":          "*WNDCLASSEX",
		"UINTPTR":               "uintptr",
		"LPVOID":                "LPVOID",
		"PVOID":                 "uintptr",
		"UINT*":                 "*UINT",
		"PUINT":                 "*uint32",
		"INT*":                  "*int32",
		"LPINT":                 "*int32",
		"LPWORD":                "*uint16",
		"DWORD_PTR":             "*uint32",
		"PLONG":                 "*int32",
		"PULONG":                "*uint32",
		"ULONG_PTR":             "*uint32",
		"PBYTE":                 "*byte",
		"LPBYTE":                "*byte",
		"PBOOL":                 "*BOOL",
		"LPSTR":                 "LPSTR",
		"WINBOOL*":              "*BOOL",
		"PDWORD_PTR":            "*uintptr",
		"PUINT_PTR":             "*uintptr",
		"INT_PTR":               "INT_PTR",
		"UINT_PTR":              "*uint32",
		"PFLOAT":                "*float32",
		"LPCH":                  "*CHAR",
		"PCONVCONTEXT":          "*CONVCONTEXT",
		"WINDOWPLACEMENT*":      "*WINDOWPLACEMENT",
		"PSECURITY_INFORMATION": "*SECURITY_INFORMATION",
		"WNDCLASSW*":            "*WNDCLASS",
		"LPWNDCLASSW":           "*WNDCLASS",
		"DWORD*":                "*uint32",
		"PCOMBOBOXINFO":         "*COMBOBOXINFO",
		"LPMENUITEMINFOW":       "*MENUITEMINFO",
		"PWINDOWINFO":           "*WINDOWINFO",
		"LPIMEPROW":             "*IMEPRO",
		"LPWSTR":                "LPWSTR",
		"PICONINFO":             "*ICONINFO",
		"PSECURITY_DESCRIPTOR":  "PSECURITY_DESCRIPTOR",
		"PTITLEBARINFO":         "*TITLEBARINFO",
		"PACL":                  "*ACL",
		"PSECURITY_QUALITY_OF_SERVICE":       "*SECURITY_QUALITY_OF_SERVICE",
		"MSGBOXPARAMSW*":                     "*MSGBOXPARAMS",
		"PCONVINFO":                          "*CONVINFO",
		"PDISPLAY_DEVICEW":                   "*DISPLAY_DEVICE",
		"PSCROLLBARINFO":                     "*SCROLLBARINFO",
		"PRAWINPUTDEVICELIST":                "*RAWINPUTDEVICELIST",
		"PGESTUREINFO":                       "*GESTUREINFO",
		"PGUITHREADINFO":                     "*GUITHREADINFO",
		"PRAWINPUTDEVICE":                    "*RAWINPUTDEVICE",
		"PLASTINPUTINFO":                     "*LASTINPUTINFO",
		"PFLASHWINFO":                        "*FLASHWINFO",
		"PCURSORINFO":                        "*CURSORINFO",
		"PRAWINPUT*":                         "uintptr",
		"PTOUCHINPUT":                        "*TOUCHINPUT",
		"PALTTABINFO":                        "*ALTTABINFO",
		"PGESTURECONFIG":                     "*GESTURECONFIG",
		"PMENUBARINFO":                       "*MENUBARINFO",
		"PRAWINPUT":                          "*RAWINPUT",
		"MENUTEMPLATEW*":                     "uintptr",
		"GLvoid*":                            "uintptr",
		"GLvoid**":                           "uintptr",
		"PLOCALMANAGEDAPPLICATION":           "*LOCALMANAGEDAPPLICATION",
		"PHKEY":                              "*HKEY",
		"PHANDLE":                            "*HANDLE",
		"PSECURITY_DESCRIPTOR_CONTROL":       "*SECURITY_DESCRIPTOR_CONTROL",
		"PPERF_COUNTERSET_INSTANCE":          "*PERF_COUNTERSET_INSTANCE",
		"PCREDENTIALW":                       "*CREDENTIAL",
		"PCREDENTIALW**":                     "uintptr",
		"GUID**":                             "uintptr",
		"LOGFONTW*":                          "*LOGFONT",
		"PGENERIC_MAPPING":                   "*GENERIC_MAPPING",
		"PPRIVILEGE_SET":                     "*PRIVILEGE_SET",
		"POBJECT_TYPE_LIST":                  "*OBJECT_TYPE_LIST",
		"PENCRYPTION_CERTIFICATE_LIST":       "*ENCRYPTION_CERTIFICATE_LIST",
		"PTOKEN_GROUPS":                      "*TOKEN_GROUPS",
		"PDWORD":                             "*DWORD",
		"PENCRYPTION_CERTIFICATE":            "PENCRYPTION_CERTIFICATE",
		"PWCHAR":                             "PWCHAR",
		"PTOKEN_PRIVILEGES":                  "*TOKEN_PRIVILEGES",
		"PSID_IDENTIFIER_AUTHORITY":          "*SID_IDENTIFIER_AUTHORITY",
		"PLUID":                              "*LUID",
		"PAUDIT_POLICY_INFORMATION":          "PAUDIT_POLICY_INFORMATION",
		"PPOLICY_AUDIT_SID_ARRAY":            "PPOLICY_AUDIT_SID_ARRAY",
		"PPOLICY_AUDIT_EVENT_TYPE":           "PPOLICY_AUDIT_EVENT_TYPE",
		"PEXPLICIT_ACCESS_W":                 "*EXPLICIT_ACCESS",
		"PTRUSTEE_W":                         "*TRUSTEE",
		"POBJECTS_AND_NAME_W":                "*OBJECTS_AND_NAME",
		"POBJECTS_AND_SID":                   "*OBJECTS_AND_SID",
		"LPSTARTUPINFOW":                     "*STARTUPINFO",
		"LPAUXCAPSW":                         "*AUXCAPS",
		"LPMIDIINCAPSW":                      "*MIDIINCAPS",
		"PPDH_FMT_COUNTERVALUE":              "*PDH_FMT_COUNTERVALUE",
		"PPDH_RAW_COUNTER":                   "*PDH_RAW_COUNTER",
		"PPDH_COUNTER_INFO_W":                "*PDH_COUNTER_INFO",
		"PPDH_COUNTER_PATH_ELEMENTS_W":       "*PDH_COUNTER_PATH_ELEMENTS",
		"PDH_COUNTER_PATH_ELEMENTS_W*":       "*PDH_COUNTER_PATH_ELEMENTS",
		"void**":                             "uintptr",
		"PCRYPT_INTEGER_BLOB":                "PCRYPT_INTEGER_BLOB",
		"PCRYPT_UINT_BLOB":                   "PCRYPT_UINT_BLOB",
		"PCRYPT_OBJID_BLOB":                  "PCRYPT_OBJID_BLOB",
		"PCERT_NAME_BLOB":                    "PCERT_NAME_BLOB",
		"PCERT_RDN_VALUE_BLOB":               "PCERT_RDN_VALUE_BLOB",
		"PCERT_BLOB":                         "PCERT_BLOB",
		"PCRL_BLOB":                          "PCRL_BLOB",
		"PDATA_BLOB":                         "PDATA_BLOB",
		"PCRYPT_DATA_BLOB":                   "PCRYPT_DATA_BLOB",
		"PCRYPT_HASH_BLOB":                   "PCRYPT_HASH_BLOB",
		"PCRYPT_DIGEST_BLOB":                 "PCRYPT_DIGEST_BLOB",
		"PCRYPT_DER_BLOB":                    "PCRYPT_DER_BLOB",
		"PCRYPT_ATTR_BLOB":                   "PCRYPT_ATTR_BLOB",
		"BSTR":                               "BSTR",
		"PRPC_MESSAGE":                       "PRPC_MESSAGE",
		"NDR_SCONTEXT":                       "NDR_SCONTEXT",
		"PMIDL_STUB_MESSAGE":                 "PMIDL_STUB_MESSAGE",
		"PCERT_EXTENSION":                    "PCERT_EXTENSION",
		"LPCHOOSECOLORW":                     "*CHOOSECOLOR",
		"LPMONIKER":                          "LPMONIKER",
		"LPUNKNOWN":                          "LPUNKNOWN",
		"PULONG64":                           "PULONG64",
		"struct GdiplusStartupInput*":        "*GdiplusStartupInput",
		"struct GdiplusStartupOutput*":       "*GdiplusStartupOutput",
		"PCRL_INFO":                          "PCRL_INFO",
		"PCCRL_CONTEXT":                      "PCCRL_CONTEXT",
		"REFIID":                             "REFIID",
		"REFGUID":                            "REFGUID",
		"TEXTMETRICW*":                       "*TEXTMETRIC",
		"unsigned char*":                     "*byte",
		"PARSEDURLW*":                        "*PARSEDURL",
		"PHUSKEY":                            "PHUSKEY",
		"struct wow_handlers16*":             "uintptr",
		"struct wow_handlers32*":             "uintptr",
		"MRUINFOW*":                          "*MRUINFO",
		"PINT_PTR":                           "PINT_PTR",
		"LPSTREAM":                           "LPSTREAM",
		"VOID**":                             "uintptr",
		"PUSHORT":                            "PUSHORT",
		"PCHAR":                              "PCHAR",
		"LPCHOOSEFONTW":                      "LPCHOOSEFONT",
		"LPFINDREPLACEW":                     "LPFINDREPLACE",
		"LPOPENFILENAMEW":                    "LPOPENFILENAME",
		"LPPAGESETUPDLGW":                    "LPPAGESETUPDLG",
		"LPPRINTDLGEXW":                      "LPPRINTDLGEX",
		"LPPRINTDLGW":                        "LPPRINTDLG",
		"LPLOGCOLORSPACEW":                   "LPLOGCOLORSPACE",
		"DEVMODEW*":                          "*DEVMODE",
		"ENUMLOGFONTEXDVW*":                  "*ENUMLOGFONTEXDV",
		"LPDEVMODEW":                         "LPDEVMODE",
		"PWSTR":                              "PWSTR",
		"PENG_TIME_FIELDS":                   "PENG_TIME_FIELDS",
		"PFD_GLYPHATTR":                      "PFD_GLYPHATTR",
		"PTRIVERTEX":                         "PTRIVERTEX",
		"LPGCP_RESULTSW":                     "LPGCP_RESULTS",
		"LPLOGFONTW":                         "LPLOGFONT",
		"LPOUTLINETEXTMETRICW":               "LPOUTLINETEXTMETRIC",
		"LPRASTERIZER_STATUS":                "LPRASTERIZER_STATUS",
		"LPTEXTMETRICW":                      "LPTEXTMETRIC",
		"PRECTFX":                            "PRECTFX",
		"POLYTEXTW*":                         "*POLYTEXT",
		"PGLYPHPOS":                          "PGLYPHPOS",
		"DOCINFOW*":                          "*DOCINFO",
		"struct realization_info*":           "uintptr",
		"PPROCESS_MEMORY_COUNTERS":           "PPROCESS_MEMORY_COUNTERS",
		"PPSAPI_WS_WATCH_INFORMATION_EX":     "PPSAPI_WS_WATCH_INFORMATION_EX",
		"signed char*":                       "*int8",
		"float*":                             "*float32",
		"double*":                            "*float64",
		"LPOLESTR":                           "LPOLESTR",
		"LPSAFEARRAY":                        "LPSAFEARRAY",
		"PPSAPI_WS_WATCH_INFORMATION":        "PPSAPI_WS_WATCH_INFORMATION",
		"PPERFORMACE_INFORMATION":            "PPERFORMACE_INFORMATION",
		"LPIMEMENUITEMINFOW":                 "LPIMEMENUITEMINFO",
		"LPSTYLEBUFW":                        "LPSTYLEBUF",
		"ADDRINFOEXW*":                       "*ADDRINFOEX",
		"ADDRINFOEXW**":                      "**ADDRINFOEX",
		"PADDRINFOW":                         "PADDRINFO",
		"LPWSAPROTOCOL_INFOW":                "LPWSAPROTOCOL_INFO",
		"PADDRINFOW*":                        "*PADDRINFO",
		"ADDRINFOW*":                         "*ADDRINFO",
		"struct timeval*":                    "*Timeval",
		"LPWSABUF":                           "LPWSABUF",
		"struct WS_sockaddr*":                "*Sockaddr",
		"LPWSANAMESPACE_INFOW":               "LPWSANAMESPACE_INFO",
		"LPWSAOVERLAPPED":                    "LPWSAOVERLAPPED",
		"LPWSASERVICECLASSINFOW":             "LPWSASERVICECLASSINFO",
		"WS_u_long*":                         "*ULONG",
		"WS_u_short*":                        "*uint16",
		"LPWSAQUERYSETW":                     "LPWSAQUERYSET",
		"char*":                              "*CHAR",
		"PMIB_IPFORWARDROW":                  "PMIB_IPFORWARDROW",
		"PMIB_IPNETROW":                      "PMIB_IPNETROW",
		"MIB_IPSTATS_LH":                     "MIB_IPSTATS_LH",
		"PIP_ADAPTER_ORDER_MAP":              "PIP_ADAPTER_ORDER_MAP",
		"PMIB_IPSTATS":                       "PMIB_IPSTATS",
		"PMIB_TCPROW":                        "PMIB_TCPROW",
		"PIP_ADAPTER_ADDRESSES":              "PIP_ADAPTER_ADDRESSES",
		"PIP_ADAPTER_UNICAST_ADDRESS_LH":     "PIP_ADAPTER_UNICAST_ADDRESS_LH",
		"PIP_ADAPTER_ANYCAST_ADDRESS_XP":     "PIP_ADAPTER_ANYCAST_ADDRESS_XP",
		"struct sockaddr*":                   "*Sockaddr",
		"PMIB_ICMP":                          "PMIB_ICMP",
		"PIP_ADAPTER_INFO":                   "PIP_ADAPTER_INFO",
		"PMIB_ICMP_EX":                       "PMIB_ICMP_EX",
		"PMIB_IFROW":                         "PMIB_IFROW",
		"PMIB_IFTABLE":                       "PMIB_IFTABLE",
		"PIP_INTERFACE_INFO":                 "PIP_INTERFACE_INFO",
		"PMIB_TCP6ROW_OWNER_MODULE":          "PMIB_TCP6ROW_OWNER_MODULE",
		"PMIB_IPADDRTABLE":                   "PMIB_IPADDRTABLE",
		"PMIB_IPFORWARDTABLE":                "PMIB_IPFORWARDTABLE",
		"PFIXED_INFO":                        "PFIXED_INFO",
		"PMIB_IPNETTABLE":                    "PMIB_IPNETTABLE",
		"PMIB_TCPROW_OWNER_MODULE":           "PMIB_TCPROW_OWNER_MODULE",
		"PMIB_UDP6ROW_OWNER_MODULE":          "PMIB_UDP6ROW_OWNER_MODULE",
		"PMIB_UDPROW_OWNER_MODULE":           "PMIB_UDPROW_OWNER_MODULE",
		"PIP_PER_ADAPTER_INFO_W2KSP1":        "PIP_PER_ADAPTER_INFO_W2KSP1",
		"PIP_PER_ADAPTER_INFO":               "PIP_PER_ADAPTER_INFO",
		"PSRWLOCK":                           "PSRWLOCK",
		"LPCPINFOEXW":                        "LPCPINFOEX",
		"LPCPINFO":                           "LPCPINFO",
		"PCONSOLE_SCREEN_BUFFER_INFO":        "PCONSOLE_SCREEN_BUFFER_INFO",
		"PCONSOLE_SCREEN_BUFFER_INFOEX":      "PCONSOLE_SCREEN_BUFFER_INFOEX",
		"PMIB_TCP6ROW":                       "PMIB_TCP6ROW",
		"PUCHAR":                             "PUCHAR",
		"PMIB_TCP6TABLE":                     "PMIB_TCP6TABLE",
		"PMIB_TCP6TABLE2":                    "PMIB_TCP6TABLE2",
		"PMIB_TCP6ROW2":                      "PMIB_TCP6ROW2",
		"PMIB_TCPSTATS":                      "PMIB_TCPSTATS",
		"PMIB_TCPTABLE":                      "PMIB_TCPTABLE",
		"PMIB_TCPTABLE2":                     "PMIB_TCPTABLE2",
		"PMIB_UDP6TABLE":                     "PMIB_UDP6TABLE",
		"PMIB_UDPSTATS":                      "PMIB_UDPSTATS",
		"PMIB_UDPTABLE":                      "PMIB_UDPTABLE",
		"PMIB_UDPROW":                        "PMIB_UDPROW",
		"PIP_UNIDIRECTIONAL_ADAPTER_ADDRESS": "PIP_UNIDIRECTIONAL_ADAPTER_ADDRESS",
		"PIP_OPTION_INFORMATION":             "PIP_OPTION_INFORMATION",
		"PIO_STATUS_BLOCK":                   "PIO_STATUS_BLOCK",
		"PSOCKADDR_IN6_LH":                   "PSOCKADDR_IN6_LH",
		"struct sockaddr_in6*":               "*SOCKADDR_IN6_LH",
		"PNET_ADDRESS_INFO":                  "PNET_ADDRESS_INFO",
		"PIP_ADAPTER_INDEX_MAP":              "PIP_ADAPTER_INDEX_MAP",
		"PSOCKADDR_IN6":                      "PSOCKADDR_IN6",
		"LPJOYCAPSW":                         "LPJOYCAPS",
		"LPJOYINFO":                          "LPJOYINFO",
		"LPMIDIOUTCAPSW":                     "LPMIDIOUTCAPS",
		"LPMMIOINFO":                         "LPMMIOINFO",
		"HPSTR":                              "HPSTR",
		"LPMMCKINFO":                         "LPMMCKINFO",
		"char _huge*":                        "HPSTR",
		"LPTIMECAPS":                         "LPTIMECAPS",
		"LPWAVEHDR":                          "LPWAVEHDR",
		"LPMIXERCAPSW":                       "LPMIXERCAPS",
		"LPMIXERLINECONTROLSW":               "LPMIXERLINECONTROLS",
		"LPMIXERCONTROLW":                    "LPMIXERCONTROL",
		"LPMIXERLINEW":                       "LPMIXERLINE",
		"LPWAVEINCAPSW":                      "LPWAVEINCAPS",
		"LPWAVEFORMATEX":                     "LPWAVEFORMATEX",
		"LPWAVEOUTCAPSW":                     "LPWAVEOUTCAPS",
		"LPDATAOBJECT":                       "LPDATAOBJECT",
		"LPITEMIDLIST":                       "LPITEMIDLIST",
		"LPCABINETSTATE":                     "LPCABINETSTATE",
		"PAPPBARDATA":                        "PAPPBARDATA",
		"LPBROWSEINFOW":                      "LPBROWSEINFO",
		"PCUITEMID_CHILD_ARRAY":              "PCUITEMID_CHILD_ARRAY",
		"LPCSFV":                             "LPCSFV",
		"PCIDLIST_ABSOLUTE_ARRAY":            "PCIDLIST_ABSOLUTE_ARRAY",
		"LPDROPSOURCE":                       "LPDROPSOURCE",
		"LPSHFILEOPSTRUCTW":                  "LPSHFILEOPSTRUCT",
		"LPSHELLFOLDER":                      "LPSHELLFOLDER",
		"REFKNOWNFOLDERID":                   "REFKNOWNFOLDERID",
		"SHFILEINFOW*":                       "*SHFILEINFO",
		"PCERT_INFO":                         "PCERT_INFO",
		"PCTL_INFO":                          "PCTL_INFO",
		"PCTL_ENTRY":                         "PCTL_ENTRY",
		"PCRYPT_ATTRIBUTE":                   "PCRYPT_ATTRIBUTE",
		"HCRYPTMSG":                          "HCRYPTMSG",
		"PIDLIST_ABSOLUTE":                   "PIDLIST_ABSOLUTE",
		"LPMALLOC":                           "LPMALLOC",
		"LPSHELLSTATE":                       "LPSHELLSTATE",
		"LPSHELLFLAGSTATE":                   "LPSHELLFLAGSTATE",
		"REFFMTID":                           "REFFMTID",
		"LPVERSIONEDSTREAM":                  "LPVERSIONEDSTREAM",
		"LPSHQUERYRBINFO":                    "LPSHQUERYRBINFO",
		"PNOTIFYICONDATAW":                   "PNOTIFYICONDATA",
		"PROPVARIANT":                        "PROPVARIANT",
		"REFPROPVARIANT":                     "REFPROPVARIANT",
		"LPMARSHAL":                          "LPMARSHAL",
		"RPC_AUTHZ_HANDLE":                   "RPC_AUTHZ_HANDLE",
		"LPMALLOCSPY":                        "LPMALLOCSPY",
		"LPMESSAGEFILTER":                    "LPMESSAGEFILTER",
		"LPBC":                               "LPBC",
		"LPRUNNINGOBJECTTABLE":               "LPRUNNINGOBJECTTABLE",
		"LPOLESTREAM":                        "LPOLESTREAM",
		"LPSTORAGE":                          "LPSTORAGE",
		"LPOLECLIENTSITE":                    "LPOLECLIENTSITE",
		"LPOLEMENUGROUPWIDTHS":               "LPOLEMENUGROUPWIDTHS",
		"LPOLEOBJECT":                        "LPOLEOBJECT",
		"LPENUMOLEVERB":                      "LPENUMOLEVERB",
		"LPPERSISTSTORAGE":                   "LPPERSISTSTORAGE",
		"LPOLEINPLACEACTIVEOBJECT":           "LPOLEINPLACEACTIVEOBJECT",
		"LPOLEINPLACEFRAME":                  "LPOLEINPLACEFRAME",
		"LPOLEINPLACEFRAMEINFO":              "LPOLEINPLACEFRAMEINFO",
		"LPDROPTARGET":                       "LPDROPTARGET",
		"HMETAFILEPICT":                      "HMETAFILEPICT",
		"SNB":                                "SNB",
		"LPRPCSTUBBUFFER":                    "*IRpcStubBuffer",
		"LPRPCPROXYBUFFER":                   "*IRpcProxyBuffer",
		"LPRPCPROXYBUFFER*":                  "**IRpcProxyBuffer",
		"LPTYPEINFO":                         "LPTYPEINFO",
		"PRPC_ASYNC_STATE":                   "PRPC_ASYNC_STATE",
		"char**":                             "**CHAR",
		"handle_t*":                          "*Handle_t",
		"NDR_CCONTEXT":                       "NDR_CCONTEXT",
		"unsigned int*":                      "*uint32",
		"PFULL_PTR_XLAT_TABLES":              "PFULL_PTR_XLAT_TABLES",
		"struct IRpcChannelBuffer*":          "*IRpcChannelBuffer",
		"struct IRpcStubBuffer*":             "*IRpcStubBuffer",
		"LPRPCCHANNELBUFFER":                 "*IRpcChannelBuffer",
		"RPC_WSTR":                           "RPC_WSTR",
		"RPC_IF_HANDLE":                      "RPC_IF_HANDLE",
		"RPC_EP_INQ_HANDLE":                  "RPC_EP_INQ_HANDLE",
		"I_RPC_HANDLE":                       "I_RPC_HANDLE",
		"RPC_PROTSEQ_VECTORW**":              "**RPC_PROTSEQ_VECTOR",
		"RPC_PROTSEQ_VECTORW*":               "*RPC_PROTSEQ_VECTOR",
		"RPC_MGR_EPV*":                       "uintptr",
		"PRPC_POLICY":                        "PRPC_POLICY",
		"twr_t*":                             "uintptr", // twr_t not found in MSDN.
		"twr_t**":                            "*uintptr",
		"PRPC_SYNTAX_IDENTIFIER":             "PRPC_SYNTAX_IDENTIFIER",
		"PCERT_PUBLIC_KEY_INFO":              "PCERT_PUBLIC_KEY_INFO",
		"PCERT_CHAIN_ENGINE_CONFIG":          "PCERT_CHAIN_ENGINE_CONFIG",
		"PCERT_EXTENSIONS":                   "PCERT_EXTENSIONS",
		"PCRYPT_ALGORITHM_IDENTIFIER":        "PCRYPT_ALGORITHM_IDENTIFIER",
		"PCRYPT_KEY_PROV_INFO":               "PCRYPT_KEY_PROV_INFO",
		"PSYSTEMTIME":                        "PSYSTEMTIME",
		"PCRYPT_KEY_PROV_PARAM":              "PCRYPT_KEY_PROV_PARAM",
		"PCERT_SIMPLE_CHAIN":                 "PCERT_SIMPLE_CHAIN",
		"PCERT_CHAIN_ELEMENT":                "PCERT_CHAIN_ELEMENT",
		"PCERT_TRUST_LIST_INFO":              "PCERT_TRUST_LIST_INFO",
		"PCERT_ENHKEY_USAGE":                 "PCERT_ENHKEY_USAGE",
		"HCRYPTPROV_OR_NCRYPT_KEY_HANDLE":    "HCRYPTPROV_OR_NCRYPT_KEY_HANDLE",
		"PCERT_PHYSICAL_STORE_INFO":          "PCERT_PHYSICAL_STORE_INFO",
		"PCERT_SYSTEM_STORE_INFO":            "PCERT_SYSTEM_STORE_INFO",
		"PCERT_NAME_INFO":                    "PCERT_NAME_INFO",
		"PCERT_RDN_ATTR":                     "PCERT_RDN_ATTR",
		"PCERT_CHAIN_PARA":                   "PCERT_CHAIN_PARA",
		"PCERT_RDN":                          "PCERT_RDN",
		"HCRYPTPROV_LEGACY":                  "HCRYPTPROV_LEGACY",
		"PCTL_USAGE":                         "PCTL_USAGE",
		"PCTL_VERIFY_USAGE_PARA":             "PCTL_VERIFY_USAGE_PARA",
		"PCTL_VERIFY_USAGE_STATUS":           "PCTL_VERIFY_USAGE_STATUS",
		"PCERT_CHAIN_POLICY_PARA":            "PCERT_CHAIN_POLICY_PARA",
		"PCERT_CHAIN_POLICY_STATUS":          "PCERT_CHAIN_POLICY_STATUS",
		"PCERT_REVOCATION_PARA":              "PCERT_REVOCATION_PARA",
		"PCERT_REVOCATION_STATUS":            "PCERT_REVOCATION_STATUS",
		"LPFILETIME":                         "LPFILETIME",
		"HCRYPTASYNC":                        "HCRYPTASYNC",
		"PHCRYPTASYNC":                       "PHCRYPTASYNC",
		"PCRYPT_DECRYPT_MESSAGE_PARA":        "PCRYPT_DECRYPT_MESSAGE_PARA",
		"PCRYPT_VERIFY_MESSAGE_PARA":         "PCRYPT_VERIFY_MESSAGE_PARA",
		"PCRYPT_DECODE_PARA":                 "PCRYPT_DECODE_PARA",
		"PCRYPT_ENCODE_PARA":                 "PCRYPT_ENCODE_PARA",
		"PCRYPT_ENCRYPT_MESSAGE_PARA":        "PCRYPT_ENCRYPT_MESSAGE_PARA",
		"HCRYPTOIDFUNCADDR":                  "HCRYPTOIDFUNCADDR",
		"HCRYPTOIDFUNCSET":                   "HCRYPTOIDFUNCSET",
		"PCRYPT_HASH_MESSAGE_PARA":           "PCRYPT_HASH_MESSAGE_PARA",
		"PCRYPT_PRIVATE_KEY_INFO":            "PCRYPT_PRIVATE_KEY_INFO",
		"PCRYPT_ATTRIBUTES":                  "PCRYPT_ATTRIBUTES",
		"PCRYPT_OID_FUNC_ENTRY":              "PCRYPT_OID_FUNC_ENTRY",
		"PCMSG_SIGNER_ENCODE_INFO":           "PCMSG_SIGNER_ENCODE_INFO",
		"PCMSG_SIGNED_ENCODE_INFO":           "PCMSG_SIGNED_ENCODE_INFO",
		"PCMSG_STREAM_INFO":                  "PCMSG_STREAM_INFO",
		"PCRYPTPROTECT_PROMPTSTRUCT":         "PCRYPTPROTECT_PROMPTSTRUCT",
		"PSIP_ADD_NEWPROVIDER":               "PSIP_ADD_NEWPROVIDER",
		"PSIP_INDIRECT_DATA":                 "PSIP_INDIRECT_DATA",
		"LPSIP_SUBJECTINFO":                  "LPSIP_SUBJECTINFO",
		"PMS_ADDINFO_FLAT":                   "PMS_ADDINFO_FLAT",
		"PMS_ADDINFO_CATALOGMEMBER":          "PMS_ADDINFO_CATALOGMEMBER",
		"PMS_ADDINFO_BLOB":                   "PMS_ADDINFO_BLOB",
		"PCRYPT_ATTRIBUTE_TYPE_VALUE":        "PCRYPT_ATTRIBUTE_TYPE_VALUE",
		"LPSIP_DISPATCH_INFO":                "LPSIP_DISPATCH_INFO",
		"PCRYPT_SIGN_MESSAGE_PARA":           "PCRYPT_SIGN_MESSAGE_PARA",
		"PCRYPT_KEY_SIGN_MESSAGE_PARA":       "PCRYPT_KEY_SIGN_MESSAGE_PARA",
		"PCRYPT_KEY_VERIFY_MESSAGE_PARA":     "PCRYPT_KEY_VERIFY_MESSAGE_PARA",
		"HLRUCACHE*":                         "*uintptr",
	}
	goType, ok = tablePtr[t]
	if ok {
		isPtr = true
		return goType, isConst, funcDecl, isPtr, nil
	}

	tableConst := map[string]string{
		"LPCSCROLLINFO":              "*SCROLLINFO",
		"LPCRECT":                    "*RECT",
		"LPCMENUITEMINFOW":           "*MENUITEMINFO",
		"LPCDLGTEMPLATEW":            "*DLGTEMPLATE",
		"LPCGUID":                    "*GUID",
		"PCRAWINPUTDEVICE":           "*RAWINPUTDEVICE",
		"LPCMENUINFO":                "*MENUINFO",
		"LPCVOID":                    "uintptr",
		"LPCPROPSHEETPAGEW":          "*PROPSHEETPAGE",
		"PCAUDIT_POLICY_INFORMATION": "PAUDIT_POLICY_INFORMATION",
		"PFORMAT_STRING":             "PFORMAT_STRING",
		"LPCSTR":                     "LPCSTR",
		"const unsigned char*":       "*byte",
		"LPCITEMIDLIST":              "LPCITEMIDLIST",
		"LPCTBBUTTON":                "*TBBUTTON",
		"LPCPROPSHEETHEADERW":        "*PROPSHEETHEADER",
		"LPCOLESTR":                  "LPCOLESTR",
		"REFCLSID":                   "REFCLSID",
		"LPCMMIOINFO":                "*MMIOINFO",
		"const char _huge*":          "HPSTR",
		"LPCWAVEFORMATEX":            "*WAVEFORMATEX",
		"LPCSHITEMID":                "*SHITEMID",
		"PCIDLIST_ABSOLUTE":          "PCIDLIST_ABSOLUTE",
		"PCZZWSTR":                   "PCZZWSTR",
		"PCCTL_CONTEXT":              "PCCTL_CONTEXT",
		"PCCERT_CONTEXT":             "PCCERT_CONTEXT",
		"PMIDL_STUB_DESC":            "PMIDL_STUB_DESC",
		"PCCERT_CHAIN_CONTEXT":       "PCCERT_CHAIN_CONTEXT",
		"PCCRYPT_OID_INFO":           "PCCRYPT_OID_INFO",
	}
	goType, ok = tableConst[t]
	if ok {
		isConst = true
		isPtr = true
		return goType, isConst, funcDecl, isPtr, nil
	}

	// Function types
	tableFunc := map[string]string{
		"WNDENUMPROC":                         "WNDENUMPROC",
		"PROPENUMPROCW":                       "PROPENUMPROC",
		"PROPENUMPROCEXW":                     "PROPENUMPROCEX",
		"WINSTAENUMPROCW":                     "WINSTAENUMPROC",
		"GRAYSTRINGPROC":                      "GRAYSTRINGPROC",
		"DLGPROC":                             "DLGPROC",
		"PFNCALLBACK":                         "PFNCALLBACK",
		"TIMERPROC":                           "TIMERPROC",
		"MONITORENUMPROC":                     "MONITORENUMPROC",
		"DRAWSTATEPROC":                       "DRAWSTATEPROC",
		"DESKTOPENUMPROCW":                    "DESKTOPENUMPROC",
		"WINEVENTPROC":                        "WINEVENTPROC",
		"HOOKPROC":                            "HOOKPROC",
		"SENDASYNCPROC":                       "SENDASYNCPROC",
		"HANDLER_FUNCTION_EX":                 "HANDLER_FUNCTION_EX",
		"LPHANDLER_FUNCTION_EX":               "HANDLER_FUNCTION_EX",
		"WNDPROC":                             "WNDPROC",
		"NDR_RUNDOWN":                         "NDR_RUNDOWN",
		"ENUMRESLANGPROCW":                    "ENUMRESLANGPROC",
		"FARPROC":                             "FARPROC",
		"THREAD_START_ROUTINE":                "THREAD_START_ROUTINE",
		"PTHREAD_START_ROUTINE":               "THREAD_START_ROUTINE",
		"LPTHREAD_START_ROUTINE":              "THREAD_START_ROUTINE",
		"LPCFHOOKPROC":                        "LPCFHOOKPROC",
		"PFNDAENUMCALLBACK":                   "DAENUMCALLBACK",
		"PFNDACOMPARE":                        "DACOMPARE",
		"PROPSHEETCALLBACK":                   "PROPSHEETCALLBACK",
		"SUBCLASSPROC":                        "SUBCLASSPROC",
		"TASKDIALOGCALLBACK":                  "TASKDIALOGCALLBACK",
		"LPFRHOOKPROC":                        "LPFRHOOKPROC",
		"LPOFNHOOKPROC":                       "LPOFNHOOKPROC",
		"LPPAGESETUPHOOK":                     "LPPAGESETUPHOOK",
		"LPPAGEPAINTHOOK":                     "LPPAGEPAINTHOOK",
		"LPSETUPHOOKPROC":                     "LPSETUPHOOKPROC",
		"ENHMFENUMPROC":                       "ENHMFENUMPROC",
		"ICMENUMPROCW":                        "ICMENUMPROC",
		"MFENUMPROC":                          "MFENUMPROC",
		"GOBJENUMPROC":                        "GOBJENUMPROC",
		"FONTENUMPROCW":                       "FONTENUMPROC",
		"LINEDDAPROC":                         "LINEDDAPROC",
		"ABORTPROC":                           "ABORTPROC",
		"PENUM_PAGE_FILE_CALLBACKW":           "PENUM_PAGE_FILE_CALLBACK",
		"IMCENUMPROC":                         "IMCENUMPROC",
		"REGISTERWORDENUMPROCW":               "REGISTERWORDENUMPROC",
		"LPLOOKUPSERVICE_COMPLETION_ROUTINE":  "LPLOOKUPSERVICE_COMPLETION_ROUTINE",
		"LPCONDITIONPROC":                     "LPCONDITIONPROC",
		"LPWSAOVERLAPPED_COMPLETION_ROUTINE":  "LPWSAOVERLAPPED_COMPLETION_ROUTINE",
		"LPWPUPOSTMESSAGE":                    "LPWPUPOSTMESSAGE",
		"PSECURE_MEMORY_CACHE_CALLBACK":       "PSECURE_MEMORY_CACHE_CALLBACK",
		"PIO_APC_ROUTINE":                     "PIO_APC_ROUTINE",
		"YIELDPROC":                           "YIELDPROC",
		"MMIOPROC":                            "MMIOPROC",
		"LPTIMECALLBACK":                      "TIMECALLBACK",
		"TIMECALLBACK":                        "TIMECALLBACK",
		"LPTASKCALLBACK":                      "TASKCALLBACK",
		"TASKCALLBACK":                        "TASKCALLBACK",
		"LPFNADDPROPSHEETPAGE":                "LPFNADDPROPSHEETPAGE",
		"LPFNDFMCALLBACK":                     "LPFNDFMCALLBACK",
		"BFFCALLBACK":                         "BFFCALLBACK",
		"PFN_RPCNOTIFICATION_ROUTINE":         "PFN_RPCNOTIFICATION_ROUTINE",
		"MIDL_ES_READ":                        "MIDL_ES_READ",
		"MIDL_ES_ALLOC":                       "MIDL_ES_ALLOC",
		"MIDL_ES_WRITE":                       "MIDL_ES_WRITE",
		"RPC_MGMT_AUTHORIZATION_FN":           "RPC_MGMT_AUTHORIZATION_FN",
		"RPC_AUTH_KEY_RETRIEVAL_FN":           "RPC_AUTH_KEY_RETRIEVAL_FN",
		"RPC_IF_CALLBACK_FN":                  "RPC_IF_CALLBACK_FN",
		"PFN_CERT_ENUM_SYSTEM_STORE":          "PFN_CERT_ENUM_SYSTEM_STORE",
		"PFN_CERT_ENUM_SYSTEM_STORE_LOCATION": "PFN_CERT_ENUM_SYSTEM_STORE_LOCATION",
		"PFN_CERT_ENUM_PHYSICAL_STORE":        "PFN_CERT_ENUM_PHYSICAL_STORE",
		"PFN_CRYPT_GET_SIGNER_CERTIFICATE":    "PFN_CRYPT_GET_SIGNER_CERTIFICATE",
		"PFN_CRYPT_ALLOC":                     "PFN_CRYPT_ALLOC",
		"PFN_CRYPT_FREE":                      "PFN_CRYPT_FREE",
		"PFN_CRYPT_ENUM_KEYID_PROP":           "PFN_CRYPT_ENUM_KEYID_PROP",
		"PFN_CRYPT_ENUM_OID_FUNC":             "PFN_CRYPT_ENUM_OID_FUNC",
		"PFN_CRYPT_ENUM_OID_INFO":             "PFN_CRYPT_ENUM_OID_INFO",
		"PFN_CRYPT_ASYNC_PARAM_FREE_FUNC*":    "PFN_CRYPT_ASYNC_PARAM_FREE_FUNC",
		"PFN_CRYPT_ASYNC_PARAM_FREE_FUNC":     "PFN_CRYPT_ASYNC_PARAM_FREE_FUNC",
		"PCRYPT_RESOLVE_HCRYPTPROV_FUNC":      "PCRYPT_RESOLVE_HCRYPTPROV_FUNC",
		"PCRYPT_DECRYPT_PRIVATE_KEY_FUNC":     "PCRYPT_DECRYPT_PRIVATE_KEY_FUNC",
		"PFN_CMSG_STREAM_OUTPUT":              "PFN_CMSG_STREAM_OUTPUT",
		"PCryptSIPGetSignedDataMsg":           "PCryptSIPGetSignedDataMsg",
		"PCryptSIPPutSignedDataMsg":           "PCryptSIPPutSignedDataMsg",
		"PCryptSIPCreateIndirectData":         "PCryptSIPCreateIndirectData",
		"PCryptSIPVerifyIndirectData":         "PCryptSIPVerifyIndirectData",
		"PCryptSIPRemoveSignedDataMsg":        "PCryptSIPRemoveSignedDataMsg",
	}
	goType, ok = tableFunc[t]
	if ok {
		fileName := TemplateFilePath(goType, "")
		if !IsFileExist(fileName) {
			panic(fmt.Errorf("Template file not found: %s", fileName))
		}
		fp, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		s := bufio.NewScanner(fp)
		reg := regexp.MustCompile(`^type [A-Za-z0-9_]+ func(.*)$`)
		funcDecl := ""
		for s.Scan() {
			line := s.Text()
			if strings.HasPrefix(line, ";") {
				continue
			}
			match := reg.FindSubmatch([]byte(line))
			if len(match) == 0 {
				continue
			}
			funcDecl = "func" + string(match[1])
			break
		}
		return goType, isConst, funcDecl, isPtr, nil
	}

	if strings.HasSuffix(t, "*") {
		tt := t
		pre := ""
		for strings.HasSuffix(tt, "*") {
			pre += "*"
			tt = strings.TrimSuffix(tt, "*")
		}
		fileName := TemplateFilePath(tt, "")
		if IsFileExist(fileName) {
			isPtr = true
			return pre + tt, isConst, funcDecl, isPtr, nil
		}
	}
	if strings.HasPrefix(t, "LP") {
		tt := strings.TrimPrefix(t, "LP")
		fileName := TemplateFilePath(tt, "")
		if IsFileExist(fileName) {
			isPtr = true
			return "*" + tt, isConst, funcDecl, isPtr, nil
		}
	}

	templateFile := TemplateFilePath(t, "")
	if IsFileExist(templateFile) {
		return t, isConst, funcDecl, isPtr, nil
	}
	return t, isConst, funcDecl, isPtr, fmt.Errorf("Unkown typename: %s", t)
}

func TemplateFilePath(typeName string, arch string) string {
	typeName = TemplateFileName(typeName)
	if len(typeName) == 0 {
		return ""
	}
	pre := string(typeName[0])
	name := typeName
	if arch != "" {
		name = name + "_" + arch
	}
	return filepath.Join("internal", "types", pre, name+".go")
}

func TemplateFileName(typeName string) string {
	// Don't use typename directly as a filename to avoid file name conflict (ex: byte.go and BYTE.go)
	str := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(str); i++ {
		c := string(str[i])
		typeName = strings.Replace(typeName, c, "-"+strings.ToUpper(c), -1)
	}
	return typeName
}
