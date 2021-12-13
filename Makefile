# Makefile

BINARY_NAME=main

GIT_HASH=$(shell git log -1 --pretty=format:"%H") 
GIT_TAG=$(shell git tag --points-at HEAD) 


build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags  "-extldflags '-static' -X main.GitCommit=$(GIT_HASH)" -o main .

clean:
	go clean
	rm ${BINARY_NAME}

git_log:
	echo GIT_HASH=${GIT_HASH}
	echo GIT_TAG=${GIT_TAG}
