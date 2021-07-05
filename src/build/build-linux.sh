#!/bin/bash
export lib=`cd $(dirname $0)/../../lib/pcre; pwd`
CGO_LDFLAGS="-lpcre -static" CGO_ENABLED=1 go build  --trimpath -o ../../bin/app ../
go build  --trimpath -o ../../bin/mclconvertor ../mclconvertor

