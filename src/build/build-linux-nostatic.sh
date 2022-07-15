#!/bin/bash
export lib=`cd $(dirname $0)/../../lib/pcre; pwd`
CGO_LDFLAGS="-lpcre" CGO_ENABLED=1 go build  --trimpath -o ../../bin/hellclient ../
go build  -tags nopcre --trimpath -o ../../bin/hellclient ../
# CGO_LDFLAGS="-lpcre" CGO_ENABLED=1 go build  --trimpath -o ../../bin/hellclient ../
go build  --trimpath -o ../../bin/mclconvertor ../mclconvertor

