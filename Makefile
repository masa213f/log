TARGET = sample
SOURCE = $(shell find . -type f -name "*.go" -not -name "*_test.go")

all: build

setup:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/lint/golint

mod:
	go mod tidy
	go mod vendor

build: mod $(TARGET)

$(TARGET): go.mod $(SOURCE)
	CGO_ENABLED=0 go build -o $(TARGET) -ldflags "-X main.version=$(VERSION)" ./cmd/$(TARGET)/...

run: $(TARGET)
	./$(TARGET)

clean:
	-rm $(TARGET)

distclean: clean
	-rm go.sum
	-rm -rf vendor

fmt:
	goimports -w $$(find . -type d -name 'vendor' -prune -o -type f -name '*.go' -print)

test:
	test -z "$$(goimports -l $$(find . -type d -name 'vendor' -prune -o -type f -name '*.go' -print) | tee /dev/stderr)"
	#test -z "$$(golint $$(go list ./... | grep -v '/vendor/') | tee /dev/stderr)"
	CGO_ENABLED=0 go test -v ./...

.PHONY: all setup mod build run clean distclean fmt test
