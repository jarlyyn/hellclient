#!/bin/bash
export lib=`cd $(dirname $0)/../../lib/pcre-musl; pwd`
CGO_LDFLAGS="-L$lib -Wl,-rpath=$lib -s -w" CC="musl-gcc" CGO_ENABLED=1 go build -tags 'musl netgo' -a --ldflags "-linkmode external -extldflags '-static'" --trimpath -o ../../bin/hellclient ../
CGO_ENABLED=0 go build  --trimpath -o ../../bin/mclconvertor ../mclconvertor


