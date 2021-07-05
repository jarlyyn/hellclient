#!/bin/bash

CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -ldflags '-extldflags "-static"' --trimpath -o ../../bin/hellclient.exe ../
CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build --trimpath -o ../../bin/mclconvertor.exe ../mclconvertor

