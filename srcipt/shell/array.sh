#!/bin/bash
# 定义一个数组
arr=("xxx" 123 "bbb")

# for循环数组方式1, i是值
echo "======for循环数组方式1, i是值"
for i in ${arr[*]}; do
  echo "$i"
done

# for循环数组方式2, i是索引
echo "======for循环数组方式2, i是索引"
for i in ${!arr[*]}; do
  echo ${arr[$i]}
done

# for循环数组方式3
echo "======for循环数组方式3, 取数组下标"
arr_len=${#arr[*]}
for (( i = 0; i < $arr_len; i++ )); do
    echo ${arr[$i]}
done
