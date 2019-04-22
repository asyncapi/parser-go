echo "Compiling for Windows x64..."

go build -o bin/cparser-windows-amd64.dll -buildmode=c-shared cparser/cparser.go
