#!/bin/bash
CGO_ENABLED=1 go build -tags 'netgo' -a --ldflags "-linkmode external -extldflags '-static'" --trimpath -o ../../bin/hellclient ../
CGO_ENABLED=0 go build  --trimpath -o ../../bin/mclconvertor ../mclconvertor


