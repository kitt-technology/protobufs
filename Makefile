GO_IMPORT := google/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor


.PHONE: all
all: test

deps:
	@go mod download

build:
	@go install .


