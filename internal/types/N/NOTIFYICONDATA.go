package win

//ref DWORD
//ref HWND
//ref UINT
//ref HICON
//ref WCHAR
//ref GUID

type NOTIFYICONDATA struct {
	CbSize           DWORD
	HWnd             HWND
	UID              UINT
	UFlags           UINT
	UCallbackMessage UINT
	HIcon            HICON
	SzTip            [128]WCHAR
	DwState          DWORD
	DwStateMask      DWORD
	SzInfo           [256]WCHAR
	union1           UINT
	SzInfoTitle      [64]WCHAR
	DwInfoFlags      DWORD
	GuidItem         GUID
	HBalloonIcon     HICON
}

func (this *NOTIFYICONDATA) GetUTimeout() UINT {
	return this.union1
}

func (this *NOTIFYICONDATA) SetUTimeout(v UINT) {
	this.union1 = v
}

func (this *NOTIFYICONDATA) GetUVersion() UINT {
	return this.union1
}

func (this *NOTIFYICONDATA) SetUVersion(v UINT) {
	this.union1 = v
}
