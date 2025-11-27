run:
		go run cmd/pharmacy/main.go

dev:
		air

lint:
		golangci-lint run ./...

fmt: 
		golangci-lint fmt

vet:
		go vet ./...