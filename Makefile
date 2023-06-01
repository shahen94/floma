default: run

.PHONY: clean build run

clean:
	echo "Cleaning up..."
	@rm -rf bin
build:
	@echo "Building..."
	@go build -o bin/floma cmd/floma/main.go
run:
	@go run cmd/floma/main.go --config floma.example.yml