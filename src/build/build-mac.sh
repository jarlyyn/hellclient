#!/bin/bash
export lib=`cd $(dirname $0)/../../lib/pcre; pwd`
CGO_CFLAGS="-I$lib" CC=o64-gcc CGO_LDFLAGS="-lpcre" CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build  --trimpath -o ../../bin/hellclient-mac ../

go build  --trimpath -o ../../bin/mclconvertor-mac ../mclconvertor

