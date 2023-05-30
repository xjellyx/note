#!/bin/bash

echo "打包环境"
echo $ENV
node -v
npm -v


function test() {
    yarn
    # shellcheck disable=SC2181
    if [ $? -eq 0 ]; then
      echo "安装依赖完成"
    else
      echo "安装依赖失败"
      exit 1
    fi
    # build
    yarn build
    if [ $? -eq 0 ];then
      echo "打包完成"
    else
      echo "打包失败"
      exit 1
    fi
    #
    # shellcheck disable=SC2164
    cd dist
    echo starwiz@123 |  sudo -S cp -r *  /home/hkatg/app/ghk-data-api/nginx/web/hkatg-frontend/
    if [ $? -eq 0 ];then
      echo "成功打包文件转移到指定目录"
    else
      exit 1
    fi
}

if [ "$ENV" == "prod" ];then
  echo "生产"
else
  echo "测试"
  test
fi
