# vi: ft=make
.PHONY: proto test get ci docker

proto:
	go get github.com/golang/protobuf/protoc-gen-go
	protoc -I . familog.proto --lile-server_out=. --go_out=plugins=grpc:$$GOPATH/src

test:
	go test -v ./... -cover

get:
	go get -u -t ./...

ci: get test

docker-build:
	docker build -t gcr.io/${PROJECT_ID}/familog:latest .

docker-push:
	gcloud docker -- push gcr.io/${PROJECT_ID}/familog:latest
