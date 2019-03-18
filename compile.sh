rm cparser/cparser.so cparser/cparser.so.dylib cparser/cparser.h

go build -o ./cparser/cparser.so -buildmode=c-shared ./cparser/cparser.go

g++ -dynamiclib -undefined suppress -flat_namespace ./cparser/*.so -o ./cparser/cparser.so.dylib
