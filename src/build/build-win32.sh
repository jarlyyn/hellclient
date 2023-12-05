#!/bin/bash
export lib=`cd $(dirname $0)/../../lib/pcre; pwd`
CGO_CFLAGS="-I$lib" CC=i686-w64-mingw32-gcc GOOS=windows GOARCH=386 CGO_ENABLED=1 go build  -tags nopcre  --trimpath -o ../../bin/hellclient.exe ../
CC=i686-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build --trimpath -o ../../bin/mclconvertor.exe ../mclconvertor