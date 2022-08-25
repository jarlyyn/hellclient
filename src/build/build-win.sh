#!/bin/bash
export lib=`cd $(dirname $0)/../../lib/pcre; pwd`
CGO_CFLAGS="-I$lib" CGO_LDFLAGS="-L$lib -lpcre" CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build  -tags nopcre -a --trimpath -o ../../bin/hellclient.exe ../
CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build --trimpath -o ../../bin/mclconvertor.exe ../mclconvertor