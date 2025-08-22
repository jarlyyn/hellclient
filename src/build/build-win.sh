#!/bin/bash
export lib=`cd $(dirname $0)/../../lib/pcre;pwd`
CGO_CFLAGS="-I$lib" CGO_LDFLAGS="-L$lib -lpcre -s -w" CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build  --trimpath -o ../../bin/hellclient.exe ../
#CGO_ENABLED=0 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build --trimpath -o ../../bin/mclconvertor.exe ../mclconvertor
