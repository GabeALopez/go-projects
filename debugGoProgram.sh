#!/bin/bash

# Check if the source file is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <source_file.go>"
    exit 1
fi

SOURCE_FILE=$1
EXECUTABLE="gdbtest"

# Compile the Go source file with debug flags
go build -gcflags "-N -l" -o "$EXECUTABLE" "$SOURCE_FILE"

# Check if the compilation was successful
if [ $? -ne 0 ]; then
    echo "Compilation failed."
    exit 1
fi

# Run the executable with gdb
gdb "$EXECUTABLE"

# Delete the executable file after exiting gdb
rm -f "$EXECUTABLE"

echo "Executable file deleted."

