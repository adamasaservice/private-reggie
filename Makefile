# Makefile


build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags  "-extldflags '-static' -X main.GitCommit=$(Build.SourceVersion)" -o main .

go_coverage:
	go test -cover
