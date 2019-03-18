go build -o ./cparser/cparser.so -buildmode=c-shared ./cparser/cparser.go

g++ -dynamiclib -undefined suppress -flat_namespace ./cparser/*.so -o ./cparser/cparser.so.dylib
