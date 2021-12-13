# Makefile

BINARY_NAME=main
GIT_HASH=$(shell git log -1 --pretty=format:"%H") 
GIT_TAG=$(shell git tag --points-at HEAD) 


build_linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags  "-extldflags '-static' -X main.GitCommit=$(GIT_HASH) main.GitTag=$(GIT_TAG)" -o main .

build_native:
	go build -ldflags="-X 'main.GitCommit=$(GIT_HASH)' -X 'main.GitTag=$(GIT_TAG)' -X 'main.Version=$(GIT_TAG)'" -o main .

clean:
	go clean
	rm ${BINARY_NAME}

git_log:
	echo GIT_HASH=${GIT_HASH}
	echo GIT_TAG=${GIT_TAG}
