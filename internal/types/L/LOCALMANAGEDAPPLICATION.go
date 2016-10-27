package win

//ref LPWSTR
type LOCALMANAGEDAPPLICATION struct {
	PszDeploymentName LPWSTR
	PszPolicyName     LPWSTR
	PszProductId      LPWSTR
	DwState           uint32
}
