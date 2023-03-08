#!/bin/bash


apps=("HC_FIELD" "HC_DATASERVICE" "HC_ASSEMBLY")

HC_FIELD_DIR="hc-field"       # 字段服务项目地址
HC_FIELD_APP_NAME="field"       # 字段服务项目名称
HC_FIELD_GIT_BRANCH="master"    # 字段服务项目使用git分支


HC_DATASERVICE_DIR="hc-dataservice"  # 数据服务项目地址
HC_DATASERVICE_APP_NAME="dataservice"  # 数据服务项目名称
HC_DATASERVICE_GIT_BRANCH="new-4s"     # 数据服务项目使用git分支


HC_ASSEMBLY_DIR="hc-assembly"     # api项目地址
HC_ASSEMBLY_APP_NAME="assembly"     # api项目名称
HC_ASSEMBLY_GIT_BRANCH="develop"    # api项目使用git分支

SHELL_FOLDER=$(cd "$(dirname "$0")" && pwd)
RUN_APPS=()
RUN_CMD="run"

function help() {
  echo "=====help====="
  echo "usage: run_tools.sh cmd app"
  echo "  [cmd]:"
  echo "    kill: 如果服务存在进程，kill掉"
  echo "    build: 仅编译"
  echo "    pullbuild: 拉取分支代码后编译"
  echo "    run: 编译并启动"
  echo "    pullrun: 拉取分支代码后编译并启动"
  echo "    onlyrun: 不再编译直接启动"
  echo "  [app]:"
  for app in ${apps[*]} ; do
      local app_dir=$(echo `eval echo '$'"$app""_DIR"`)
      local app_name=$(echo `eval echo '$'"$app""_APP_NAME"`)
      local app_branch=$(echo `eval echo '$'"$app""_GIT_BRANCH"`)
      echo "    ""$app_name: [项目目录:$app_dir  代码分支:<$app_branch>]"
  done
  echo "    ""all 启动以上所有app"
}

# 获取运行方式
if [[ ! "$1" ]]; then
   help && exit
fi
case "$1" in
  "kill")
    RUN_CMD="kill";;
  "build")
    RUN_CMD="build";;
  "pullbuild")
    RUN_CMD="pullbuild";;
  "run")
    RUN_CMD="run";;
  "pullrun")
    RUN_CMD="pullrun";;
  "onlyrun")
    RUN_CMD="onlyrun";;
  *)
    echo "未知的启动命令: $1"
    help && exit
esac

# 获取运行的项目
if [[ ! "$2" ]]; then
   help && exit
fi
if [[ "$2" == "all" ]]; then
   RUN_APPS=${apps[*]}
else
   for app in ${apps[*]} ; do
      app_name=$(echo `eval echo '$'"$app""_APP_NAME"`)
      if [[ "$app_name" == "$2" ]]; then
        RUN_APPS=("$app")
      fi
   done
fi
if [[ ${#RUN_APPS[*]} -le 0 ]]; then
  echo "未知的app: $2"
  help && exit
fi

# 拉镜像
function git_pull() {
  local app=$1
  local app_dir=$(echo `eval echo '$'"$app""_DIR"`)
  local app_name=$(echo `eval echo '$'"$app""_APP_NAME"`)
  local app_branch=$(echo `eval echo '$'"$app""_GIT_BRANCH"`)

  echo "拉取$app_name代码, 分支<$app_branch>"
  cd $SHELL_FOLDER || exit
  cd $app_dir || exit
  git checkout $app_branch && git pull
}

# 杀进程
function kill_app() {
    local app=$1
    local app_name=$(echo `eval echo '$'"$app""_APP_NAME"`)
    local bin_dir="bin_${app_name}"
    local pid=`ps aux | grep "${bin_dir}/${app_name}" | grep -v grep | awk '{print $2}'`

    if [[ -n "${pid}" ]]; then
      kill -n 9 ${pid}
      echo "kill 正在运行的服务[$app_name] pid:$pid"
    else
      echo "服务[$app_name]未运行 不kill了~"
    fi
}

# 编译
function build_app() {
  local app=$1
  local app_dir=$(echo `eval echo '$'"$app""_DIR"`)
  local app_name=$(echo `eval echo '$'"$app""_APP_NAME"`)
  local bin_dir="bin_${app_name}"
  local conf_dir="${app_dir}/conf"

  echo "开始编译$app_name项目"

  if [ ! -d "$bin_dir" ]; then
        echo "创建app_name项目执行目录:$bin_dir"
        mkdir $bin_dir
  fi

  cd $SHELL_FOLDER || exit
  cd $app_dir || exit

  go mod tidy || exit
  go build -o "$app_name" "./cmd/main.go" || exit

  cd $SHELL_FOLDER || exit
  mv "${app_dir}/${app_name}" "${bin_dir}" || exit
  echo "移动${app_name}可执行文件至 --> ${app_dir}/${app_name}"
  cp -r $conf_dir $bin_dir || exit
  echo "拷贝${app_name}conf文件${conf_dir} --> ${bin_dir}目录"
}

# 运行
function run_app() {
    local app=$1
    local app_name=$(echo `eval echo '$'"$app""_APP_NAME"`)
    local bin_dir="bin_${app_name}"

    # 先杀了进程，在启动
    kill_app $app

    cd $SHELL_FOLDER || exit
    cd $bin_dir || exit

    nohup "../${bin_dir}/${app_name}" >> "../${bin_dir}/console.out" 2>&1 &

    local pid=`ps aux | grep "${bin_dir}/${app_name}" | grep -v grep | awk '{print $2}'`

    echo "启动$app_name完成, pid:$pid"
}


for app in ${RUN_APPS[*]} ; do
    app_name=$(echo `eval echo '$'"$app""_APP_NAME"`)
    echo ">>>>>>>>>>>> app:$app_name, cmd:$RUN_CMD"
    case $RUN_CMD in
    "kill")
      # 杀进程
      kill_app $app
      ;;
    "build")
      # 编译
      build_app $app
      ;;
    "pullbuild")
      # 拉代码
      git_pull $app
      # 编译
      build_app $app
      ;;
    "run")
      # 编译
      build_app $app
      # 运行
      run_app $app
      ;;
    "pullrun")
      # 拉代码
      git_pull $app
      # 编译
      build_app $app
      # 运行
      run_app $app
      ;;
    "onlyrun")
      # 运行
      run_app $app
      ;;
    *)
      echo "不知道的命令:$RUN_CMD"
      help && exit
    esac

done