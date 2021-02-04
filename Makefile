GO_IMPORT := google/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor


.PHONE: all
all: test

.PHONY: clean
clean:
	@rm -rf *.pb.go

deps:
	@go mod download

protoc:
	@protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./*.proto

build: clean protoc

publish:
	scripts/publish.sh

build_publish: build publish

