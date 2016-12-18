package win

//ref DWORD
//ref LPWSTR

type LPSERVICE_MAIN_FUNCTION func(dwNumServicesArgs DWORD, lpServiceArgVectors *LPWSTR)
