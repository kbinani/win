package win

//ref HWND
//ref IContextMenuCB
//ref PCIDLIST_ABSOLUTE
//ref IShellFolder
//ref UINT
//ref PCUITEMID_CHILD_ARRAY
//ref IUnknown
//ref HKEY

type DEFCONTEXTMENU struct {
	Hwnd                HWND
	Pcmcb               *IContextMenuCB
	PidlFolder          PCIDLIST_ABSOLUTE
	Psf                 *IShellFolder
	Cidl                UINT
	Apidl               PCUITEMID_CHILD_ARRAY
	PunkAssociationInfo *IUnknown
	CKeys               UINT
	AKeys/*const*/ *HKEY
}
