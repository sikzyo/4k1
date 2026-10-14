default:
  @just --list

build:
  go build -o bin/4k1 ./cmd

run:
  go run ./cmd

clear: 
  rm -rf ./bin/
