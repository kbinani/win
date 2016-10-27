package win

//ref GUID
//ref ULONG
type AUDIT_POLICY_INFORMATION struct {
	AuditSubCategoryGuid GUID
	AuditingInformation  ULONG
	AuditCategoryGuid    GUID
}
