package win

//ref SIZE_T
//ref LARGE_INTEGER

type QUOTA_LIMITS struct {
	PagedPoolLimit        SIZE_T
	NonPagedPoolLimit     SIZE_T
	MinimumWorkingSetSize SIZE_T
	MaximumWorkingSetSize SIZE_T
	PagefileLimit         SIZE_T
	padding1              [pad0for64_4for32]byte
	TimeLimit             LARGE_INTEGER
}
