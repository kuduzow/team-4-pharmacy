.PHONY: all run dev lint fmt vet clean

APP_CMD := .\cmd\pharmacy
BIN := .\tmp\pharmacy

all: run


run:
	go run $(APP_CMD)

dev:
	air

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -rf ./tmp
