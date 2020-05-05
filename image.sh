#!/usr/bin/env bash
e=$1

echo "当前环境"-$e
app_name=atlantis
version_num=$e

default_registry=hub.docker.com/repository/docker
default_project=lvpf/golang-project
default_user=lpengfei66@163.com
default_pwd=lpf12345@

if [ "$e" != pre ]; then
	default_registry=hub.docker.com/repository/docker
	default_user=lvpf
fi

echo $default_registry

echo $(date "+%Y-%m-%d %H:%M:%S")-开始制作镜像
docker build --rm -t $app_name:$version_num .
echo $(date "+%Y-%m-%d %H:%M:%S")-镜像制作完成

echo $(date "+%Y-%m-%d %H:%M:%S")-开始打tag
echo $app_name:$version_num
docker tag $app_name:$version_num $default_registry/$default_project/$app_name\:$version_num
echo $(date "+%Y-%m-%d %H:%M:%S")-打tag完成

echo $(date "+%Y-%m-%d %H:%M:%S")-开始push镜像
docker login -u $default_user -p $default_pwd
docker push $default_registry/$default_project/$app_name\:$version_num
echo $(date "+%Y-%m-%d %H:%M:%S")-push镜像完成
