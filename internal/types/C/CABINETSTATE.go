package win

//ref WORD
//ref UINT
//ref BOOL

type CABINETSTATE struct {
	CLength         WORD
	NVersion        WORD
	flags1          uint16
	FMenuEnumFilter UINT
}

func (this *CABINETSTATE) FFullPathTitle() BOOL {
	return (BOOL)(0x1 & (this.flags1 >> 15))
}

func (this *CABINETSTATE) FSaveLocalView() BOOL {
	return (BOOL)(0x1 & (this.flags1 >> 14))
}

func (this *CABINETSTATE) FNotShell() BOOL {
	return (BOOL)(0x1 & (this.flags1 >> 13))
}

func (this *CABINETSTATE) FSimpleDefault() BOOL {
	return (BOOL)(0x1 & (this.flags1 >> 12))
}

func (this *CABINETSTATE) FDontShowDescBar() BOOL {
	return (BOOL)(0x1 & (this.flags1 >> 11))
}

func (this *CABINETSTATE) FNewWindowMode() BOOL {
	return (BOOL)(0x1 & (this.flags1 >> 10))
}

func (this *CABINETSTATE) FShowCompColor() BOOL {
	return (BOOL)(0x1 & (this.flags1 >> 9))
}

func (this *CABINETSTATE) FDontPrettyNames() BOOL {
	return (BOOL)(0x1 & (this.flags1 >> 8))
}

func (this *CABINETSTATE) FAdminsCreateCommonGroups() BOOL {
	return (BOOL)(0x1 & (this.flags1 >> 7))
}
