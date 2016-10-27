@echo off

go run internal\cmd\gen\gen.go

set GOARCH=amd64
go run internal\cmd\test\test.go

set GOARCH=386
go run internal\cmd\test\test.go

pushd internal\ctests\build\amd64
del CMakeCache.txt
del cmake_install.cmake
del ctests.VC.db
rmdir /S /Q .vs
rmdir /S /Q CMakeFiles
rmdir /S /Q ctests.dir
rmdir /S /Q ipch
rmdir /S /Q x64
rmdir /Q /Q Debug
cmake ..\.. -G "Visual Studio 14 2015 Win64"
cmake --build .
Debug\ctests.exe
popd

pushd internal\ctests\build\386
del CMakeCache.txt
del cmake_install.cmake
del ctests.VC.db
rmdir /S /Q .vs
rmdir /S /Q CMakeFiles
rmdir /S /Q ctests.dir
rmdir /S /Q ipch
rmdir /S /Q Win32
rmdir /Q /Q Debug
cmake ..\.. -G "Visual Studio 14 2015"
cmake --build .
Debug\ctests.exe
popd
