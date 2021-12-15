#!/usr/bin/env bash

# 安装air命令, 自动重启
AIR_EXIST="$(command -v air | wc -l | awk '{print $1}')"

if [ "$AIR_EXIST" == "0" ]; then
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
else
  echo "air already installed"
fi

# 安装swag命令
SWAG_EXIST="$(command -v swag | wc -l | awk '{print $1}')"

if [ "$SWAG_EXIST" == "0" ]; then
  go get github.com/swaggo/swag/cmd/swag
else
  echo "swag already installed"
fi


air -v && swag -v
