@echo off

call :main "%~1" "%~2"
exit /b %errorlevel%

:main
    set OUTPUT=%~1
    set COMMIT_ID=%~2

    mkdir "%OUTPUT%" >nul 2>&1
    rmdir "%OUTPUT%\.git" /S /Q >nul 2>&1
    call :exec xcopy ..\win "%OUTPUT%\" /S /E /H /Q /Y
    if not %errorlevel% == 0 (
        exit /b 2
    )
    
    call :exec pushd "%OUTPUT%"
    call :exec git checkout -- .
    if not %errorlevel% == 0 (
        call :exec popd
        exit /b 3
    )
    call :exec git submodule deinit .
    if not %errorlevel% == 0 (
        call :exec popd
        exit /b 3
    )
    call :exec git checkout master
    if not %errorlevel% == 0 (
        call :exec popd
        exit /b 3
    )
    call :exec git clean -xdf
    if not %errorlevel% == 0 (
        call :exec popd
        exit /b 3
    )
    call :exec del *.go
    if not %errorlevel% == 0 (
        call :exec popd
        exit /b 4
    )
    call :exec popd
    
    call :exec go run internal\cmd\gen\gen.go --output="%OUTPUT%"
    if not %errorlevel% == 0 (
        exit /b 5
    )

    call :exec pushd "%OUTPUT%"
    call :exec git add . -A
    if not %errorlevel% == 0 (
        call :exec popd
        exit /b 6
    )
    call :exec git commit -m "Generated from: https://github.com/kbinani/win/commit/%COMMIT_ID%"
    if not %errorlevel% == 0 (
        call :exec popd
        exit /b 7
    )
    call :exec git push origin master
    if not %errorlevel% == 0 (
        call :exec popd
        exit /b 8
    )
    call :exec popd

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
