SWAG_EXIST="$(command -v swag | wc -l | awk '{print $1}')"

if [[ $SWAG_EXIST == "0" ]]; then
  echo "Please execute command: go get -u github.com/swaggo/swag/cmd/swag"
  exit 0
fi

swag init --parseDependency --parseInternal -g main.go
