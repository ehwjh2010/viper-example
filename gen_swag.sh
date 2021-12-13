#!/usr/bin/env bash
PROJECT_PATH="/Users/jh/Documents/Project/self/GoProject/cobra-example"

SWAG_EXIST="$(command -v swag | wc -l | awk '{print $1}')"

if [[ $SWAG_EXIST == "0" ]]; then
  echo "Please execute command: go get github.com/swaggo/swag/cmd/swag"
  exit 0
fi

cd "$PROJECT_PATH" && swag init --parseDependency --parseInternal -g main.go
