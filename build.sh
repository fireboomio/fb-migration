#!/bin/bash

# 编译为不同平台架构的二进制文件
GOOS=linux GOARCH=amd64 go build -o fb-migration-linux
GOOS=windows GOARCH=amd64 go build -o fb-migration.exe
GOOS=darwin GOARCH=amd64 go build -o fb-migration-amd64

