default:
  @just --list

build-mac:
  GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o dist/4k1-darwin-arm64 ./cmd

build-linux:
  GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o dist/4k1-linux-amd64 ./cmd

build-all: build-mac build-linux

run:
  go run ./cmd

clear: 
  rm -rf ./dist/
