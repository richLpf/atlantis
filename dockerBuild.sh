#!/usr/bin/env bash
branch=$(git symbolic-ref HEAD | sed -e 's,.*/\(.*\),\1,')

echo "当前分支----"$branch
app_name=atlantis"_""$branch"
#默认推送 https://hub.docker.com/repository/docker/lvpf/golang-project

default_registry=https://hub.docker.com/repository/docker/lvpf/golang-project
default_project=atlantis

git pull origin $branch

