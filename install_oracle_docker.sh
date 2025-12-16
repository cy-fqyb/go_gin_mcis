#!/bin/bash

# --------------------------------------
# 脚本说明：
# 自动拉取 Oracle XE 21c Docker 镜像，并启动容器
# 容器名称：oracle-xe
# 监听端口：1521 (SQL*Net), 5500 (EM Express)
# 默认 SYS/SYSTEM 用户密码由你设置
# --------------------------------------

# 配置参数
CONTAINER_NAME="oracle-xe"
SYS_PASSWORD="Oracle123"    # 修改为你自己的密码
PDB_NAME="XEPDB1"           # 默认 pluggable database
HOST_PORT_DB=1521
HOST_PORT_EM=5500
IMAGE_NAME="gvenzl/oracle-xe:21-slim"

# 检查 Docker 是否安装
if ! command -v docker &> /dev/null
then
    echo "Docker 未安装，请先安装 Docker"
    exit 1
fi

# 拉取 Oracle XE 镜像
echo "正在拉取 Oracle XE Docker 镜像..."
docker pull $IMAGE_NAME

# 检查是否已有同名容器在运行
if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" ]; then
    echo "已有容器 $CONTAINER_NAME，先停止并删除它..."
    docker stop $CONTAINER_NAME
    docker rm $CONTAINER_NAME
fi

# 启动容器
echo "启动 Oracle XE 容器..."
docker run -d \
    --name $CONTAINER_NAME \
    -p $HOST_PORT_DB:1521 \
    -p $HOST_PORT_EM:5500 \
    -e ORACLE_PASSWORD=$SYS_PASSWORD \
    $IMAGE_NAME

# 等待数据库启动（约 30 秒）
echo "等待数据库启动中..."
sleep 30

# 检查容器状态
docker ps -f name=$CONTAINER_NAME

echo "Oracle XE Docker 容器已启动！"
echo "连接信息："
echo "  主机: localhost"
echo "  端口: $HOST_PORT_DB"
echo "  服务名: $PDB_NAME"
echo "  用户: SYS/SYSTEM"
echo "  密码: $SYS_PASSWORD"
echo "  以 SYSDBA 连接: sqlplus sys/$SYS_PASSWORD@localhost:$HOST_PORT_DB/$PDB_NAME as sysdba"
