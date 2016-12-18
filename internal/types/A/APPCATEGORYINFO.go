package win

//ref LCID
//ref LPWSTR
//ref GUID

type APPCATEGORYINFO struct {
	Locale         LCID
	PszDescription LPWSTR
	AppCategoryId  GUID
}
