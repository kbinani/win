@echo off

call :main
exit /b %errorlevel%

:main
	(call :exec go run internal\cmd\gen\gen.go) > generate.log
	if not %errorlevel% == 0 (
		exit /b 1
	)

	call :setv GOARCH amd64
	call :exec go run internal\cmd\test\test.go
	if not %errorlevel% == 0 (
		call :setv GOARCH
		exit /b 2
	)

	call :setv GOARCH 386
	call :exec go run internal\cmd\test\test.go
	if not %errorlevel% == 0 (
		call :setv GOARCH
		exit /b 3
	)
	call :setv GOARCH

	call :exec pushd internal\ctests\build\amd64
	call :test_with_generator "Visual Studio 14 2015 Win64" amd64
	if not %errorlevel% == 0 (
		call :exec popd
		exit /b 4
	)
	call :exec popd

	call :exec pushd internal\ctests\build\386
	call :test_with_generator "Visual Studio 14 2015" 386
	if not %errorlevel% == 0 (
		call :exec popd
		exit /b 5
	)
	call :exec popd
	exit /b 0


:test_with_generator
	(call :exec cmake ..\.. -G "%~1") > ..\..\..\..\build_%2.log
	if not %errorlevel% == 0 (
		exit /b 1
	)

	(call :exec cmake --build .) >> ..\..\..\..\build_%2.log
	if not %errorlevel% == 0 (
		exit /b 1
	)
	(call :exec Debug\ctests.exe --gtest_output=xml:..\..\..\..\test_%~2.xml) >> ..\..\..\..\test_%~2.log
	if not %errorlevel% == 0 (
		exit /b 1
	)
	exit /b 0


:exec
	echo %* 1>&2
	%* 2>&1
	set ret=%errorlevel%
	if not %ret% == 0 (
		echo FAIL 1>&2
		exit /b %ret%
	)
	exit /b 0

:setv
	echo set %1=%2
	set %1=%2

