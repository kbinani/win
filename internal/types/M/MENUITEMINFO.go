package win

//ref HMENU
//ref HBITMAP
type MENUITEMINFO struct {
	CbSize        uint32
	FMask         uint32
	FType         uint32
	FState        uint32
	WID           uint32
	HSubMenu      HMENU
	HbmpChecked   HBITMAP
	HbmpUnchecked HBITMAP
	DwItemData    uintptr
	DwTypeData    *uint16
	Cch           uint32
	HbmpItem      HBITMAP
}
