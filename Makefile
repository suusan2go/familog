# vi: ft=make
.PHONY: proto test get ci docker

proto:
	go get github.com/golang/protobuf/protoc-gen-go
	protoc -I . -I ./thirdparty familog.proto \
	  --lile-server_out=. \
		--go_out=plugins=grpc:$$GOPATH/src --descriptor_set_out=api_descriptor.pb \
    --include_imports \
    --include_source_info

test:
	go test -v ./... -cover

get:
	go get -u -t ./...

ci: get test

docker-build:
	$(eval HASH := $(shell git rev-parse --short HEAD))
	docker build -t gcr.io/${PROJECT_ID}/familog:${HASH} .
	docker tag gcr.io/${PROJECT_ID}/familog:${HASH} gcr.io/${PROJECT_ID}/familog:${HASH}

docker-push:
	$(eval HASH := $(shell git rev-parse --short HEAD))
	gcloud docker -- push gcr.io/${PROJECT_ID}/familog:${HASH}
	gcloud docker -- push gcr.io/${PROJECT_ID}/familog:latest
