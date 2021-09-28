#!/bin/bash
GOHOSTARCH=amd64
GOHOSTOS=linux
GOOS=linux
GOARCH=arm GOARM=6 go build -tags=arm -o=build/node-install-armv6l arm.go main.go;
GOARCH=arm GOARM=7 go build -tags=arm -o=build/node-install-armv7l arm.go main.go;
GOARCH=arm64 go build -o=build/node-install-arm64 x64.go main.go;
GOARCH=amd64 go build -o=build/node-install-x86_64 x64.go main.go;
sha256sum build/node-install-arm64 | awk '{print $1}' > build/sha256-node-install-arm64.checksum
sha256sum build/node-install-armv6l | awk '{print $1}' > build/sha256-node-install-armv6l.checksum
sha256sum build/node-install-armv7l | awk '{print $1}' > build/sha256-node-install-armv7l.checksum
sha256sum build/node-install-x86_64 | awk '{print $1}' > build/sha256-node-install-x86_64.checksum
# scp build/node-install-armv7l pi@node1.local:/home/pi