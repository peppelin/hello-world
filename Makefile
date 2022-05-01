SHELL=/bin/bash -e -o pipefail
PWD = $(shell pwd)

# constants
GOLANGCI_VERSION = 1.18
IMAGE_TAG = latest

######################################################
# misc
######################################################

out:
	@mkdir -p out

out/bin:
	@mkdir -p out/bin

######################################################
# go
######################################################

go.mod:
	go mod init github.com/peppelin/hello-world

.PHONY: tidy
tidy: ## clean up go.mod and go.sum
	go mod tidy

.PHONY: vendor
vendor: ## vendor all packages to ./vendor
	go mod vendor

.PHONY: download
download: ## downloads the dependencies
	go mod download -x

.PHONY: run
run: ## run the service broker
	go run main.go

######################################################
# clean
######################################################

.PHONY: clean-bin
clean-bin: ## clean local binary folders
	@rm -rf bin testbin

.PHONY: clean-outputs
clean-outputs: ## clean output folders out, vendor
	@rm -rf out vendor api/proto/google api/proto/validate

.PHONY: clean-docker
clean-docker: ## clean previous docker images
	@docker rmi -f hello-world

.PHONY: clean
clean: clean-bin clean-outputs clean-docker ## clean up everything

######################################################
# lint
######################################################

bin/golangci-lint: bin/golangci-lint-$(GOLANGCI_VERSION)
	@ln -sf golangci-lint-$(GOLANGCI_VERSION) $@

bin/golangci-lint-$(GOLANGCI_VERSION):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v$(GOLANGCI_VERSION)
	@mv bin/golangci-lint $@

.PHONY: lint
lint: bin/golangci-lint out download ## lint all code with golangci-lint
	bin/golangci-lint run ./... --timeout 15m0s

######################################################
# test
######################################################

.PHONY: test
test: out download ## run all tests
	go test -v -coverpkg=./... -coverprofile=coverage.cov ./...

######################################################
# build
######################################################

.PHONY: build
build: download out/bin ## build all binaries
	CGO_ENABLED=0 go build -ldflags="-w -s" -o out/bin ./...

.PHONY: build-linux
build-linux: download out/bin ## build all binaries for linux
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-w -s" -o out/bin ./...

.PHONY: build-arm
build-arm: download out/bin ## build all binaries for arm
	CGO_ENABLED=0 GOARCH=arm GOOS=linux go build -ldflags="-w -s" -o out/bin ./...

######################################################
# docker
######################################################

.PHONY: docker
docker: clean build ## build docker image
	@docker build -t hello-world .
docker-linux: clean build-linux ## build docker image for linux
	@docker build -t hello-world .
docker-arm: clean build-arm ## build docker image for arm
	@docker build -t hello-world .

######################################################
# help
######################################################

.PHONY: help
help: ## show help
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
        awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ''