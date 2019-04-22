#!/bin/bash

echo "Compiling for Linux x64..."

go build -o bin/cparser-linux-amd64.so -buildmode=c-shared cparser/cparser.go
