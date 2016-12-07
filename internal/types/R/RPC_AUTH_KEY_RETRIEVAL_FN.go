package win

//ref PVOID
//ref RPC_WSTR
//ref RPC_STATUS

type RPC_AUTH_KEY_RETRIEVAL_FN func(Arg PVOID, ServerPrincName RPC_WSTR, KeyVer uint32, Key *PVOID, Status *RPC_STATUS)
