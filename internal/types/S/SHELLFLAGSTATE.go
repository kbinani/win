package win

type SHELLFLAGSTATE struct {
	storage1 BOOL
}

func (this *SHELLFLAGSTATE) FShowAllObjects() BOOL {
	return BOOL(0x1 & (this.storage1 >> 31))
}

func (this *SHELLFLAGSTATE) FShowExtensions() BOOL {
	return BOOL(0x1 & (this.storage1 >> 30))
}

func (this *SHELLFLAGSTATE) FNoConfirmRecycle() BOOL {
	return BOOL(0x1 & (this.storage1 >> 29))
}

func (this *SHELLFLAGSTATE) FShowSysFiles() BOOL {
	return BOOL(0x1 & (this.storage1 >> 28))
}

func (this *SHELLFLAGSTATE) FShowCompColor() BOOL {
	return BOOL(0x1 & (this.storage1 >> 27))
}

func (this *SHELLFLAGSTATE) FDoubleClickInWebView() BOOL {
	return BOOL(0x1 & (this.storage1 >> 26))
}

func (this *SHELLFLAGSTATE) FDesktopHTML() BOOL {
	return BOOL(0x1 & (this.storage1 >> 25))
}

func (this *SHELLFLAGSTATE) FWin95Classic() BOOL {
	return BOOL(0x1 & (this.storage1 >> 24))
}

func (this *SHELLFLAGSTATE) FDontPrettyPath() BOOL {
	return BOOL(0x1 & (this.storage1 >> 23))
}

func (this *SHELLFLAGSTATE) FShowAttribCol() BOOL {
	return BOOL(0x1 & (this.storage1 >> 22))
}

func (this *SHELLFLAGSTATE) FMapNetDrvBtn() BOOL {
	return BOOL(0x1 & (this.storage1 >> 21))
}

func (this *SHELLFLAGSTATE) FShowInfoTip() BOOL {
	return BOOL(0x1 & (this.storage1 >> 20))
}

func (this *SHELLFLAGSTATE) FHideIcons() BOOL {
	return BOOL(0x1 & (this.storage1 >> 19))
}

func (this *SHELLFLAGSTATE) FAutoCheckSelect() BOOL {
	return BOOL(0x1 & (this.storage1 >> 18))
}

func (this *SHELLFLAGSTATE) FIconsOnly() BOOL {
	return BOOL(0x1 & (this.storage1 >> 17))
}
