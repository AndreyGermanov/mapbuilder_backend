#!/bin/bash
systemctl stop mapbuilder
mkdir -p /opt/mapbuilder
/usr/local/go/bin/go build -o mapbuilder cmd/mapbuilder/mapbuilder.go
mv -f ir /opt/mapbuilder/
ln -s geodata /opt/mapbuilder/geodata
systemctl start mapbuilder