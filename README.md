Build status
===

[![Build status](https://ci.appveyor.com/api/projects/status/r4r0dlqm9fj3c9la/branch/generator?svg=true)](https://ci.appveyor.com/project/kbinani/win/branch/generator)

How to
===

Generate Go source files
---

```
go run internal/cmd/gen/gen.go
```

Test generated source files 
---

`internal/cmd/test/test.go` command-line utility generates C++ test case files by using `reflect` to check size of `struct` and offset of struct member fields.

* amd64
    ```
    set GOARCH=amd64
    go run internal\cmd\test\test.go
    cd internal\ctests\build\amd64
    cmake ..\..\ -G "Visual Studio 14 2015 Win64"
    cmake --build .
    Debug\ctests.exe
    ```

* 386
    ```
    set GOARCH=386
    go run internal\cmd\test\test.go
    cd internal\ctests\build\386
    cmake ..\..\ -G "Visual Studio 14 2015"
    cmake --build .
    Debug\ctests.exe
    ```

Design of generator
===

Workflow of the code generator
---

1. Download MinGW tarball
1. Download Wine tarball
1. Create list of WinAPI functions from `*.def` files in MinGW for each dll (`kernel32.dll` etc.)
1. Search API definition from `*.h` files in MinGW and Wine tarball
1. Generate API wrapper. `struct` and `func` definitions are loaded from `internal/types` or `internal/funcs` if needed

`struct` and `func` templates
---

All `struct`'s are manually crafted, and stored under `internal/types` directory.
Some `func`'s are also stored under `internal/funcs`. Other WinAPI functions except them are automatically generated.
These template files may have `_386.go` or `_amd64.go` suffix. These suffix means that the file is specifically made for the architecture.
