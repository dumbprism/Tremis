@echo off

set EXE=tremis.exe

if not exist %EXE% (
    echo Building the project for the first time...
    go build -o %EXE% main.go math.go

    if %ERRORLEVEL% neq 0 (
        echo Build failed.
        exit /b %ERRORLEVEL%
    )
    echo Build succeeded.
) else (
    echo Executable already exists. Skipping build.
)

echo Running the project...
%EXE%
