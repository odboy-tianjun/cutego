#!/bin/bash
# 更新 vendor 依赖
set -e
echo ">>> 清理无用依赖..."
go mod tidy
echo ">>> 同步 vendor 目录..."
go mod vendor
echo ">>> 完成"
