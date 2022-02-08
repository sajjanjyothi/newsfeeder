OPENAPI:=api/news.yaml
GOPATH:=$(shell go env GOPATH)
BUILD_NAME:=newsfeeder
RELEASE:=newsfeeder-0.1.0

codegen:
	mkdir -p newsfeeder
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen
	$(GOPATH)/bin/oapi-codegen -package newsfeeder $(OPENAPI) > newsfeeder/newsfeeder.gen.go
build:
	go build -o bin/$(BUILD_NAME)
build_docker:
	docker build . -t newsfeeder
deploy:
	 helm upgrade --install $(RELEASE) \
		deployment/newsfeeder
test:
	go test -cover ./...
clean:
	rm -rf bin