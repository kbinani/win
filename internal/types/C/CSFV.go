package win

//ref LPFNVIEWCALLBACK
//ref UINT
//ref IShellFolder
//ref IShellView
//ref PCIDLIST_ABSOLUTE
//ref LONG
//ref FOLDERVIEWMODE

type CSFV struct {
	CbSize      UINT
	Pshf        *IShellFolder
	PsvOuter    *IShellView
	Pidl        PCIDLIST_ABSOLUTE
	LEvents     LONG
	PfnCallback LPFNVIEWCALLBACK
	Fvm         FOLDERVIEWMODE
}
