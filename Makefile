PACKAGE := github.com/sue445/go-doorkeeper

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
