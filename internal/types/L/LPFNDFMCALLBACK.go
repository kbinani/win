package win

//ref IShellFolder
//ref HWND
//ref IDataObject
//ref UINT
//ref WPARAM
//ref LPARAM
//ref HRESULT

type LPFNDFMCALLBACK func(psf *IShellFolder, hwnd HWND, pdtobj *IDataObject, uMsg UINT, wParam WPARAM, lParam LPARAM) HRESULT
