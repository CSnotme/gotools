#!/bin/bash


# 获取执行脚本的当前目录绝对路径
echo "=======获取执行脚本的当前目录绝对路径"
#DIR_PATH=$(cd "$(dirname "$0")" && pwd)  # 通过dirname+pwd获取
#DIR_PATH=$(dirname "$(readlink -f "$0")")  # 通过dirname+readlink获取
DIR_PATH=$(pwd)
echo "${DIR_PATH}"

# 获取执行脚本的文件名
echo "=======获取执行脚本的文件名"
FILE_NAME="$(basename "$0")"
echo "${FILE_NAME}"

# 获取执行脚本的绝对路径
echo "=======获取执行脚本的绝对路径"
#FILE_PATH="$(pwd)/$(basename "$0")"  # 通过pwd+basename获取
FILE_PATH="$(readlink -f "$0")"
echo "${FILE_PATH}"



