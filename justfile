default:
  @just --list

build:
  go build -o bin/4k1_darwin ./cmd

run:
  go run ./cmd

clear: 
  rm -rf ./bin/
