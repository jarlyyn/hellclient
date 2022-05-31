#!/bin/bash
CC=aarch64-linux-gnu-gcc CC_FOR_TARGET=gcc-aarch64-linux-gnu CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -tags nopcre --trimpath -o ../../bin/hellclient.arm ../
GOARCH=arm64 GOOS=linux CC=aarch64-linux-gnu-gcc CC_FOR_TARGET=gcc-aarch64-linux-gnu go build  --trimpath -o ../../bin/mclconvertor.arm ../mclconvertor

