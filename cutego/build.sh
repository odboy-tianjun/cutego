#!/bin/bash
# 编译指定环境的二进制文件
# 用法: ./build.sh prod    → SujiejunServer
#       ./build.sh dev     → SujiejunServerTest

set -e

env="${1:-dev}"
case "$env" in
  prod) binary="cutego" ;;
  dev|test) binary="cutego_test" ;;
  *) echo "用法: $0 {dev|prod}" >&2; exit 1 ;;
esac

echo ">>> 编译 $env 环境 → $binary (linux amd64)..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o "$binary" .
echo ">>> 完成: $binary"
