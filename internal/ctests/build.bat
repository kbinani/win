@echo off

call :main
exit /b %errorlevel%

:main
	go run internal\cmd\gen\gen.go
	if not %errorlevel% == 0 (
		exit /b 1
	)

	set GOARCH=amd64
	go run internal\cmd\test\test.go
	if not %errorlevel% == 0 (
		exit /b 2
	)

	set GOARCH=386
	go run internal\cmd\test\test.go
	if not %errorlevel% == 0 (
		exit /b 3
	)

	pushd internal\ctests\build\amd64
	call :test_with_generator "Visual Studio 14 2015 Win64"
	if not %errorlevel% == 0 (
		popd
		exit /b 4
	)
	popd

	pushd internal\ctests\build\386
	call :test_with_generator "Visual Studio 14 2015"
	if not %errorlevel% == 0 (
		popd
		exit /b 5
	)
	popd
	exit /b 0


:test_with_generator
	call :clean_cmake_cache
	cmake ..\.. -G "%~1"
	if not %errorlevel% == 0 (
		exit /b %errorlevel%
	)
	cmake --build .
	if not %errorlevel% == 0 (
		exit /b %errorlevel%
	)
	Debug\ctests.exe
	if not %errorlevel% == 0 (
		exit /b %errorlevel%
	)
	exit /b 0


:clean_cmake_cache
	del CMakeCache.txt
	del cmake_install.cmake
	del ctests.VC.db
	rmdir /S /Q .vs
	rmdir /S /Q CMakeFiles
	rmdir /S /Q ctests.dir
	rmdir /S /Q ipch
	rmdir /S /Q Win32
	rmdir /S /Q x64
	rmdir /S /Q Debug
	exit /b 0
