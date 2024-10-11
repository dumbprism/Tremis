#!/bin/bash

# Set the output executable name
EXE="tremis"

# Check if the executable exists
if [ ! -f "$EXE" ]; then
    echo "Welcome to tremis, setting up the server and getting things ready..."
    # Build the Go project
    go build -o "$EXE" main.go math.go sets.go

    if [ $? -ne 0 ]; then
        echo "Build failed."
        exit 1
    fi
    echo "Build succeeded. Enjoy using Tremis, happy hacking !"
else
    echo "Executable already exists. Skipping build."
fi

# Run the executable
echo "Running the project..."
./"$EXE"
