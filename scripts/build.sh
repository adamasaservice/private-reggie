#!/bin/bash

NAME=example.host,com/example/private-reggie

# get last commit hash 
function parse_git_hash() {
  git rev-list -1 HEAD
}

GIT_COMMIT=$(parse_git_hash)
VERSION=1.0.0

echo BUILD GO BINARY
cd cmd/main
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags  "-extldflags '-static' -X main.GitCommit=${GIT_COMMIT}" -o main .

echo BUILD DOCKER IMAGE
docker build -t ${NAME}:${VERSION} .

echo TAG DOCKER IMAGE
docker tag ${NAME}:${VERSION} ${NAME}:latest






