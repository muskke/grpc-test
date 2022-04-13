#enviroments
NAME=grpc-test-server
CGO_ENABLED=0
GO111MODULE=on

#vars
GOPATH:=$(shell go env GOPATH)
GOROOT:=$(shell go env GOROOT)
empty :=
space := $(empty) $(empty)
version:=v3

#go imports
VALIDATE_IMPORT := Mvalidate/validate.proto=github.com/envoyproxy/protoc-gen-validate/validate

GO_IMPORT_SPACES := ${VALIDATE_IMPORT},\
	Mgoogle/protobuf/any.proto=google.golang.org/protobuf/types/known/anypb,\
	Mgoogle/protobuf/duration.proto=google.golang.org/protobuf/types/known/durationpb,\
	Mgoogle/protobuf/struct.proto=google.golang.org/protobuf/types/known/structpb,\
	Mgoogle/protobuf/timestamp.proto=google.golang.org/protobuf/types/known/timestamppb,\
	Mgoogle/protobuf/wrappers.proto=google.golang.org/protobuf/types/known/wrapperspb,\
	Mgoogle/protobuf/descriptor.proto=google.golang.org/protobuf/types/descriptorpb

GO_IMPORT:=$(subst $(space),,$(GO_IMPORT_SPACES))

all: build

.PHONY: init
init:
	go get -u -d google.golang.org/protobuf/proto
	go get -u -d github.com/envoyproxy/protoc-gen-validate
	go get -u -d github.com/grpc-ecosystem/grpc-gateway
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/micro-community/micro/v3/cmd/protoc-gen-micro@master

build:
	go env -w GO111MODULE=auto
	go env -w GOPROXY=http://nexus.infore-robotics.cn/repository/golang,https://goproxy.io,https://goproxy.cn,direct
	go env -w CGO_ENABLED=0
	go mod download
	go build -v -o app/$(NAME)

run:
	./app/$(NAME)
	
.PHONY: test
test:
	go test -v ./... -cover

.PHONY: protos
protos:
	protoc -I "${GOPATH}/src"  -I proto/${version} \
	-I "${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis" \
	--plugin="protoc-gen-go=${GOPATH}/bin/protoc-gen-go.exe"  \
	--plugin="protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro.exe"  \
	--plugin="protoc-gen-validate=${GOPATH}/bin/protoc-gen-validate.exe"  \
	--go_out="${GO_IMPORT},paths=source_relative:proto/${version}"  \
	--go-grpc_out="${GO_IMPORT},paths=source_relative:proto/${version}"  \
	--micro_out="${GO_IMPORT},plugins=grpc,paths=source_relative:proto/${version}"   \
	--validate_out="lang=go:proto/${version}"   \
	--swagger_out=logtostderr=true:proto/${version}  \
	grpc_test.proto

