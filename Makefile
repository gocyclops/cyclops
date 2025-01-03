.PHONY:fmt vet build install
fmt:
	go fmt ./...
vet:
	go vet ./...
build:
	go build
install:
	go install
