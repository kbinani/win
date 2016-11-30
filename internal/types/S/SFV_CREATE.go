package win

//ref UINT
//ref IShellFolder
//ref IShellView
//ref IShellFolderViewCB

type SFV_CREATE struct {
	CbSize   UINT
	Pshf     *IShellFolder
	PsvOuter *IShellView
	Psfvcb   *IShellFolderViewCB
}
