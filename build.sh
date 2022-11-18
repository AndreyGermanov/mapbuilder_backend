#!/bin/bash
systemctl stop mapbuilder
mkdir -p /opt/mapbuilder
/usr/local/go/bin/go build -o mapbuilder cmd/mapbuilder/mapbuilder.go
mv -f mapbuilder /opt/mapbuilder/
ln -sf geodata /opt/mapbuilder/geodata
systemctl start mapbuilder