PACKAGE := github.com/sue445/go-doorkeeper
VERSION  := $(shell cat version.go | grep 'Version = ' | sed -E 's/^.*Version = "(.+)".*/\1/g')

.DEFAULT_GOAL := test

bin/$(NAME): $(SRCS)
	go build -o bin/$(NAME)

.PHONY: test
test:
	go test -count=1 $${TEST_ARGS} $(PACKAGE)

.PHONY: testrace
testrace:
	go test -count=1 $${TEST_ARGS} -race $(PACKAGE)

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golint -set_exit_status $(PACKAGE)

.PHONY: vet
vet:
	go vet $(PACKAGE)

.PHONY: tag
tag:
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push --tags

.PHONY: release
release: tag
	git push origin master
