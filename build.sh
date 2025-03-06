#!/usr/bin/zsh

if ! [[ -d bin ]]; then
  mkdir bin
fi

if ! [[ -d dist ]]; then
  mkdir dist
fi

if [[ $# -eq 0 ]]; then
  echo "you must specify the name of cmd/<package> as first arg"
  exit 1
fi

if ! [[ -d cmd/"$1" ]]; then
  echo "cmd/$1 doesn't exist"
  exit 1
fi

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
go build -o bin/"$1"/bootstrap cmd/"$1"/main.go && \
echo "completed building binary at $1/bootstrap"; \
chmod +x bin/"$1"/bootstrap && \
rm dist/"$1".zip;
cd bin/"$1" && \
7z a -tzip -mx=9 ../../dist/"$1".zip bootstrap;
