#!/bin/bash

go build  -ldflags '-extldflags "-static"' --trimpath -o ../../bin/hellclient ../
go build  --trimpath -o ../../bin/mclconvertor ../mclconvertor

