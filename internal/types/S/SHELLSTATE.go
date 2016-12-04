package win

//ref BOOL
//ref DWORD
//ref UINT
//ref LONG

type SHELLSTATE struct {
	storage1       BOOL
	DwWin95Unused  DWORD
	UWin95Unused   UINT
	LParamSort     LONG
	ISortDirection int32
	Version        UINT
	UNotUsed       UINT
	storage2       uint8
	storage3       uint16
}

func (this *SHELLSTATE) FShowAllObjects() BOOL {
	return BOOL(0x1 & (this.storage1 >> 31))
}

func (this *SHELLSTATE) FShowExtensions() BOOL {
	return BOOL(0x1 & (this.storage1 >> 30))
}

func (this *SHELLSTATE) FNoConfirmRecycle() BOOL {
	return BOOL(0x1 & (this.storage1 >> 29))
}

func (this *SHELLSTATE) FShowSysFiles() BOOL {
	return BOOL(0x1 & (this.storage1 >> 28))
}

func (this *SHELLSTATE) FShowCompColor() BOOL {
	return BOOL(0x1 & (this.storage1 >> 27))
}

func (this *SHELLSTATE) FDoubleClickInWebView() BOOL {
	return BOOL(0x1 & (this.storage1 >> 26))
}

func (this *SHELLSTATE) FDesktopHTML() BOOL {
	return BOOL(0x1 & (this.storage1 >> 25))
}

func (this *SHELLSTATE) FWin95Classic() BOOL {
	return BOOL(0x1 & (this.storage1 >> 24))
}

func (this *SHELLSTATE) FDontPrettyPath() BOOL {
	return BOOL(0x1 & (this.storage1 >> 23))
}

func (this *SHELLSTATE) FShowAttribCol() BOOL {
	return BOOL(0x1 & (this.storage1 >> 22))
}

func (this *SHELLSTATE) FMapNetDrvBtn() BOOL {
	return BOOL(0x1 & (this.storage1 >> 21))
}

func (this *SHELLSTATE) FShowInfoTip() BOOL {
	return BOOL(0x1 & (this.storage1 >> 20))
}

func (this *SHELLSTATE) FHideIcons() BOOL {
	return BOOL(0x1 & (this.storage1 >> 19))
}

func (this *SHELLSTATE) FWebView() BOOL {
	return BOOL(0x1 & (this.storage1 >> 18))
}

func (this *SHELLSTATE) FFilter() BOOL {
	return BOOL(0x1 & (this.storage1 >> 17))
}

func (this *SHELLSTATE) FShowSuperHidden() BOOL {
	return BOOL(0x1 & (this.storage1 >> 16))
}

func (this *SHELLSTATE) FNoNetCrawling() BOOL {
	return BOOL(0x1 & (this.storage1 >> 15))
}

func (this *SHELLSTATE) FSepProcess() BOOL {
	return BOOL(0x1 & (this.storage2 >> 15))
}

func (this *SHELLSTATE) FStartPanelOn() BOOL {
	return BOOL(0x1 & (this.storage2 >> 14))
}

func (this *SHELLSTATE) FShowStartPage() BOOL {
	return BOOL(0x1 & (this.storage2 >> 13))
}

func (this *SHELLSTATE) FAutoCheckSelect() BOOL {
	return BOOL(0x1 & (this.storage2 >> 12))
}

func (this *SHELLSTATE) FIconsOnly() BOOL {
	return BOOL(0x1 & (this.storage2 >> 11))
}

func (this *SHELLSTATE) FShowTypeOverlay() BOOL {
	return BOOL(0x1 & (this.storage2 >> 10))
}

func (this *SHELLSTATE) FShowStatusBar() BOOL {
	return BOOL(0x1 & (this.storage2 >> 9))
}

func (this *SHELLSTATE) FSpareFlags() UINT {
	return UINT(0x1FF & (this.storage3 >> 7))
}
