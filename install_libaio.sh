#!/bin/bash

# 设置变量
LIBAIO_PATH="/usr/lib/x86_64-linux-gnu"
LIBAIO_T64="libaio.so.1t64"
LIBAIO="libaio.so.1"

# 检查 libaio1t64 是否已安装
if ! dpkg -l | grep -q libaio1t64; then
  echo "正在安装 libaio1t64..."
  sudo apt update
  sudo apt install -y libaio1t64
else
  echo "libaio1t64 已安装。"
fi

# 创建符号链接
if [ ! -f "$LIBAIO_PATH/$LIBAIO" ]; then
  echo "创建符号链接 $LIBAIO_PATH/$LIBAIO..."
  sudo ln -s "$LIBAIO_PATH/$LIBAIO_T64" "$LIBAIO_PATH/$LIBAIO"
else
  echo "符号链接 $LIBAIO_PATH/$LIBAIO 已存在。"
fi

# 更新动态链接器缓存
echo "更新动态链接器缓存..."
sudo ldconfig

echo "libaio.so.1 安装和配置完成。"
