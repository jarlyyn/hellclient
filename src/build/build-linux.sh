#!/bin/bash
export lib=`cd $(dirname $0)/../../lib/pcre;pwd`
CGO_CFLAGS="-I/$lib/" CGO_LDFLAGS="-L$lib/ -static" CC="musl-gcc" CGO_ENABLED=1 go build -tags 'musl' -a --ldflags '-linkmode external -extldflags "-static"' --trimpath -o ../../bin/hellclient ../
# CGO_LDFLAGS="-lpcre" CGO_ENABLED=1 go build  --trimpath -o ../../bin/hellclient ../
go build  --trimpath -o ../../bin/mclconvertor ../mclconvertor

