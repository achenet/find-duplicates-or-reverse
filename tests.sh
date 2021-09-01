#!/bin/sh

echo "Testing"

# Build
go build -o find-duplicates-or-reverse solution/*.go

# Test with various test files
./find-duplicates-or-reverse test-files/1
./find-duplicates-or-reverse test-files/2

# Test with command line input
