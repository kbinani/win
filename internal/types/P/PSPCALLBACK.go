package win

//ref HWND
//ref UINT
//ref PROPSHEETPAGE
type PSPCALLBACK func(hwnd HWND, uMsg UINT, ppsp *PROPSHEETPAGE) UINT
