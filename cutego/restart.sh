#!/bin/bash
# 编译 + 停止旧进程 + 启动新进程（一键重启）
# 用法: ./restart.sh prod   → 生产环境
#       ./restart.sh dev    → 开发环境

set -e

env="${1:-dev}"
config="application-${env}.properties"

if [ ! -f "$config" ]; then
  echo "错误: 未找到 $config" >&2
  exit 1
fi

port=$(grep -E '^server\.port=' "$config" | head -1 | cut -d= -f2 | tr -d ' ')
binary=$( [ "$env" = "prod" ] && echo "cutego" || echo "cutego_test" )

echo "=== 重启 $env 环境 ==="

# 1. 编译
echo ">>> 编译 $binary (linux amd64)..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o "$binary" .

# 2. 停止旧进程
echo ">>> 查找端口 $port 上的进程..."
pid=$(lsof -ti:"$port" 2>/dev/null || true)
if [ -n "$pid" ]; then
  echo ">>> 停止 PID: $pid..."
  kill -15 "$pid" 2>/dev/null || true
  sleep 2
fi

# 3. 启动
echo ">>> 启动 $binary (active=$env, port=$port)..."
nohup "./$binary" "active=$env" > application.log 2>&1 &
sleep 2

if kill -0 $! 2>/dev/null; then
  echo "✅ 启动成功  PID: $!  端口: $port"
else
  echo "❌ 启动失败，查看 application.log"
  exit 1
fi
