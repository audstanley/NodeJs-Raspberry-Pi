#!/bin/bash
GOHOSTARCH=amd64
GOHOSTOS=linux
GOOS=linux
GOARCH=arm GOARM=6 go build -tags=arm -o=build/node-install-armv6l arm.go main.go;
GOARCH=arm GOARM=7 go build -tags=arm -o=build/node-install-armv7l arm.go main.go;
GOARCH=arm64 go build -o=build/node-install-arm64 x64.go main.go;
GOARCH=amd64 go build -o=build/node-install-x86_64 x64.go main.go;
scp build/node-install-armv7l pi@node1.local:/home/pi