# PowerShell-specific build configuration
APP_NAME = markdown-parser
SRC = .\cmd\markdown-parser\main.go
BUILD_DIR = build
BINARY = ${BUILD_DIR}/${APP_NAME}

# Format the code
.PHONY: fmt
fmt:
	go fmt ./...

# Run Go vet
.PHONY: vet
vet: fmt
	go vet ./...

# Build the application
.PHONY: build
build: vet
	go build -o ./${BINARY}.exe ${SRC}

# Run the binary
.PHONY: run
run: build
	./${BINARY}.exe

# Run the main.go
.PHONY: frun
frun:
	go run ${SRC}

# Run the binary
.PHONY: exeqt
exeqt:
	./${BINARY}.exe

# Clean the build
.PHONY: clean
clean:
	go clean
	rm ./${BINARY}.exe

# Test the application
.PHONY: test
test:
	go test -v ./...
