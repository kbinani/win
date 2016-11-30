package win

//ref IShellView
//ref IShellFolder
//ref HWND
//ref UINT
//ref WPARAM
//ref LPARAM
//ref HRESULT

type LPFNVIEWCALLBACK func(psvOuter *IShellView, psf *IShellFolder, hwndMain HWND, uMsg UINT, wParam WPARAM, lParam LPARAM) HRESULT
