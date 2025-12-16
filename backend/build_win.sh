#!/bin/bash
# ================================================
# build_win.sh - Linux 下交叉编译 Go 程序为 Windows exe
# ================================================

# 1️⃣ 配置
APP_NAME="gin_fqyb_end_case"              # 可执行文件名（不含扩展名）
MAIN_FILE="main.go"           # Go 程序入口文件
OUTPUT_DIR="dist"             # 输出目录
CONFIG_FILES=("config.yaml")  # 需要打包的配置文件
GOOS_TARGET="windows"         # 目标系统
GOARCH_TARGET="amd64"         # 目标架构 (64 位)
LDFLAGS="-s -w"               # 可选: 去掉符号和 DWARF，减小体积

# 2️⃣ 创建输出目录
mkdir -p $OUTPUT_DIR

# 3️⃣ 设置环境变量
export GOOS=$GOOS_TARGET
export GOARCH=$GOARCH_TARGET

# 4️⃣ 编译
echo "🔧 编译 $MAIN_FILE 为 $GOOS/$GOARCH..."
go build -ldflags "$LDFLAGS" -o "$OUTPUT_DIR/$APP_NAME.exe" $MAIN_FILE
if [ $? -ne 0 ]; then
    echo "❌ 编译失败"
    exit 1
fi

echo "✅ 编译成功: $OUTPUT_DIR/$APP_NAME.exe"

# 5️⃣ 复制配置文件
for cfg in "${CONFIG_FILES[@]}"; do
    if [ -f "$cfg" ]; then
        cp "$cfg" "$OUTPUT_DIR/"
        echo "📄 复制配置文件: $cfg"
    fi
done

echo "🎉 打包完成，输出目录: $OUTPUT_DIR"
