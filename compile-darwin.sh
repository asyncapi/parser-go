#!/bin/bash

echo "Compiling for Mac OS X x64 (darwin)..."

go build -o bin/cparser-darwin-amd64.so.dylib -buildmode=c-shared cparser/cparser.go
