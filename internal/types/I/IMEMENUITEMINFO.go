package win

//ref UINT
//ref HBITMAP
//ref DWORD
type IMEMENUITEMINFO struct {
	CbSize        UINT
	FType         UINT
	FState        UINT
	WID           UINT
	HbmpChecked   HBITMAP
	HbmpUnchecked HBITMAP
	DwItemData    DWORD
	SzString      [IMEMENUITEM_STRING_SIZE]WCHAR
	HbmpItem      HBITMAP
}
