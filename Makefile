
export GOPATH :=${shell pwd}
export PATH := ${PATH}:${GOPATH}/bin
export GOBIN := ${GOPATH}/bin
main:
	go run main.go
build:
	go install main.go
images:
	docker build -t gotest .
run:
	docker run -it -v /Users/suifengluo/Documents/gitlab_work/gotest:/gopath --rm gotest
docker: images run


get-tools:
	go get github.com/suifengRock/go-tools

get:
	go get github.com/garyburd/redigo

hutu:
	go run src/hutu/hutu.go

log:
	go run src/golib/example/log.go

rand:
	go run src/golib/example/rand.go

ticker:
	go run src/golib/example/ticker.go

time:
	go run src/golib/example/time.go

test:
	go test -v -file src/golib/example/time.go

benchmark:
	go test -file src/golib/example/time.go  -test.bench="*" -v 