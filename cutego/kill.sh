#!/bin/bash
# 停止指定环境的服务器进程
# 用法: ./kill.sh prod    → 从 application-prod.properties 读取端口
#       ./kill.sh dev     → 从 application-dev.properties 读取端口

env="${1:-dev}"
config="application-${env}.properties"

if [ ! -f "$config" ]; then
  echo "错误: 未找到 $config" >&2
  exit 1
fi

port=$(grep -E '^server\.port=' "$config" | head -1 | cut -d= -f2 | tr -d ' ')
if [ -z "$port" ]; then
  echo "错误: 在 $config 中未找到 server.port" >&2
  exit 1
fi

echo ">>> 查找端口 $port 上的进程..."
pid=$(lsof -ti:"$port" 2>/dev/null || true)
if [ -n "$pid" ]; then
  echo ">>> 找到 PID: $pid，正在停止..."
  kill -15 "$pid" 2>/dev/null || true
  sleep 2
  echo ">>> 已停止"
else
  echo ">>> 端口 $port 上无运行中的进程"
fi
