 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=GinRestAPi
BINARY_UNIX=$(BINARY_NAME)Unix
BINARY_WIN=$(BINARY_NAME)Win.exe
SHELL := /bin/bash
BASEDIR = $(shell pwd)

build: gotool
	go build -v  .

swagger:
	swag init

run:
	go run main.go

runswag: swagger
	go run main.go

gotool:
	gofmt -w .

lint:
	golangci-lint run

ca:
	openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ca files"

cross_build: build_windows build-linux build_mac
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

build_windows:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_WIN) -v

build_mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_MAC) -v

exec_win:
	@./main_win.exe

install:
	go install
cloc:
	cloc . --exclude-dir=.idea --out=cloc.txt	#--by-file  #order by file line height

wrk:
	wrk -t12 -c400 -d30s --latency http://127.0.0.1:6663/vps/health

.PHONY: clean gotool ca help build_cross build_windows build-linux exec_win cloc swagger run runswag