#pragma once

#include <SDKDDKVer.h>

#include <stdio.h>
#include <tchar.h>
#include <math.h>

#define CINTERFACE

#include <Winsock2.h>
#include <ws2ipdef.h>
#include <WinDNS.h>
#include <Objbase.h>
#include <GuidDef.h>
#include <windows.h>
#include <objidl.h>
#include <gdiplus.h>
#include <commctrl.h>
#include <NTSecAPI.h>
#include <Pdh.h>
#include <IPHlpApi.h>
#include <aclapi.h>
#include <winddi.h>
#include <Uxtheme.h>
#include <Psapi.h>
#include <Rpc.h>
#include <WinCred.h>
#include <Shtypes.h>
#include <Prsht.h>
#include <Unknwn.h>
#include <OCIdl.h>
#include <Perflib.h>
#include <Shlwapi.h>
#include <Pdh.h>
#include <Olectl.h>
#include <winnls32.h>
#include <Dimm.h>
#include <Wincrypt.h>
#include <gl/GL.h>
#include <Wct.h>
#include <Ws2tcpip.h>
#include <WinSafer.h>
#include <fltdefs.h>
#include <Mmddk.h>
#include <Propvarutil.h>

typedef struct _IO_STATUS_BLOCK {
    union {
        NTSTATUS Status;
        PVOID Pointer;
    } DUMMYUNIONNAME;
    ULONG_PTR Information;
} IO_STATUS_BLOCK, *PIO_STATUS_BLOCK;


#define _PROPERTYKEY_EQUALITY_OPERATORS_
#include <Shlobj.h>

#include <gtest/gtest.h>

using namespace Gdiplus;
