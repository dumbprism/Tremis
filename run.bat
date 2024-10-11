@echo off

set EXE=tremis.exe

if not exist %EXE% (
    echo Welcome to tremis, setting up the server and getting things ready...
    go build -o %EXE% main.go math.go sets.go lists.go

    if %ERRORLEVEL% neq 0 (
        echo Build failed.
        exit /b %ERRORLEVEL%
    )
    echo Build succeeded. Enjoy using Tremis, happy hacking !
) else (
    echo Executable already exists. Skipping build.
)

echo Running the project...
%EXE%
